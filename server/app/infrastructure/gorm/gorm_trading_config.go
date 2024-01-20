package gorm

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/myerror"
	"cointrading/app/domain/repositories"
	"cointrading/app/infrastructure/gorm/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type GormTradingConfig struct {
	db *gorm.DB
}

func NewTradingConfigRepository(db *gorm.DB) repositories.TradingConfigRepository {
	return &GormTradingConfig{db}
}

func (g *GormTradingConfig) Find() (*entities.TradingConfig, error) {
	model := &models.TradingConfig{}

	if err := g.db.First(model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w: Record not found: %v", myerror.ErrRecordNotFound, err)
		}
		return nil, err
	}

	return g.modelToEntity(model)
}

func (g *GormTradingConfig) Create(tradingConfig *entities.TradingConfig) error {
	model := g.entityToModel(tradingConfig)

	if err := g.db.Model(&models.TradingConfig{}).Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (g *GormTradingConfig) Update(tradingConfig *entities.TradingConfig) error {
	model := g.entityToModel(tradingConfig)

	if err := g.db.Model(&models.TradingConfig{}).Updates(model).Error; err != nil {
		return err
	}

	return nil
}

func (g *GormTradingConfig) modelToEntity(model *models.TradingConfig) (*entities.TradingConfig, error) {
	return entities.NewTradingConfig(model.Exchange, model.Symbol, model.Duration)
}

func (g *GormTradingConfig) entityToModel(tradingConfig *entities.TradingConfig) *models.TradingConfig {
	return &models.TradingConfig{
		Exchange: tradingConfig.Exchange().Value(),
		Symbol:   tradingConfig.Symbol().Value(),
		Duration: tradingConfig.Duration().Value().String(),
	}
}
