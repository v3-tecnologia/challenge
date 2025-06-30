package models

import "time"

type GPS struct {
	ID        uint64    `gorm:"primaryKey" json:"id,omitempty"`
	MAC       string    `gorm:"index;not null" json:"mac" validate:"required"`
	Latitude  float64   `gorm:"not null" json:"latitude" validate:"required"`
	Longitude float64   `gorm:"not null" json:"longitude" validate:"required"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
