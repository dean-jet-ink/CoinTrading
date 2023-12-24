package repositories

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/valueobjects"
)

type TradingAPIClient interface {
	GetBalance() ([]*entities.Balance, error)
	GetTicker(symbol *valueobjects.Symbol) (*entities.Ticker, error)
	GetRealTimeTicker(symbol *valueobjects.Symbol, tickerChan chan<- *entities.Ticker)
	SendOrder(order *entities.Order) (string, error)
	GetOrder(orderId string) (*entities.Order, error)
	GetOrders(symbol *valueobjects.Symbol) ([]*entities.Order, error)
}
