package model

type Photo struct {
	Device
	PhotoBase64 string `json:"photo_base64"`
}
