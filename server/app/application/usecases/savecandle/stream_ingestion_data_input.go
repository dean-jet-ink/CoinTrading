package savecandle

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/valueobjects"
)

type SaveCandleInput struct {
	Exchange *valueobjects.Exchange
	Duration *valueobjects.Duration
	Ticker   *entities.Ticker
}
