package interactors

import (
	"cointrading/app/application/usecases/dto"
	"cointrading/app/application/usecases/getexchanges"
	"cointrading/app/domain/valueobjects"
)

type GetExchangesInteractor struct {
}

func NewGetExchangesUsecase() getexchanges.GetExchangesUsecase {
	return &GetExchangesInteractor{}
}

func (g *GetExchangesInteractor) Handle() *getexchanges.GetExchangesOutput {
	exchanges := valueobjects.Exchanges()

	output := &getexchanges.GetExchangesOutput{}

	for _, exchange := range exchanges {
		exchangeDTO := &dto.Exchange{
			Name:  exchange.DisplayValue(),
			Value: exchange.Value(),
		}

		output.Exchanges = append(output.Exchanges, exchangeDTO)
	}

	return output
}
