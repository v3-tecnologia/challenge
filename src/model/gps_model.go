package model

type GPS struct {
	BaseTelemetry
	BaseModel
	Latitude  int `gorm:"not null"`
	Longitude int `gorm:"not null"`
}
