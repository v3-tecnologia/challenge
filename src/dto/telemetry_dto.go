package dto

import "time"

type BaseTelemetry struct {
	MacAddr           string    `json:"mac_addr" binding:"required"`
	DateTimeCollected time.Time `json:"date_time_collected" binding:"required"`
}

type Gyroscope struct {
	BaseTelemetry
	AxisX int `json:"x" binding:"required"`
	AxisY int `json:"y" binding:"required"`
	AxisZ int `json:"z" binding:"required"`
}

type GPS struct {
	BaseTelemetry
	Latitude  int `json:"latitude" binding:"required"`
	Longitude int `json:"longitude" binding:"required"`
}
