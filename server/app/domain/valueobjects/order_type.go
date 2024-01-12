package valueobjects

import (
	"cointrading/app/domain/myerror"
	"fmt"
)

var (
	Market = &OrderType{1}
	Limit  = &OrderType{2}
	Stop   = &OrderType{3}
)

var orderTypes = map[OrderType]string{
	*Market: "MARKET",
	*Limit:  "LIMIT",
	*Stop:   "STOP",
}

type OrderType struct {
	value int
}

func NewOrderType(value int) (*OrderType, error) {
	orderType := &OrderType{
		value: value,
	}

	if _, ok := orderTypes[*orderType]; !ok {
		err := fmt.Errorf("%w: Unexpected order type code: %v", myerror.ErrUnexpectedOrderType, value)

		return nil, err
	}

	return orderType, nil
}

func (o *OrderType) Value() int {
	return o.value
}

func (o *OrderType) DisplayValue() string {
	return orderTypes[*o]
}

func (o *OrderType) IsMarket() bool {
	return *o == *Market
}

func (o *OrderType) IsLimit() bool {
	return *o == *Limit
}

func (o *OrderType) IsStop() bool {
	return *o == *Stop
}

func OrderTypes() []*OrderType {
	return []*OrderType{
		Market,
		Limit,
		Stop,
	}
}
