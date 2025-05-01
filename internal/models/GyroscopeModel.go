package models

type GyroscopeModel struct {
	X         float64 `json:"x" binding:"required"`
	Y         float64 `json:"y" binding:"required"`
	Z         float64 `json:"z" binding:"required"`
	MAC       string  `json:"mac" binding:"required,mac"`
	Timestamp int64   `json:"timestamp" binding:"required"`
}
