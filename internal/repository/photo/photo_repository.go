package photo

import (
	"errors"
	"fmt"
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

func (r *photoRepository) Create(d *domain.Photo) (*domain.Photo, error) {

	if d.DeviceID == "" {
		return nil, ErrDeviceIDEmpty
	}
	if d.FilePath == "" {
		return nil, ErrFilePathEmpty
	}
	if d.Timestamp.IsZero() {
		return nil, ErrTimestampEmpty
	}

	if err := r.DB.Create(d).Error; err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateFailed, err)
	}

	return d, nil
}
