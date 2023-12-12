package valueobject

import "strings"

var codeMap = map[string]string{
	"btc": "Bitcoin",
	"eth": "Ethereum",
}

type CurrencyCode struct {
	value string
}

func NewCurrencyCode(value string) *CurrencyCode {
	lowerValue := strings.ToLower(value)

	return &CurrencyCode{
		value: lowerValue,
	}
}

func (c *CurrencyCode) Value() string {
	return c.value
}

func (c *CurrencyCode) DisplayValue() string {
	coinName, ok := codeMap[c.value]

	if !ok {
		return "Unknown"
	}

	return coinName
}
