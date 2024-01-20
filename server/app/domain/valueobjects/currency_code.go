package valueobjects

import (
	"cointrading/app/domain/myerror"
	"fmt"
)

var (
	JPY = &CurrencyCode{1}
	BTC = &CurrencyCode{2}
	ETH = &CurrencyCode{3}
	XRP = &CurrencyCode{4}
)

var currencyCodes = map[CurrencyCode]string{
	*JPY: "円",
	*BTC: "BTC",
	*ETH: "ETH",
	*XRP: "XRP",
}

type CurrencyCode struct {
	value int
}

func NewCurrencyCode(value int) (*CurrencyCode, error) {
	currencyCode := &CurrencyCode{
		value: value,
	}

	if _, ok := currencyCodes[*currencyCode]; !ok {
		err := fmt.Errorf("%w: Unexpected currency code: %v", myerror.ErrUnexpectedCurrencyCode, value)

		return nil, err
	}

	return currencyCode, nil
}

func (c *CurrencyCode) Value() int {
	return c.value
}

func (c *CurrencyCode) DisplayValue() string {
	return currencyCodes[*c]
}

func CurrencyCodes() []*CurrencyCode {
	return []*CurrencyCode{
		JPY,
		BTC,
		ETH,
		XRP,
	}
}
