package interactors

import (
	"cointrading/app/application/usecases/initializetradingconfig"
	"cointrading/app/domain/entities"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/shared"
	"log"
)

type InitializeTradingConfigInteractor struct {
	tradingConfigRepository repositories.TradingConfigRepository
}

func NewInitializeTradingConfigUsecase(tradingConfigRepository repositories.TradingConfigRepository) initializetradingconfig.InitializeTradingConfigUsecase {
	return &InitializeTradingConfigInteractor{
		tradingConfigRepository: tradingConfigRepository,
	}
}

func (g *InitializeTradingConfigInteractor) Handle() {
	tradingConfig, err := g.tradingConfigRepository.Find()
	if err != nil {
		entity, err := entities.NewTradingConfig(shared.Exchange().Value(), shared.Symbol().Value(), shared.Duration().Value().String())
		if err != nil {
			log.Println(err)
			return
		}

		g.tradingConfigRepository.Create(entity)
	} else {
		shared.SetExchange(tradingConfig.Exchange())
		shared.SetSymbol(tradingConfig.Symbol())
		shared.SetDuration(tradingConfig.Duration())
	}

}
