package dto

import "time"

type CreateGyroscopeRequestDTO struct {
	X         float64   `json:"x" binding:"required"`
	Y         float64   `json:"y" binding:"required"`
	Z         float64   `json:"z" binding:"required"`
	DeviceId  string    `json:"device_id" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}
