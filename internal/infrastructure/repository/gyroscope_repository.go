package repository

import (
	"challenge-v3-backend/internal/domain/entity"
	"context"
	"gorm.io/gorm"
)

type GyroscopeTelemetryRepository struct {
	db *gorm.DB
}

func NewGyroscopeTelemetryRepository(db *gorm.DB) *GyroscopeTelemetryRepository {
	return &GyroscopeTelemetryRepository{
		db: db,
	}

}

func (r *GyroscopeTelemetryRepository) CreateGyroscopeTelemetry(ctx context.Context, entity *entity.GyroscopeTelemetry) error {
	result := r.db.WithContext(ctx).Create(entity)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
