package entities

import (
	"cointrading/app/domain/valueobjects"
	"time"
)

type TradingConfig struct {
	exchange *valueobjects.Exchange
	symbol   *valueobjects.Symbol
	duration *valueobjects.Duration
}

func NewTradingConfig(exchange int, symbol int, duration string) (*TradingConfig, error) {
	e, err := valueobjects.NewExchange(exchange)
	if err != nil {
		return nil, err
	}

	s, err := valueobjects.NewSymbol(symbol)
	if err != nil {
		return nil, err
	}

	parsedDuration, err := time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}

	d, err := valueobjects.NewDuration(parsedDuration)
	if err != nil {
		return nil, err
	}

	return &TradingConfig{
		exchange: e,
		symbol:   s,
		duration: d,
	}, nil
}

func (t *TradingConfig) Exchange() *valueobjects.Exchange {
	return t.exchange
}

func (t *TradingConfig) Symbol() *valueobjects.Symbol {
	return t.symbol
}

func (t *TradingConfig) Duration() *valueobjects.Duration {
	return t.duration
}
