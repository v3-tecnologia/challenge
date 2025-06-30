package models

import "time"

type Gyroscope struct {
	ID        uint64    `gorm:"primaryKey" json:"id,omitempty"`
	MAC       string    `gorm:"index;not null" json:"mac"`
	X         float64   `gorm:"not null" json:"x"`
	Y         float64   `gorm:"not null" json:"y"`
	Z         float64   `gorm:"not null" json:"z"`
	Timestamp time.Time `gorm:"autoCreateTime" json:"timestamp"`
}
