package interactors

import (
	"cointrading/app/application/usecases/dto"
	"cointrading/app/application/usecases/getsymbols"
	"cointrading/app/domain/valueobjects"
)

type GetSymbolsInteractor struct {
	getSymbolsUsecase getsymbols.GetSymbolsUsecase
}

func NewGetSymbolsUsecase() getsymbols.GetSymbolsUsecase {
	return &GetSymbolsInteractor{}
}

func (g *GetSymbolsInteractor) Handle() *getsymbols.GetSymbolsOutput {
	symbols := valueobjects.Symbols()

	output := &getsymbols.GetSymbolsOutput{}

	for _, symbol := range symbols {
		symbolDTO := &dto.Symbol{
			Name:  symbol.DisplayValue(),
			Value: symbol.Value(),
		}

		output.Symbols = append(output.Symbols, symbolDTO)
	}

	return output
}
