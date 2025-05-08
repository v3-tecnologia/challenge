package gyroscope

import (
	domain "github.com/iamrosada0/v3/internal/domain/gyroscope"
)

type GyroscopeRepository interface {
	Create(d *domain.GyroscopeData) (*domain.GyroscopeData, error)
}
