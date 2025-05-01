package models

type GpsModel struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	MAC       string  `json:"mac" binding:"required,mac"`
	Timestamp int64   `json:"timestamp" binding:"required"`
}
