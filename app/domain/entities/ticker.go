package entities

import (
	"cointrading/app/domain/valueobjects"
	"time"
)

type Ticker struct {
	symbol   *valueobjects.Symbol
	dateTime *valueobjects.DateTime
	bestAsk  float64
	bestBid  float64
	volume   float64
}

func NewTicker(symbol *valueobjects.Symbol, dateTime time.Time, bestAsk, bestBid float64, volume float64) *Ticker {
	return &Ticker{
		symbol:   symbol,
		dateTime: valueobjects.NewDateTime(dateTime),
		bestAsk:  bestAsk,
		bestBid:  bestBid,
		volume:   volume,
	}
}

func (t *Ticker) MidPrice() float64 {
	return (t.bestAsk + t.bestBid) / 2
}
