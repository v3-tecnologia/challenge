package models

type PhotoModel struct {
	ImageBase64 string `json:"image_base_64"`
	MAC         string `json:"mac"`
	Timestamp   int64  `json:"timestamp"`
}
