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

func NewTicker(symbol int, dateTime time.Time, bestAsk, bestBid float64, volume float64) (*Ticker, error) {
	s, err := valueobjects.NewSymbol(symbol)
	if err != nil {
		return nil, err
	}

	return &Ticker{
		symbol:   s,
		dateTime: valueobjects.NewDateTime(dateTime),
		bestAsk:  bestAsk,
		bestBid:  bestBid,
		volume:   volume,
	}, nil
}

func (t *Ticker) Symbol() *valueobjects.Symbol {
	return t.symbol
}

func (t *Ticker) DateTime() *valueobjects.DateTime {
	return t.dateTime
}

func (t *Ticker) BestAsk() float64 {
	return t.bestAsk
}

func (t *Ticker) BestBid() float64 {
	return t.bestBid
}

func (t *Ticker) Volume() float64 {
	return t.volume
}

func (t *Ticker) MidPrice() float64 {
	return (t.bestAsk + t.bestBid) / 2
}