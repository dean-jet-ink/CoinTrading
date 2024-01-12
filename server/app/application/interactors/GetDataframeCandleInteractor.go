package interactors

import (
	"cointrading/app/application/usecases/dto"
	"cointrading/app/application/usecases/getdataframecandle"
	"cointrading/app/config"
	"cointrading/app/domain/entities"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/shared"
)

type GetDataframeCandleInteractor struct {
	candleRepo repositories.CandleRepository
}

func NewGetDataframeCandleUsecase(candleRepo repositories.CandleRepository) getdataframecandle.GetDataframeCandleUsecase {
	return &GetDataframeCandleInteractor{
		candleRepo: candleRepo,
	}
}

func (g *GetDataframeCandleInteractor) Handle(input *getdataframecandle.GetDataframeCandleInput) (*getdataframecandle.GetDataframeCandleOutput, error) {
	candle := entities.NewCandle(shared.Exchange(), shared.Symbol(), shared.Duration(), nil, 0, 0, 0, 0, 0)

	candles, err := g.candleRepo.FindAllWithLimit(candle, config.GetCandleLimit())
	if err != nil {
		return nil, err
	}

	candleDTOs := make([]*dto.Candle, len(candles))

	for i, c := range candles {
		candleDTO := g.candleEntityToDTO(c)

		candleDTOs[i] = candleDTO
	}

	return &getdataframecandle.GetDataframeCandleOutput{
		Candles: candleDTOs,
	}, nil
}

func (g *GetDataframeCandleInteractor) candleEntityToDTO(candle *entities.Candle) *dto.Candle {
	return &dto.Candle{
		Time:   candle.DisplayTruncatedDateTime(),
		Open:   candle.Open(),
		High:   candle.High(),
		Low:    candle.Low(),
		Close:  candle.Close(),
		Volume: candle.Volume(),
	}
}
