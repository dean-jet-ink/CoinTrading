package entities

import (
	"cointrading/app/domain/valueobject"
	"time"
)

type Ticker struct {
	symbol   *valueobject.Symbol
	dateTime *valueobject.DateTime
	price    float64
	volume   float64
}

func NewTicker(symbol int, dateTime time.Time, price float64, volume float64) *Ticker {
	return &Ticker{
		symbol:   valueobject.NewSymbol(symbol),
		dateTime: valueobject.NewDateTime(dateTime),
		price:    price,
		volume:   volume,
	}
}
