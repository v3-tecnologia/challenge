package repositories

import (
	"github.com/wellmtx/challenge/internal/infra/database"
	"github.com/wellmtx/challenge/internal/infra/entities"
)

type GyroscopeRepository interface {
	Create(data entities.GyroscopeEntity) (entities.GyroscopeEntity, error)
}

type gyroscopeRepository struct {
	db *database.Database
}

func NewGyroscopeRepository(db *database.Database) GyroscopeRepository {
	return &gyroscopeRepository{
		db: db,
	}
}

func (r *gyroscopeRepository) Create(data entities.GyroscopeEntity) (entities.GyroscopeEntity, error) {
	if err := r.db.DB.Create(&data).Error; err != nil {
		return entities.GyroscopeEntity{}, err
	}
	return data, nil
}
