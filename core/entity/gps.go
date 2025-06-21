package entity

type GPS struct {
	Latitude  *float64 `json:"latitude"`  `validate:"required"`
	Longitude *float64 `json:"longitude"` `validate:"required"`
	Timestamp string   `json:"timestamp"` `validate:"required"`
	DeviceID  string   `json:"device_id"` `validate:"required"`
}
