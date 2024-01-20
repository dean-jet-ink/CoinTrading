package repositories

import "cointrading/app/domain/entities"

type TradingConfigRepository interface {
	Find() (*entities.TradingConfig, error)
	Create(tradingConfig *entities.TradingConfig) error
	Update(tradingConfig *entities.TradingConfig) error
}
