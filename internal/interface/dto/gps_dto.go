package dto

import "time"

type CreateGPSRequestDTO struct {
	Latitude  float64   `json:"latitude" binding:"required"`
	Longitude float64   `json:"longitude" binding:"required"`
	DeviceId  string    `json:"device_id" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}
