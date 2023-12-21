package repositories

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/valueobjects"
)

type TradingAPIClient interface {
	GetBalance() ([]*entities.Balance, error)
	GetTicker(symbol *valueobjects.Symbol) (*entities.Ticker, error)
	GetRealTimeTicker(symbol *valueobjects.Symbol, tickerChan chan<- *entities.Ticker)
}
