package valueobjects

import (
	"cointrading/app/domain/errors"
	"fmt"
	"net/http"
)

var (
	JPY = &CurrencyCode{1}
	BTC = &CurrencyCode{2}
	ETH = &CurrencyCode{3}
)

var currencyCodes = map[int]string{
	1: "円",
	2: "BTC",
	3: "ETH",
}

type CurrencyCode struct {
	value int
}

func NewCurrencyCode(value int) (*CurrencyCode, error) {
	if _, ok := currencyCodes[value]; !ok {
		message := "非対応の貨幣コードです"
		original := fmt.Sprintf("Unexpected currency code %v", value)
		myerr := errors.NewMyError(message, original, http.StatusInternalServerError)

		return nil, myerr
	}

	return &CurrencyCode{
		value: value,
	}, nil
}

func (c *CurrencyCode) Value() int {
	return c.value
}

func (c *CurrencyCode) DisplayValue() string {
	return currencyCodes[c.value]
}

func CurrencyCodes() []*CurrencyCode {
	return []*CurrencyCode{
		JPY,
		BTC,
		ETH,
	}
}
