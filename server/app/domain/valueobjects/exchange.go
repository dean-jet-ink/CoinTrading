package valueobjects

import (
	"cointrading/app/domain/myerror"
	"fmt"
	"strings"
)

var (
	Bitflyer = &Exchange{1}
	Bybit    = &Exchange{2}
	Bitget   = &Exchange{3}
)

var exchangeNames = map[Exchange]string{
	*Bitflyer: "Bitflyer",
	*Bybit:    "Bybit",
	*Bitget:   "Bitget",
}

type Exchange struct {
	value int
}

func NewExchange(value int) (*Exchange, error) {
	exchange := &Exchange{
		value: value,
	}

	if _, ok := exchangeNames[*exchange]; !ok {
		err := fmt.Errorf("%w: Unexpected exchange code %v", myerror.ErrUnexpectedExchange, value)

		return nil, err
	}

	return exchange, nil
}

func (e *Exchange) Value() int {
	return e.value
}

func (e *Exchange) DisplayValue() string {
	return exchangeNames[*e]
}

func (e *Exchange) DisplayValueForTableName() string {
	return strings.ToLower(exchangeNames[*e])
}

func (e *Exchange) IsBitflyer() bool {
	return *e == *Bitflyer
}

func (e *Exchange) IsBybit() bool {
	return *e == *Bybit
}

func (e *Exchange) IsBitget() bool {
	return *e == *Bitget
}

func Exchanges() []*Exchange {
	return []*Exchange{
		Bitflyer,
		// Bybit,
		// Bitget,
	}
}
