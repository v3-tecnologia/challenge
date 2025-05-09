package model

import "time"

type BaseTelemetry struct {
	MacAddr           string    `gorm:"not null"`
	DateTimeCollected time.Time `gorm:"not null"`
}
