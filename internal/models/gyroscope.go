package models

import (
	"time"

	"github.com/google/uuid"
)

type Gyroscope struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	DeviceID  string    `gorm:"not null" json:"device_id" binding:"required"`
	X         float64   `json:"x" gorm:"not null" binding:"required"`
	Y         float64   `json:"y" gorm:"not null" binding:"required"`
	Z         float64   `json:"z" gorm:"not null" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
