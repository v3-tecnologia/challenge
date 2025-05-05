package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID         uuid.UUID // Unique identifier for the photo entry
	DeviceID   string    // Unique identifier of the device (MAC address)
	FilePath   string    // Path where the image is stored
	Recognized bool      // Indicates if the photo was matched using Rekognition
	Timestamp  time.Time // Timestamp when the photo was captured
	CreatedAt  time.Time // Timestamp when the entry was saved in the system
}

func NewPhoto(deviceID, filePath string, timestamp time.Time, recognized bool) (*Photo, error) {
	if deviceID == "" {
		return nil, errors.New("deviceID cannot be empty")
	}
	if filePath == "" {
		return nil, errors.New("filePath cannot be empty")
	}
	if timestamp.IsZero() {
		return nil, errors.New("timestamp is required")
	}

	return &Photo{
		ID:         uuid.New(),
		DeviceID:   deviceID,
		FilePath:   filePath,
		Timestamp:  timestamp,
		Recognized: recognized,
		CreatedAt:  time.Now(),
	}, nil
}
