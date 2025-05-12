package domain

import (
	"errors"
	"time"
	"v3/internal/adapter/uuid"
)

var (
	ErrDeviceIDPhoto                  = errors.New("device ID not found")
	ErrTimestampPhoto                 = errors.New("timestamp not found")
	ErrPhotoData                      = errors.New("photo data not found")
	ErrProcessPhotoWithAWSRekognition = errors.New("failed to process photo with AWS Rekognition")
	ErrSavePhotoData                  = errors.New("failed to save photo data")
	ErrMissingPhotoInvalidFields      = errors.New("missing or invalid photo fields")
)

type PhotoDto struct {
	DeviceID  string `form:"deviceId" binding:"required"`
	Timestamp int64  `form:"timestamp" binding:"required"`
}

type Photo struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	DeviceID   string    `json:"deviceId" gorm:"index;not null"`
	FilePath   string    `json:"file_path" gorm:"not null"`
	Recognized bool      `json:"recognized" gorm:"default:false"`
	Timestamp  time.Time `json:"timestamp" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func NewPhotoData(d *PhotoDto) (*Photo, error) {
	id := uuid.NewAdapter().Generate()
	if d.DeviceID == "" {
		return nil, ErrDeviceIDPhoto
	}
	timestamp := time.Unix(d.Timestamp, 0).UTC()
	if timestamp.IsZero() {
		return nil, ErrTimestampPhoto
	}
	return &Photo{
		ID:         id,
		DeviceID:   d.DeviceID,
		FilePath:   "", // Set after S3 upload
		Timestamp:  timestamp,
		Recognized: false,
	}, nil
}
