package response

import (
	"time"

	"github.com/google/uuid"
)

type CreateGyroscopeResponse struct {
	ID        uuid.UUID `json:"id"`
	DeviceID  string    `json:"device_id"`
	X         float64   `json:"x"`
	Y         float64   `json:"y"`
	Z         float64   `json:"z"`
	Timestamp time.Time `json:"timestamp"`
}

type CreateGPSResponse struct {
	ID        uuid.UUID `json:"id"`
	DeviceID  string    `json:"device_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
	CreatedAt time.Time `json:"created_at"`
}

type GyroscopeResponse struct {
	ID        uuid.UUID `json:"id"`
	DeviceID  string    `json:"device_id"`
	X         float64   `json:"x"`
	Y         float64   `json:"y"`
	Z         float64   `json:"z"`
	Timestamp time.Time `json:"timestamp"`
	CreatedAt time.Time `json:"created_at"`
}

type GPSResponse struct {
	ID        uuid.UUID `json:"id"`
	DeviceID  string    `json:"device_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
	CreatedAt time.Time `json:"created_at"`
}

type TelemetryPhotoResponse struct {
	ID        uuid.UUID `json:"id"`
	DeviceID  string    `json:"device_id"`
	Photo     string    `json:"photo"`
	Timestamp time.Time `json:"timestamp"`
	CreatedAt time.Time `json:"created_at"`
}

type DeviceResponse struct {
	DeviceID        string    `json:"device_id"`
	LastSeen        time.Time `json:"last_seen"`
	GyroscopeCount  int64     `json:"gyroscope_count"`
	GPSCount        int64     `json:"gps_count"`
	PhotoCount      int64     `json:"photo_count"`
	TotalDataPoints int64     `json:"total_data_points"`
}

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

type GyroscopeListResponse struct {
	Data       []GyroscopeResponse `json:"data"`
	Pagination PaginationMeta      `json:"pagination"`
}

type GPSListResponse struct {
	Data       []GPSResponse  `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

type TelemetryPhotoListResponse struct {
	Data       []TelemetryPhotoResponse `json:"data"`
	Pagination PaginationMeta           `json:"pagination"`
}

type DeviceListResponse struct {
	Data       []DeviceResponse `json:"data"`
	Pagination PaginationMeta   `json:"pagination"`
}
