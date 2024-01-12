package bitflyer

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/myerror"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/valueobjects"
	"cointrading/app/infrastructure/apiclients"
	"cointrading/app/lib"
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

type bitflyerClient struct {
	key        string
	secret     string
	httpClient *apiclients.ClientBase
}

func NewBitflyerClient(key, secret string) repositories.TradingAPIClient {
	return &bitflyerClient{
		key:        key,
		secret:     secret,
		httpClient: apiclients.NewClientBase("https://api.bitflyer.com/v1/"),
	}
}

func (b *bitflyerClient) GetBalance() ([]*entities.Balance, error) {
	method := "GET"
	path := "me/getbalance"

	header, err := b.privateAPIHeader(method, path, nil)
	if err != nil {
		return nil, err
	}

	balances := []*balance{}

	if err := b.httpClient.Request(method, path, nil, header, nil, balances); err != nil {
		return nil, err
	}

	var balanceEntities []*entities.Balance
	for _, balance := range balances {
		code, ok := b.convertCurrencyCodeToInt(balance.CurrencyCode)
		if !ok {
			continue
		}

		balanceEntity, err := entities.NewBalance(code, balance.Amount, balance.Available)
		if err != nil {
			return nil, err
		}

		balanceEntities = append(balanceEntities, balanceEntity)
	}

	return balanceEntities, nil
}

func (b *bitflyerClient) GetTicker(symbol *valueobjects.Symbol) (*entities.Ticker, error) {
	method := "GET"
	path := "ticker"

	symbolStr, err := b.convertSymbolToString(symbol)
	if err != nil {
		return nil, err
	}

	query := map[string]string{
		"product_code": symbolStr,
	}

	header := b.publicAPIHeader()

	ticker := &ticker{}

	if err := b.httpClient.Request(method, path, query, header, nil, ticker); err != nil {
		return nil, err
	}

	dateTime, err := lib.StringToDateTime(ticker.Timestamp)
	if err != nil {
		return nil, err
	}

	entity, err := entities.NewTicker(symbol.Value(), dateTime, ticker.BestAsk, ticker.BestBid, ticker.Volume)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (b *bitflyerClient) GetRealTimeTicker(ctx context.Context, symbol *valueobjects.Symbol, tickerChan chan<- *entities.Ticker, errChan chan<- error) {
	// wss://ws.lightstream.bitflyer.com/json-rpc
	u := url.URL{
		Scheme: "wss",
		Host:   "ws.lightstream.bitflyer.com",
		Path:   "json-rpc",
	}

	dialer := apiclients.NewWebSocketDialer()

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		errChan <- err
		return
	}

	defer conn.Close()

	symbolStr, err := b.convertSymbolToString(symbol)
	if err != nil {
		errChan <- err
		return
	}

	channel := fmt.Sprintf("lightning_ticker_%s", symbolStr)

	message := &apiclients.JsonRPC2{
		Version: apiclients.Version,
		Method:  apiclients.Subscribe,
		Params: map[string]string{
			"channel": channel,
		},
	}

	if err := conn.WriteJSON(message); err != nil {
		errChan <- err
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			b.processMessage(conn, symbol, tickerChan, errChan)
		}

	}
}

func (b *bitflyerClient) processMessage(conn *websocket.Conn, symbol *valueobjects.Symbol, tickerChan chan<- *entities.Ticker, errChan chan<- error) {
	message := &apiclients.JsonRPC2{}
	if err := conn.ReadJSON(message); err != nil {
		errChan <- err
		return
	}

	if message.Method == "channelMessage" {
		switch params := message.Params.(type) {
		case map[string]any:
			for key, binary := range params {
				if key == "message" {
					marshalTicker, err := json.Marshal(binary)
					if err != nil {
						// 空のデータが送信されている場合
						return
					}

					ticker := &ticker{}
					if err := json.Unmarshal(marshalTicker, ticker); err != nil {
						// Ticker以外のデータが送信されている場合
						return
					}

					dateTime, err := lib.StringToDateTime(ticker.Timestamp)
					if err != nil {
						errChan <- err
						return
					}

					entity, err := entities.NewTicker(symbol.Value(), dateTime, ticker.BestAsk, ticker.BestBid, ticker.Volume)
					if err != nil {
						errChan <- err
						return
					}

					tickerChan <- entity
				}
			}
		}
	}
}

func (b *bitflyerClient) SendOrder(ord *entities.Order) (string, error) {
	method := "POST"
	path := "me/sendchildorder"

	orderModel, err := b.entityToOrder(ord)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(orderModel)
	if err != nil {
		return "", err
	}

	header, err := b.privateAPIHeader(method, path, data)
	if err != nil {
		return "", err
	}

	orderResponse := &orderResponse{}

	if err := b.httpClient.Request(method, path, nil, header, data, orderResponse); err != nil {
		return "", err
	}

	return orderResponse.ChildOrderAcceptanceID, nil
}

func (b *bitflyerClient) GetOrder(orderId string) (*entities.Order, error) {
	method := "GET"
	path := "me/getchildorders"

	header, err := b.privateAPIHeader(method, path, nil)
	if err != nil {
		return nil, err
	}

	query := map[string]string{
		"child_order_acceptance_id": orderId,
	}

	orders := []*order{}

	if err := b.httpClient.Request(method, path, query, header, nil, orders); err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		err := fmt.Errorf("%w: Order is not found by id: %v", myerror.ErrNotFoundOrder, orderId)
		return nil, err
	}

	order := orders[0]

	entity, err := b.orderToEntity(order)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (b *bitflyerClient) GetOrders(symbol *valueobjects.Symbol) ([]*entities.Order, error) {
	method := "GET"
	path := "me/getchildorders"

	symbolStr, err := b.convertSymbolToString(symbol)
	if err != nil {
		return nil, err
	}

	header, err := b.privateAPIHeader(method, path, nil)
	if err != nil {
		return nil, err
	}

	query := map[string]string{
		"product_code": symbolStr,
	}

	orders := []*order{}

	if err := b.httpClient.Request(method, path, query, header, nil, orders); err != nil {
		return nil, err
	}

	orderEntities := []*entities.Order{}

	for _, order := range orders {
		entity, err := b.orderToEntity(order)
		if err != nil {
			return nil, err
		}

		orderEntities = append(orderEntities, entity)
	}

	return orderEntities, nil
}
