package entity

type GPS struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
	Timestamp string   `json:"timestamp"`
	DeviceID  string   `json:"device_id"`
}
