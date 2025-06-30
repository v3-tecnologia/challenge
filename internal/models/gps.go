package models

import "time"

type GPS struct {
	ID        uint64    `gorm:"primaryKey" json:"id,omitempty"`
	MAC       string    `gorm:"index;not null" json:"mac"`
	Latitude  float64   `gorm:"not null" json:"latitude"`
	Longitude float64   `gorm:"not null" json:"longitude"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
