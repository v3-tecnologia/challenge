package gyroscope

import "github.com/iamrosada0/v3/internal/domain"

type GyroscopeRepository interface {
	Create(d *domain.Gyroscope) (*domain.Gyroscope, error)
}
