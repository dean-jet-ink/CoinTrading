package valueobjects

import (
	"cointrading/app/domain/errors"
	"fmt"
	"net/http"
)

var (
	market = &OrderType{1}
	limit  = &OrderType{2}
	stop   = &OrderType{3}
)

var orderTypes = map[int]string{
	1: "MARKET",
	2: "LIMIT",
	3: "STOP",
}

type OrderType struct {
	value int
}

func NewOrderType(value int) (*OrderType, error) {
	if _, ok := orderTypes[value]; !ok {
		message := "無効なオーダータイプです"
		original := fmt.Sprintf("Invalied order type code: %v", value)
		myerr := errors.NewMyError(message, original, http.StatusInternalServerError)

		return nil, myerr
	}

	return &OrderType{
		value: value,
	}, nil
}

func (o *OrderType) Value() int {
	return o.value
}

func (o *OrderType) DisplayValue() string {
	return orderTypes[o.value]
}

func (o *OrderType) IsMarket() bool {
	return *o == *market
}

func (o *OrderType) IsLimit() bool {
	return *o == *limit
}

func (o *OrderType) IsStop() bool {
	return *o == *stop
}

func OrderTypes() []*OrderType {
	return []*OrderType{
		market,
		limit,
		stop,
	}
}
