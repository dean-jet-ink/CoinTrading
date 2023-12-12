package apiclient

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/repositories"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

	type Balance struct {
		CurrencyCode string  `json:"currency_code"`
		Amount       float64 `json:"amount"`
		Available    float64 `json:"available"`
	}

	var balances []*Balance
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
