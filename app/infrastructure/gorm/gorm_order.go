package gorm

import (
	"cointrading/app/domain/repositories"

	"gorm.io/gorm"
)

type sqliteOrder struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repositories.OrderRepository {
	return &sqliteOrder{
		db: db,
	}
}

func (s *sqliteOrder) Create() {

}

func (s *sqliteOrder) Update() {

}

func (s *sqliteOrder) FindByID() {

}

func (s *sqliteOrder) FindOrdersBySymbol() {

}
