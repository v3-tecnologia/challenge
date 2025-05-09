package dto

import "time"

type BaseTelemetry struct {
	MacAddr           string    `json:"mac_addr" binding:"required"`
	DateTimeCollected time.Time `json:"date_time_collected" binding:"required"`
}

type Gyroscope struct {
	BaseTelemetry
	AxisX float64 `json:"x" binding:"required"`
	AxisY float64 `json:"y" binding:"required"`
	AxisZ float64 `json:"z" binding:"required"`
}

type GPS struct {
	BaseTelemetry
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
