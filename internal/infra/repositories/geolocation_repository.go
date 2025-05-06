package repositories

import (
	"github.com/wellmtx/challenge/internal/infra/database"
	"github.com/wellmtx/challenge/internal/infra/entities"
)

type GeolocationRepository interface {
	Create(data entities.GeolocationEntity) (entities.GeolocationEntity, error)
}

type geolocationRepository struct {
	db *database.Database
}

func NewGeolocationRepository(db *database.Database) GeolocationRepository {
	return &geolocationRepository{
		db: db,
	}
}

func (r *geolocationRepository) Create(data entities.GeolocationEntity) (entities.GeolocationEntity, error) {
	if err := r.db.DB.Create(&data).Error; err != nil {
		return entities.GeolocationEntity{}, err
	}
	return data, nil
}
