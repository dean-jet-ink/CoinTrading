package apiclient

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/valueobject"
	"cointrading/app/lib"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type BitflyerClient struct {
	key        string
	secret     string
	httpClient *clientBase
}

func NewBitflyerClient(key, secret string) repositories.TradingAPIClient {
	return &BitflyerClient{
		key:        key,
		secret:     secret,
		httpClient: newClientBase("https://api.bitflyer.com/v1/"),
	}
}

type balance struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Available    float64 `json:"available"`
}

func (b *BitflyerClient) GetBalance() ([]*entities.Balance, error) {
	method := "GET"
	path := "me/getbalance"

	endpoint, err := b.httpClient.createEndpoint(path)
	if err != nil {
		return nil, err
	}

	header, err := b.privateAPIHeader(method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := b.httpClient.request(method, endpoint, nil, header, nil)
	if err != nil {
		return nil, err
	}

	var balances []*balance
	if err := b.httpClient.unmarshalResponse(resp, balances); err != nil {
		return nil, err
	}

	var balanceEntities []*entities.Balance
	for _, balance := range balances {
		balanceEntity := entities.NewBalance(balance.CurrencyCode, balance.Amount, balance.Available)
		balanceEntities = append(balanceEntities, balanceEntity)
	}

	return balanceEntities, nil
}

type ticker struct {
	ProductCode     string  `json:"product_code"`
	State           string  `json:"state"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func (t *ticker) midPrice() float64 {
	return (t.BestAsk + t.BestBid) / 2
}

func (b *BitflyerClient) GetTicker(symbol *valueobject.Symbol) (*entities.Ticker, error) {
	method := "GET"
	path := "ticker"

	symbolStr := b.convertSymbol(symbol)

	endpoint, err := b.httpClient.createEndpoint(path)
	if err != nil {
		return nil, err
	}

	query := map[string]string{
		"product_code": symbolStr,
	}

	header := b.publicAPIHeader()

	resp, err := b.httpClient.request(method, endpoint, query, header, nil)
	if err != nil {
		return nil, err
	}

	var ticker *ticker
	if err := b.httpClient.unmarshalResponse(resp, ticker); err != nil {
		return nil, err
	}

	dateTime, err := lib.StringToDateTime(ticker.Timestamp)
	if err != nil {
		return nil, err
	}

	entity := entities.NewTicker(symbol.Value(), dateTime, ticker.midPrice(), ticker.Volume)

	return entity, nil
}

func (b *BitflyerClient) GetRealTimeTicker(symbol *valueobject.Symbol, tickerChan chan<- *entities.Ticker) error {
	// wss://ws.lightstream.bitflyer.com/json-rpc
	u := url.URL{
		Scheme: "wss",
		Host:   "ws.lightstream.bitflyer.com",
		Path:   "json-rpc",
	}

	dialer := NewWebSocketDialer()

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	defer conn.Close()

	symbolStr := b.convertSymbol(symbol)

	channel := fmt.Sprintf("lightning_ticker_%s", symbolStr)

	message := &JsonRPC2{
		Version: "2.0",
		Method:  "subscribe",
		Params: map[string]string{
			"channel": channel,
		},
	}

	if err := conn.WriteJSON(message); err != nil {
		return err
	}

OUTER:
	for {
		message = &JsonRPC2{}
		if err := conn.ReadJSON(message); err != nil {
			return err
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

						var ticker *entities.Ticker
						if err := json.Unmarshal(marshalTicker, ticker); err != nil {
							// Ticker以外のデータが送信されている場合
							continue OUTER
						}

						tickerChan <- ticker
					}
				}
			}
		}
	}
}

func (b *BitflyerClient) privateAPIHeader(method, endpoint string, body []byte) (map[string]string, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	message := timestamp + method + endpoint + string(body)

	h := hmac.New(sha256.New, []byte(b.secret))

	if _, err := h.Write([]byte(message)); err != nil {
		return nil, err
	}

	sign := hex.EncodeToString(h.Sum(nil))

	header := map[string]string{
		"ACCESS-KEY":       b.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}

	return header, nil
}

func (b *BitflyerClient) publicAPIHeader() map[string]string {
	header := map[string]string{
		"Content-Type": "application/json",
	}

	return header
}

func (b *BitflyerClient) convertSymbol(symbol *valueobject.Symbol) string {
	switch {
	case symbol.IsBTCJPY():
		return "BTC_JPY"
	case symbol.IsBTCUSD():
		return "BTC_USD"
	case symbol.IsETHJPY():
		return "ETH_JPY"
	case symbol.IsETHUSD():
		return "ETH_USD"
	default:
		return "Unknown"
	}
}
