package model

type Gyroscope struct {
	BaseModel
	BaseTelemetry
	AxisX int `gorm:"not null"`
	AxisY int `gorm:"not null"`
	AxisZ int `gorm:"not null"`
}
