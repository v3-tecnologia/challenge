package photo

import (
	"errors"
	"v3/internal/domain"

	"gorm.io/gorm"
)

var (
	ErrDeviceIDEmpty  = errors.New("NOT NULL constraint failed: photos.device_id")
	ErrFilePathEmpty  = errors.New("NOT NULL constraint failed: photos.file_path")
	ErrTimestampEmpty = errors.New("NOT NULL constraint failed: photos.timestamp")
	ErrCreateFailed   = errors.New("failed to create photo record")
)

type PhotoRepository interface {
	Create(d *domain.Photo) (*domain.Photo, error)
}

type photoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{DB: db}
}
