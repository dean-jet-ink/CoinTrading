package dbfactory

import (
	"cointrading/app/config"
	"cointrading/app/domain/factories"
	"cointrading/app/domain/repositories"
	mygorm "cointrading/app/infrastructure/gorm"

	"gorm.io/gorm"
)

type gormFactory struct {
	db *gorm.DB
}

func newGORMFactory() factories.DBFactory {
	var db *gorm.DB

	switch config.DB() {
	case 1:
		db = mygorm.SQLiteConnection()
	}

	return &gormFactory{
		db: db,
	}
}

func (g *gormFactory) DB() *gorm.DB {
	return g.db
}

func (g *gormFactory) NewCandleRepository() repositories.CandleRepository {
	return mygorm.NewCandleRepository(g.db)
}

func (g *gormFactory) NewOrderRepository() repositories.OrderRepository {
	return mygorm.NewOrderRepository(g.db)
}

func (g *gormFactory) NewTradingConfigRepository() repositories.TradingConfigRepository {
	return mygorm.NewTradingConfigRepository(g.db)
}
