package gorm

import (
	"cointrading/app/domain/repositories"

	"gorm.io/gorm"
)

type sqliteCandle struct {
	db *gorm.DB
}

func NewCandleRepository(db *gorm.DB) repositories.CandleRepository {
	return &sqliteCandle{
		db: db,
	}
}

func (s *sqliteCandle) Create() {

}
