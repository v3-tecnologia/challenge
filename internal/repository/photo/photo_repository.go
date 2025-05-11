package photo

import (
	"github.com/iamrosada0/v3/internal/domain"
	"gorm.io/gorm"
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
