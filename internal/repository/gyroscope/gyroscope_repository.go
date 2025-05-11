package gyroscope

import (
	"github.com/iamrosada0/v3/internal/domain"
	"gorm.io/gorm"
)

type GyroscopeRepository interface {
	Create(d *domain.Gyroscope) (*domain.Gyroscope, error)
}
type gyroscopeRepository struct {
	DB *gorm.DB
}

func NewGyroscopeRepository(db *gorm.DB) GyroscopeRepository {
	return &gyroscopeRepository{DB: db}
}
func (r *gyroscopeRepository) Create(d *domain.Gyroscope) (*domain.Gyroscope, error) {
	if err := r.DB.Create(d).Error; err != nil {
		return nil, err
	}
	return d, nil
}
