package dto

type BaseTelemetry struct {
	MacAddr string `json:"mac_addr" binding:"required"`
}

type Gyroscope struct {
	BaseTelemetry BaseTelemetry
	X             int `json:"x" binding:"required"`
	Y             int `json:"y" binding:"required"`
	Z             int `json:"z" binding:"required"`
}

type GPS struct {
	BaseTelemetry BaseTelemetry
	Latitude      int `json:"latitude" binding:"required"`
	Longitude     int `json:"longitude" binding:"required"`
}
