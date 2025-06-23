package entity

type Photo struct {
	Image     []byte `json:"image"`
	DeviceID  string `json:"device_id"`
	Timestamp string `json:"timestamp"`
}
