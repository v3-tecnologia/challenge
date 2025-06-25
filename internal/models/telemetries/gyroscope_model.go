package models

import "time"

type GyroscopeModel struct {
	X         float64   `json:"x" bson:"x"`
	Y         float64   `json:"y" bson:"y"`
	Z         float64   `json:"z" bson:"z"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
