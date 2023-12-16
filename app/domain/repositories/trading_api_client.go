package repositories

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/valueobject"
)

type TradingAPIClient interface {
	GetBalance() ([]*entities.Balance, error)
	GetTicker(symbol *valueobject.Symbol) (*entities.Ticker, error)
	GetRealTimeTicker(symbol *valueobject.Symbol, tickerChan chan<- *entities.Ticker) error
}
