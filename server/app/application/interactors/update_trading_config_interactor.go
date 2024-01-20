package interactors

import (
	"cointrading/app/application/usecases/updatetradingconfig"
	"cointrading/app/domain/entities"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/shared"
	"cointrading/app/domain/valueobjects"
	"time"
)

type UpdateTradingConfigInteractor struct {
	tradingConfigRepository repositories.TradingConfigRepository
}

func NewUpdateTradingConfigUsecase(tradingConfigRepository repositories.TradingConfigRepository) updatetradingconfig.UpdateTradingConfigUsecase {
	return &UpdateTradingConfigInteractor{
		tradingConfigRepository: tradingConfigRepository,
	}
}

func (u *UpdateTradingConfigInteractor) Handle(input *updatetradingconfig.UpdateTradingConfigInput) error {
	if input.Exchange != 0 {
		exchange, err := valueobjects.NewExchange(input.Exchange)
		if err != nil {
			return err
		}

		shared.SetExchange(exchange)
	} else if input.Symbol != 0 {
		symbol, err := valueobjects.NewSymbol(input.Symbol)
		if err != nil {
			return err
		}

		shared.SetSymbol(symbol)
	} else if input.Duration != "" {
		parsedDuration, err := time.ParseDuration(input.Duration)
		if err != nil {
			return err
		}

		duration, err := valueobjects.NewDuration(parsedDuration)
		if err != nil {
			return err
		}

		shared.SetDuration(duration)
	}

	entity, err := entities.NewTradingConfig(input.Exchange, input.Symbol, input.Duration)
	if err != nil {
		return err
	}

	return u.tradingConfigRepository.Update(entity)
}
