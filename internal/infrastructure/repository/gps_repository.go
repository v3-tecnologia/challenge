package repository

import (
	"challenge-v3-backend/internal/domain/entity"
	"context"
	"gorm.io/gorm"
)

type GPSTelemetryRepository struct {
	db *gorm.DB
}

func NewGPSRepository(db *gorm.DB) *GPSTelemetryRepository {
	return &GPSTelemetryRepository{
		db: db,
	}

}

func (r *GPSTelemetryRepository) CreateGPSTelemetry(ctx context.Context, entity *entity.GPSTelemetry) error {
	result := r.db.WithContext(ctx).Create(entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
