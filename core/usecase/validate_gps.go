package usecase

import (
	"errors"

	"github.com/yanvic/challenge/core/entity"
)

func ValidateGPS(data entity.GPS) error {
	if data.DeviceID == "" {
		return errors.New("device_id is required")
	}
	if data.Timestamp == "" {
		return errors.New("timestamp is required")
	}
	if data.Longitude == nil {
		return errors.New("latitude is required")
	}
	if data.Longitude == nil {
		return errors.New("longitude is required")
	}
	return nil
}
