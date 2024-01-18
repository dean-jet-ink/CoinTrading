package gorm

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/repositories"
	"cointrading/app/domain/valueobjects"
	"cointrading/app/infrastructure/gorm/models"

	"gorm.io/gorm"
)

type gormCandle struct {
	db *gorm.DB
}

func NewCandleRepository(db *gorm.DB) repositories.CandleRepository {
	return &gormCandle{
		db: db,
	}
}

func (g *gormCandle) FindByTime(candle *entities.Candle) (*entities.Candle, error) {
	model := &models.Candle{}

	if err := g.db.Table(candle.GetTableName()).Where("time = ?", candle.TruncatedDateTime()).Find(model).Error; err != nil {
		return nil, err
	}

	if model.Time.IsZero() {
		return nil, nil
	}

	entity := g.modelToEntity(candle.Exchange(), candle.Symbol(), candle.Duration(), model)

	return entity, nil
}

func (g *gormCandle) FindAllWithLimit(candle *entities.Candle, limit int) ([]*entities.Candle, error) {
	models := []*models.Candle{}

	subQuery := g.db.Table(candle.GetTableName()).Order("time DESC").Limit(limit)

	err := g.db.Table("(?) as c", subQuery).Order("time ASC").Find(&models).Error
	if err != nil {
		return nil, err
	}

	entities := []*entities.Candle{}

	for _, model := range models {
		entity := g.modelToEntity(candle.Exchange(), candle.Symbol(), candle.Duration(), model)

		entities = append(entities, entity)
	}

	return entities, nil
}

func (g *gormCandle) Create(candle *entities.Candle) error {
	model := g.entityToModel(candle)

	if err := g.db.Table(candle.GetTableName()).Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (g *gormCandle) Update(candle *entities.Candle) error {
	model := g.entityToModel(candle)

	if err := g.db.Table(candle.GetTableName()).Updates(model).Error; err != nil {
		return err
	}

	return nil
}

func (g *gormCandle) entityToModel(entity *entities.Candle) *models.Candle {
	return &models.Candle{
		Time:   entity.TruncatedDateTime(),
		Open:   entity.Open(),
		Close:  entity.Close(),
		High:   entity.High(),
		Low:    entity.Low(),
		Volume: entity.Volume(),
	}
}

func (g *gormCandle) modelToEntity(exchange *valueobjects.Exchange, symbol *valueobjects.Symbol, duration *valueobjects.Duration, model *models.Candle) *entities.Candle {
	entity := entities.NewCandle(
		exchange,
		symbol,
		duration,
		model.Time,
		model.Open,
		model.Close,
		model.High,
		model.Low,
		model.Volume,
	)

	return entity
}
