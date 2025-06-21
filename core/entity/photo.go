package entity

type Photo struct {
	ImageBase64 string `json:"image_base64"` `validate:"required"`
	Timestamp   string `json:"timestamp"`	 `validate:"required"`
	DeviceID    string `json:"device_id"`	 `validate:"required"`
}
