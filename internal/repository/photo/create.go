package photo

import (
	"fmt"
	"v3/internal/domain"
)

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
