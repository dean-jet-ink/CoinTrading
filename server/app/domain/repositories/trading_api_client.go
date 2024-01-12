package repositories

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/valueobjects"
	"context"
)

type TradingAPIClient interface {
	GetBalance() ([]*entities.Balance, error)
	GetTicker(symbol *valueobjects.Symbol) (*entities.Ticker, error)
	GetRealTimeTicker(ctx context.Context, symbol *valueobjects.Symbol, tickerChan chan<- *entities.Ticker, errChan chan<- error)
	SendOrder(order *entities.Order) (string, error)
	GetOrder(orderId string) (*entities.Order, error)
	GetOrders(symbol *valueobjects.Symbol) ([]*entities.Order, error)
}
