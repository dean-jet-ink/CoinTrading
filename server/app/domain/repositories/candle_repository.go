package repositories

import (
	"cointrading/app/domain/entities"
)

type CandleRepository interface {
	FindByTime(candle *entities.Candle) (*entities.Candle, error)
	FindAllWithLimit(candle *entities.Candle, limit int) ([]*entities.Candle, error)
	Create(candle *entities.Candle) error
	Update(candle *entities.Candle) error
}
