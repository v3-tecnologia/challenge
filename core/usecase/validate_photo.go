package usecase

import (
	"bytes"
	"errors"
	"image"

	"github.com/yanvic/challenge/core/entity"
)

func ValidatePhoto(data entity.Photo) error {
	if data.DeviceID == "" {
		return errors.New("device_id is required")
	}
	if data.Timestamp == "" {
		return errors.New("timestamp is required")
	}
	if len(data.Image) == 0 {
		return errors.New("image data is required")
	}

	_, _, err := image.Decode(bytes.NewReader(data.Image))
	if err != nil {
		return errors.New("image data is not a valid image format")
	}

	return nil
}
