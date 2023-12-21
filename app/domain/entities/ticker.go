package entities

import (
	"cointrading/app/domain/valueobjects"
	"time"
)

type Ticker struct {
	symbol   *valueobjects.Symbol
	dateTime *valueobjects.DateTime
	price    float64
	volume   float64
}

func NewTicker(symbol *valueobjects.Symbol, dateTime time.Time, price float64, volume float64) *Ticker {
	return &Ticker{
		symbol:   symbol,
		dateTime: valueobjects.NewDateTime(dateTime),
		price:    price,
		volume:   volume,
	}
}
