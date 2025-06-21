package entity

type Gyroscope struct {
	X         *float64 `json:"x"` 			`validate:"required"`
	Y         *float64 `json:"y"` 			`validate:"required"`
	Z         *float64 `json:"z"`			`validate:"required"`
	Timestamp string   `json:"timestamp"`	`validate:"required"`
	DeviceID  string   `json:"device_id"`	`validate:"required"`
}
