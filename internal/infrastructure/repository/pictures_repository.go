package repository

import (
	"challenge-v3-backend/internal/domain/entity"
	"context"
	"gorm.io/gorm"
)

type PicturesRepository struct {
	db *gorm.DB
}

func NewPicturesRepository(db *gorm.DB) *PicturesRepository {
	return &PicturesRepository{
		db: db,
	}

}

func (r *PicturesRepository) CreatePictures(ctx context.Context, entity *entity.Picture) (*entity.Picture, error) {
	result := r.db.WithContext(ctx).Create(entity)

	if result.Error != nil {
		return nil, result.Error
	}

	return entity, nil
}
