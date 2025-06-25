package models

import "time"

type GpsModel struct {
	Latitude  *float64  `json:"latitude" bson:"latitude"`
	Longitude *float64  `json:"longitude" bson:"longitude"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}
