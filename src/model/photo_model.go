package model

type Photo struct {
	BaseModel
	BaseTelemetry
	Name  string `gorm:"not null"`
	Image []byte `gorm:"not null"`
}
