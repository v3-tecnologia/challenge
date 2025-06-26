package requests

import "time"

type CreateGyroscopeRequest struct {
	DeviceID  string    `json:"device_id" binding:"required"`
	X         float64   `json:"x" binding:"required"`
	Y         float64   `json:"y" binding:"required"`
	Z         float64   `json:"z" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
}

type CreateGPSRequest struct {
	DeviceID  string    `json:"device_id" binding:"required"`
	Latitude  float64   `json:"latitude" binding:"required"`
	Longitude float64   `json:"longitude" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
}

type CreateTelemetryPhotoRequest struct {
	DeviceID  string    `json:"device_id" binding:"required"`
	Photo     string    `json:"photo" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
}
