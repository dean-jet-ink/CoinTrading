package valueobjects

import (
	"cointrading/app/domain/errors"
	"fmt"
)

var (
	bitflyer = &Exchange{1}
	bybit    = &Exchange{2}
	bitget   = &Exchange{3}
)

var exchangeNames = map[Exchange]string{
	*bitflyer: "Bitflyer",
	*bybit:    "Bybit",
	*bitget:   "Bitget",
}

type Exchange struct {
	value int
}

func NewExchange(value int) (*Exchange, error) {
	exchange := &Exchange{
		value: value,
	}

	if _, ok := exchangeNames[*exchange]; !ok {
		message := fmt.Sprint("非対応の取引所です")
		original := fmt.Sprintf("Unexpected exchange code %v", value)

		myerr := errors.NewMyError(message, original, 500)

		return nil, myerr
	}

	return exchange, nil
}

func (e *Exchange) Value() int {
	return e.value
}

func (e *Exchange) String() string {
	exchangeName, _ := exchangeNames[*e]

	return exchangeName
}

func (e *Exchange) IsBitflyer() bool {
	return *e == *bitflyer
}

func (e *Exchange) IsBybit() bool {
	return *e == *bybit
}

func (e *Exchange) IsBitget() bool {
	return *e == *bitget
}

func Exchanges() []*Exchange {
	return []*Exchange{
		bitflyer,
		bybit,
		bitget,
	}
}
