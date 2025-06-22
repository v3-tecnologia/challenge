package entity

type Photo struct {
	ImageBase64 string `json:"image_base64"`
	Timestamp   string `json:"timestamp"`
	DeviceID    string `json:"device_id"`
}
