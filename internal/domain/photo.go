package domain

import (
	"errors"
	"time"

	"github.com/iamrosada0/v3/internal/adapter/uuid"
)

var (
	ErrDeviceIDPhoto  = errors.New("device ID not found")
	ErrTimestampPhoto = errors.New("timestamp not found")
	ErrFilePathPhoto  = errors.New("file path not found")
)

type PhotoDto struct {
	DeviceID  string `json:"deviceId" binding:"required"`
	Timestamp int64  `json:"timestamp" binding:"required"`
	FilePath  string `json:"file_path" binding:"required"`
}

type Photo struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	DeviceID   string    `json:"device_id" gorm:"index;not null"`
	FilePath   string    `json:"file_path" gorm:"not null"`
	Recognized bool      `json:"recognized" gorm:"default:false"`
	Timestamp  time.Time `json:"timestamp" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func NewPhotoData(d *PhotoDto) (*Photo, error) {
	id := uuid.NewAdapter().Generate()

	dev, err := NewDevice(d.DeviceID)
	if err != nil {
		return nil, ErrDeviceIDPhoto
	}

	timestamp := time.Unix(d.Timestamp, 0)
	if timestamp.IsZero() {
		return nil, ErrTimestampPhoto
	}

	if d.FilePath == "" {
		return nil, ErrFilePathPhoto
	}
}
