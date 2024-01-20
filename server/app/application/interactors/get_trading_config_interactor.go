package interactors

import (
	"cointrading/app/application/usecases/dto"
	"cointrading/app/application/usecases/gettradingconfig"
	"cointrading/app/domain/shared"
	"cointrading/app/domain/valueobjects"
)

type GetTradingConfigInteractor struct {
}

func NewGetTradingConfigUsecase() gettradingconfig.GetTradingConfigUsecase {
	return &GetTradingConfigInteractor{}
}

func (g *GetTradingConfigInteractor) Handle() *gettradingconfig.GetTradingConfigOutput {
	exchange := shared.Exchange()
	symbol := shared.Symbol()
	duration := shared.Duration()

	return g.createOutput(exchange, symbol, duration)
}

func (g *GetTradingConfigInteractor) createOutput(exchange *valueobjects.Exchange, symbol *valueobjects.Symbol, duration *valueobjects.Duration) *gettradingconfig.GetTradingConfigOutput {
	output := &gettradingconfig.GetTradingConfigOutput{
		Exchange: &dto.Exchange{
			Name:  exchange.DisplayValue(),
			Value: exchange.Value(),
		},
		Symbol: &dto.Symbol{
			Name:  symbol.DisplayValue(),
			Value: symbol.Value(),
		},
		Duration: &dto.Duration{
			Value:        duration.Value().String(),
			DisplayValue: duration.DisplayValue(),
		},
	}

	return output
}
