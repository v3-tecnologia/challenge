package domain

import "time"

type GpsDomain struct {
	ID             int
	Latitude       float64
	Longitude      float64
	DeviceID       int
	MacAddress     string
	CollectionDate time.Time
}
