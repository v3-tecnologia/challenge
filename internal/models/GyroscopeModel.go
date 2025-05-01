package models

type GyroscopeModel struct {
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
	MAC       string  `json:"mac"`
	Timestamp int64   `json:"timestamp"`
}
