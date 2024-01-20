package interactors

import (
	"cointrading/app/application/usecases/getrealtimeticker"
	"cointrading/app/domain/repositories"
	"cointrading/app/infrastructure/factories/tradingfactory"
)

type getRealTimeTickerInteractor struct {
	candleRepo repositories.CandleRepository
}

func NewGetRealTimeTickerUsecase(candleRepo repositories.CandleRepository) getrealtimeticker.GetRealTimeTickerUsecase {
	return &getRealTimeTickerInteractor{
		candleRepo: candleRepo,
	}
}

func (g *getRealTimeTickerInteractor) Handle(input *getrealtimeticker.GetRealTimeTickerInput) {
	tradingFactory := tradingfactory.NewTradingFactory(input.Exchange)

	client := tradingFactory.NewTradingAPIClient()

	client.GetRealTimeTicker(input.CTX, input.Symbol, input.TickerChan, input.ErrChan)
}
