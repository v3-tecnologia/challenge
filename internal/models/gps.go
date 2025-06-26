package models

import (
	"time"

	"github.com/google/uuid"
)

type GPS struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	DeviceID  string    `gorm:"not null" json:"device_id" binding:"required"`
	Latitude  float64   `gorm:"not null" json:"latitude" binding:"required"`
	Longitude float64   `gorm:"not null" json:"longitude" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
