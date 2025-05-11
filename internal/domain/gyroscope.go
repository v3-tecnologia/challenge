package domain

import (
	"errors"
	"time"

	"github.com/iamrosada0/v3/internal/adapter/uuid"
)

var (
	ErrDeviceIDGyroscope      = errors.New("device ID not found")
	ErrTimestampGyroscope     = errors.New("timestamp not found")
	ErrInvalidGyroscopeValues = errors.New("invalid gyroscope values")
)

type GyroscopeDto struct {
	DeviceID  string  `json:"deviceId" binding:"required"`
	Timestamp int64   `json:"timestamp" binding:"required"`
	X         float64 `json:"x" binding:"required"`
	Y         float64 `json:"y" binding:"required"`
	Z         float64 `json:"z" binding:"required"`
}

type Gyroscope struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	DeviceID  string    `json:"device_id" gorm:"index;not null"`
	X         float64   `json:"x" gorm:"not null"`
	Y         float64   `json:"y" gorm:"not null"`
	Z         float64   `json:"z" gorm:"not null"`
	Timestamp time.Time `json:"timestamp" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func NewGyroscopeData(d *GyroscopeDto) (*Gyroscope, error) {

	id := uuid.NewAdapter().Generate()

	dev, err := NewDevice(d.DeviceID)
	if err != nil {
		return nil, ErrDeviceIDGyroscope
	}

	timestamp := time.Unix(d.Timestamp, 0)
	if timestamp.IsZero() {
		return nil, ErrTimestampGyroscope
	}

}
