package factories

import (
	"cointrading/app/domain/repositories"
)

type DBFactory interface {
	NewCandleRepository() repositories.CandleRepository
	NewOrderRepository() repositories.OrderRepository
	NewTradingConfigRepository() repositories.TradingConfigRepository
}
