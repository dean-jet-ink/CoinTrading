package dbfactory

import (
	"cointrading/app/config"
	"cointrading/app/domain/factories"
)

func NewDBFactory() factories.DBFactory {
	db := config.DB()

	switch db {
	case 1:
		return newGORMFactory()
	default:
		return newGORMFactory()
	}
}
