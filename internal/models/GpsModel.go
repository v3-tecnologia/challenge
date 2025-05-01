package models

type GpsModel struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	MAC       string  `json:"mac"`
	Timestamp int64   `json:"timestamp"`
}
