package repositories

import "cointrading/app/domain/entities"

type TradingAPIClient interface {
	GetBalance() ([]*entities.Balance, error)
}
