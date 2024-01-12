package getrealtimeticker

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/valueobjects"
	"context"
)

type GetRealTimeTickerInput struct {
	CTX        context.Context
	TickerChan chan *entities.Ticker
	ErrChan    chan error
	Exchange   *valueobjects.Exchange
	Symbol     *valueobjects.Symbol
}
