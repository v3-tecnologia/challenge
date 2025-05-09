package model

type Gyroscope struct {
	BaseModel
	BaseTelemetry
	AxisX float64 `gorm:"not null"`
	AxisY float64 `gorm:"not null"`
	AxisZ float64 `gorm:"not null"`
}
