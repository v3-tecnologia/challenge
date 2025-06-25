package entity

type Gyroscope struct {
	X         *float64 `json:"x"`
	Y         *float64 `json:"y"`
	Z         *float64 `json:"z"`
	Timestamp string   `json:"timestamp"`
	DeviceID  string   `json:"device_id"`
}
