package model

type GyroscopeData struct {
	ID int     `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
	Z  float64 `json:"z"`
}

type GPSData struct {
	ID        int     `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
}

type PhotoData struct {
	ID       int    `json:"id"`
	Filename string `json:"filename"`
	Data     string `json:"data"`
}
