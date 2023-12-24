package entities

import "cointrading/app/domain/valueobjects"

type Order struct {
	id          string
	symbol      *valueobjects.Symbol
	side        *valueobjects.Side
	orderType   *valueobjects.OrderType
	price       int
	size        float64
	timeInForce *valueobjects.TimeInForce
	status      *valueobjects.OrderStatus
}

func NewOrder(id string, symbol int, side string, orderType, price int, size float64, timeInForce, orderStatus int) (*Order, error) {
	sy, err := valueobjects.NewSymbol(symbol)
	if err != nil {
		return nil, err
	}

	si, err := valueobjects.NewSide(side)
	if err != nil {
		return nil, err
	}

	ot, err := valueobjects.NewOrderType(orderType)
	if err != nil {
		return nil, err
	}

	t, err := valueobjects.NewTimeInForce(timeInForce)
	if err != nil {
		return nil, err
	}

	os, err := valueobjects.NewOrderStatus(orderStatus)
	if err != nil {
		return nil, err
	}

	return &Order{
		id:          id,
		symbol:      sy,
		side:        si,
		orderType:   ot,
		price:       price,
		size:        size,
		timeInForce: t,
		status:      os,
	}, nil
}

func (o *Order) Id() string {
	return o.id
}

func (o *Order) Symbol() *valueobjects.Symbol {
	return o.symbol
}

func (o *Order) Side() *valueobjects.Side {
	return o.side
}

func (o *Order) OrderType() *valueobjects.OrderType {
	return o.orderType
}

func (o *Order) Price() int {
	return o.price
}

func (o *Order) Size() float64 {
	return o.size
}

func (o *Order) TimeInForce() *valueobjects.TimeInForce {
	return o.timeInForce
}

func (o *Order) Status() *valueobjects.OrderStatus {
	return o.status
}
