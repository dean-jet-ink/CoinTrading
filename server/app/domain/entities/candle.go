package entities

import (
	"cointrading/app/domain/valueobjects"
	"fmt"
	"time"
)

type Candle struct {
	exchange *valueobjects.Exchange
	symbol   *valueobjects.Symbol
	duration *valueobjects.Duration
	time     *valueobjects.DateTime
	open     float64
	close    float64
	high     float64
	low      float64
	volume   float64
}

func NewCandle(exchange *valueobjects.Exchange, symbol *valueobjects.Symbol, duration *valueobjects.Duration, time time.Time, open float64, close float64, high float64, low float64, volume float64) *Candle {
	return &Candle{
		exchange: exchange,
		symbol:   symbol,
		duration: duration,
		time:     valueobjects.NewDateTime(time),
		open:     open,
		close:    close,
		high:     high,
		low:      low,
		volume:   volume,
	}
}

func (c *Candle) Exchange() *valueobjects.Exchange {
	return c.exchange
}

func (c *Candle) Symbol() *valueobjects.Symbol {
	return c.symbol
}

func (c *Candle) Duration() *valueobjects.Duration {
	return c.duration
}

func (c *Candle) Time() *valueobjects.DateTime {
	return c.time
}

func (c *Candle) Open() float64 {
	return c.open
}

func (c *Candle) Close() float64 {
	return c.close
}

func (c *Candle) High() float64 {
	return c.high
}

func (c *Candle) Low() float64 {
	return c.low
}

func (c *Candle) Volume() float64 {
	return c.volume
}

func (c *Candle) TruncatedDateTime() time.Time {
	return c.time.TruncateDateTime(c.duration.Value())
}

func (c *Candle) DisplayTruncatedDateTime() string {
	layout := "2006/01/02 15:04:05"

	switch {
	case c.duration.IsSecond():
		layout = "2006/01/02 15:04:05"
	case c.duration.IsMinute():
		layout = "2006/01/02 15:04"
	case c.duration.IsHour():
		layout = "2006/01/02 15:00"
	case c.duration.IsDay() || c.duration.IsWeek():
		layout = "2006/01/02"
	case c.duration.IsMonth():
		layout = "2006/01"
	}

	return c.time.Value().Format(layout)
}

func (c *Candle) GetTableName() string {
	exchange := c.exchange.DisplayValueForTableName()
	symbol := c.symbol.DisplayValueForTableName()
	duration := c.duration.DisplayValueForTableName()

	return fmt.Sprintf("%s_%s_%s", exchange, symbol, duration)
}
