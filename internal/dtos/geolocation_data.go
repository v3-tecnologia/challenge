package dtos

type GeolocationDataDto struct {
	BaseDTO
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
