package model

type GPSData struct {
	Device
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
