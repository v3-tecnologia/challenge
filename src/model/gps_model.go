package model

type GPS struct {
	BaseTelemetry
	BaseModel
	Latitude  float64 `gorm:"not null"`
	Longitude float64 `gorm:"not null"`
}
