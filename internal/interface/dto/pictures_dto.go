package dto

import "time"

type CreatePictureRequestDTO struct {
	PictureData []byte    `json:"picture_data" binding:"required"`
	DeviceId    string    `json:"device_id" binding:"required"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
	PictureType string    `json:"picture_type" binding:"required"`
}

type CreatePictureResponseDTO struct {
	Id           string    `json:"id"`
	DeviceId     string    `json:"device_id"`
	CreatedAt    time.Time `json:"created_at"`
	ReceivedAt   time.Time `json:"received_at"`
	IsRecognized bool      `json:"is_recognized"`
	PictureURL   string    `json:"picture_url"`
}
