package domain

import (
	"errors"
	"time"

	"github.com/iamrosada0/v3/internal/adapter/uuid"
)

var (
	ErrDeviceIDGPS      = errors.New("device ID not found")
	ErrTimestampGPS     = errors.New("timestamp not found")
	ErrInvalidGPSValues = errors.New("invalid GPS values")
)

type GPSDto struct {
	DeviceID  string  `json:"deviceId" binding:"required"`
	Timestamp int64   `json:"timestamp" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type GPS struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	DeviceID  string    `json:"device_id" gorm:"index;not null"`
	Latitude  float64   `json:"latitude" gorm:"not null"`
	Longitude float64   `json:"longitude" gorm:"not null"`
	Timestamp time.Time `json:"timestamp" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func NewGPSData(d *GPSDto) (*GPS, error) {
	id := uuid.NewAdapter().Generate()

	dev, err := NewDevice(d.DeviceID)
	if err != nil {
		return nil, ErrDeviceIDGPS
	}

	timestamp := time.Unix(d.Timestamp, 0)
	if timestamp.IsZero() {
		return nil, ErrTimestampGPS
	}
}
