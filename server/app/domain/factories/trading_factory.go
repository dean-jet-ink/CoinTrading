package factories

import (
	"cointrading/app/domain/repositories"
)

type ITradingFactory interface {
	NewTradingAPIClient() repositories.TradingAPIClient
}
