package models

import "time"

// Telemetry models for gyroscope, GPS, and photo data
type Gyroscope struct {
	ID        int64     `gorm:"primaryKey"`
	X         float64   `gorm:"not null"`
	Y         float64   `gorm:"not null"`
	Z         float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type GPS struct {
	ID        int64     `gorm:"primaryKey"`
	Latitude  float64   `gorm:"not null"`
	Longitude float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Photo struct {
	ID        int64     `gorm:"primaryKey"`
	Image     []byte    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// DTOs for request validation
type GyroscopeDTO struct {
	X float64 `json:"x" binding:"required"`
	Y float64 `json:"y" binding:"required"`
	Z float64 `json:"z" binding:"required"`
}

type GPSDTO struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type PhotoDTO struct {
	Image     []byte    `gorm:"not null"`
}
