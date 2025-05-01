package models

type PhotoModel struct {
	ImageBase64 string `json:"image_base_64" binding:"required,base64"`
	MAC         string `json:"mac" binding:"required,mac"`
	Timestamp   int64  `json:"timestamp" binding:"required"`
}
