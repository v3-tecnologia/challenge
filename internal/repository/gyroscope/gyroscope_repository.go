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
