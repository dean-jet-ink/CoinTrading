package interactors

import (
	"cointrading/app/application/usecases/savecandle"
	"cointrading/app/domain/entities"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/valueobjects"
	"time"
)

type SaveCandleInteractor struct {
	candleRepo repositories.CandleRepository
}

func NewSaveCandleUsecase(candleRepo repositories.CandleRepository) savecandle.SaveCandleUsecase {
	return &SaveCandleInteractor{
		candleRepo: candleRepo,
	}
}

func (s *SaveCandleInteractor) Handle(input *savecandle.SaveCandleInput) (*savecandle.SaveCandleOutput, error) {
	exchange := input.Exchange
	symbol := input.Ticker.Symbol()
	duration := input.Duration
	time := input.Ticker.Time()
	price := input.Ticker.MidPrice()
	volume := input.Ticker.Volume()

	candle := entities.NewCandle(exchange, symbol, duration, time, 0, 0, 0, 0, 0)

	candle, err := s.candleRepo.FindByTime(candle)
	if err != nil {
		return nil, err
	}

	output := &savecandle.SaveCandleOutput{
		IsCreated: false,
	}

	if candle == nil {
		if err = s.createCandle(exchange, symbol, duration, time, price, volume); err != nil {
			return nil, err
		}

		output.IsCreated = true

		return output, nil
	}

	if err = s.updateCandle(price, volume, candle); err != nil {
		return nil, err
	}

	return output, nil
}

func (s *SaveCandleInteractor) createCandle(exchange *valueobjects.Exchange, symbol *valueobjects.Symbol, duration *valueobjects.Duration, time time.Time, price, volume float64) error {
	candle := entities.NewCandle(
		exchange,
		symbol,
		duration,
		time,
		price,
		price,
		price,
		price,
		volume,
	)

	if err := s.candleRepo.Create(candle); err != nil {
		return err
	}

	return nil
}

func (s *SaveCandleInteractor) updateCandle(price, volume float64, candle *entities.Candle) error {
	high := candle.High()
	low := candle.Low()
	close := price

	switch {
	case candle.High() <= price:
		high = price
	case candle.Low() >= price:
		low = price
	}

	candle = entities.NewCandle(
		candle.Exchange(),
		candle.Symbol(),
		candle.Duration(),
		candle.Time().Value(),
		candle.Open(),
		close,
		high,
		low,
		volume,
	)

	if err := s.candleRepo.Update(candle); err != nil {
		return err
	}

	return nil
}
