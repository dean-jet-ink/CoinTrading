package bitflyer

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/valueobjects"
	"cointrading/app/infrastructure/apiclients"
	"cointrading/app/lib"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
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
		code, ok := b.convertCurrencyCode(balance.CurrencyCode)
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

	symbolStr, err := b.convertSymbol(symbol)
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

	entity := entities.NewTicker(symbol, dateTime, ticker.BestAsk, ticker.BestBid, ticker.Volume)

	return entity, nil
}

func (b *bitflyerClient) GetRealTimeTicker(symbol *valueobjects.Symbol, tickerChan chan<- *entities.Ticker) {
	// wss://ws.lightstream.bitflyer.com/json-rpc
	u := url.URL{
		Scheme: "wss",
		Host:   "ws.lightstream.bitflyer.com",
		Path:   "json-rpc",
	}

	dialer := apiclients.NewWebSocketDialer()

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	symbolStr, err := b.convertSymbol(symbol)
	if err != nil {
		log.Println(err)
		return
	}

	channel := fmt.Sprintf("lightning_ticker_%s", symbolStr)

	message := &apiclients.JsonRPC2{
		Version: "2.0",
		Method:  "subscribe",
		Params: map[string]string{
			"channel": channel,
		},
	}

	if err := conn.WriteJSON(message); err != nil {
		log.Println(err)
		return
	}

OUTER:
	for {
		message = &apiclients.JsonRPC2{}
		if err := conn.ReadJSON(message); err != nil {
			log.Println(err)
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
							continue OUTER
						}

						ticker := &ticker{}
						if err := json.Unmarshal(marshalTicker, ticker); err != nil {
							// Ticker以外のデータが送信されている場合
							continue OUTER
						}

						dateTime, err := lib.StringToDateTime(ticker.Timestamp)
						if err != nil {
							continue OUTER
						}

						entity := entities.NewTicker(symbol, dateTime, ticker.BestAsk, ticker.BestBid, ticker.Volume)

						tickerChan <- entity
					}
				}
			}
		}
	}
}

func (b *bitflyerClient) SendOrder(ord *entities.Order) (string, error) {
	method := "POST"
	path := "me/sendchildorder"

	header, err := b.privateAPIHeader(method, path, nil)
	if err != nil {
		return "", err
	}

	symbolStr, err := b.convertSymbol(ord.Symbol())
	if err != nil {
		return "", err
	}

	orderTypeStr, err := b.convertOrderType(ord.OrderType())
	if err != nil {
		return "", err
	}

	timeInForceStr, err := b.convertTimeInForce(ord.TimeInForce())
	if err != nil {
		return "", err
	}

	orderModel := &order{
		ProductCode:    symbolStr,
		ChildOrderType: orderTypeStr,
		Side:           strings.ToUpper(ord.Side().Value()),
		Price:          ord.Price(),
		Size:           ord.Size(),

		// 期限1日
		MinuteToExpire: 1440,

		TimeInForce: timeInForceStr,
	}

	data, err := json.Marshal(orderModel)
	if err != nil {
		return "", err
	}

	orderResponse := &orderResponse{}

	if err := b.httpClient.Request(method, path, nil, header, data, orderResponse); err != nil {
		return "", err
	}

	return orderResponse.ChildOrderAcceptanceID, nil
}
