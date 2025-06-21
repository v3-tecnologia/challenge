package usecase

import (
	"errors"

	"github.com/yanvic/challenge/core/entity"
)

func ValidatePhoto(data entity.Photo) error {
	if data.DeviceID == "" {
		return errors.New("device_id is required")
	}
	if data.Timestamp == "" {
		return errors.New("timestamp is required")
	}
	if data.ImageBase64 == "" {
		return errors.New("image_base64 is required")
	}
	return nil
}
