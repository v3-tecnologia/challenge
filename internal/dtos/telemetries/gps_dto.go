package dtos

type CreateGpsDto struct {
	Latitude  *float64 `json:"latitude"  binding:"required,gt=-90,lt=90"`
	Longitude *float64 `json:"longitude" binding:"required,gt=-180,lt=180"`
}
