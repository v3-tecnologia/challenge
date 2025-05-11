package gps

import "github.com/iamrosada0/v3/internal/domain"

type GPSRepository interface {
	Create(d *domain.GPS) (*domain.GPS, error)
}
