package factories

import (
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/valueobjects"
)

type ITradingFactory interface {
	CreateTradingAPIClient(exchenge *valueobjects.Exchange) repositories.TradingAPIClient
}
