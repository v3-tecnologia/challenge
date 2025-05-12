package gps

import (
	"fmt"
	"v3/internal/domain"
)

func (r *gpsRepository) Create(d *domain.GPS) (*domain.GPS, error) {

	if d.DeviceID == "" {
		return nil, ErrDeviceIDEmpty
	}
	if d.Timestamp.IsZero() {
		return nil, ErrTimestampEmpty
	}
	if d.Latitude == 0 {
		return nil, ErrLatitudeEmpty
	}
	if d.Longitude == 0 {
		return nil, ErrLongitudeEmpty
	}

	if err := r.DB.Create(d).Error; err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateFailed, err)
	}

	return d, nil
}
