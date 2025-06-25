package usecase

import (
	"errors"

	"github.com/yanvic/challenge/core/entity"
)

func ValidateGyroscope(data entity.Gyroscope) error {
	if data.DeviceID == "" {
		return errors.New("device_id is required")
	}
	if data.Timestamp == "" {
		return errors.New("timestamp is required")
	}
	if data.X == nil {
		return errors.New("x is required")
	}
	if data.Y == nil {
		return errors.New("y is required")
	}
	if data.Z == nil {
		return errors.New("z is required")
	}
	return nil
}
