package gyroscope

import (
	"fmt"
	"v3/internal/domain"
)

func (r *gyroscopeRepository) Create(d *domain.Gyroscope) (*domain.Gyroscope, error) {
	if d.DeviceID == "" {
		return nil, ErrDeviceIDEmpty
	}
	if d.X == 0 {
		return nil, ErrXEmpty
	}
	if d.Y == 0 {
		return nil, ErrYEmpty
	}
	if d.Z == 0 {
		return nil, ErrZEmpty
	}
	if d.Timestamp.IsZero() {
		return nil, ErrTimestampEmpty
	}

	if err := r.DB.Create(d).Error; err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateFailed, err)
	}

	return d, nil
}
