package gps

import "github.com/iamrosada0/v3/internal/repository/gps"

type GPSInputDto struct {
	DeviceID  string  `json:"deviceId"`
	Timestamp int64   `json:"timestamp"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CreateGPSUseCase struct {
	Repo gps.GPSRepository
}

func NewCreateGPSUseCase(repo gps.GPSRepository) *CreateGPSUseCase {
	return &CreateGPSUseCase{Repo: repo}
}
