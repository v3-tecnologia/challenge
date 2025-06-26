package models

import (
	"time"

	"github.com/google/uuid"
)

type TelemetryPhoto struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	DeviceID  string    `gorm:"not null" json:"device_id" binding:"required"`
	Photo     string    `gorm:"not null" json:"photo" binding:"required"`
	Timestamp time.Time `gorm:"not null" json:"timestamp" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
