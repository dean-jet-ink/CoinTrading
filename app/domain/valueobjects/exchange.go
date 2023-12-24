package valueobjects

import (
	"cointrading/app/domain/errors"
	"fmt"
	"net/http"
)

var (
	bitflyer = &Exchange{1}
	bybit    = &Exchange{2}
	bitget   = &Exchange{3}
)

var exchangeNames = map[int]string{
	1: "Bitflyer",
	2: "Bybit",
	3: "Bitget",
}

type Exchange struct {
	value int
}

func NewExchange(value int) (*Exchange, error) {
	if _, ok := exchangeNames[value]; !ok {
		message := fmt.Sprint("非対応の取引所です")
		original := fmt.Sprintf("Unexpected exchange code %v", value)

		myerr := errors.NewMyError(message, original, http.StatusInternalServerError)

		return nil, myerr
	}

	return &Exchange{
		value: value,
	}, nil
}

func (e *Exchange) Value() int {
	return e.value
}

func (e *Exchange) DisplayValue() string {
	return exchangeNames[e.value]
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
