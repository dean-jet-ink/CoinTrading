package interactors

import (
	"cointrading/app/application/usecases/dto"
	"cointrading/app/application/usecases/getdurations"
	"cointrading/app/domain/valueobjects"
)

type GetDurationsInteractor struct {
}

func NewGetDurationsUsecase() getdurations.GetDurationsUsecase {
	return &GetDurationsInteractor{}
}

func (g *GetDurationsInteractor) Handle() *getdurations.GetDurationsOutput {
	durations := valueobjects.Durations()

	output := &getdurations.GetDurationsOutput{}

	for _, duration := range durations {
		durationDTO := &dto.Duration{
			Value:        duration.Value().String(),
			DisplayValue: duration.DisplayValue(),
		}

		output.Durations = append(output.Durations, durationDTO)
	}

	return output
}
