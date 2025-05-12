package entity

import (
	"errors"
	"time"
)

type GyroscopeTelemetry struct {
	BaseEntity
	X float64 `json:"x" gorm:"type:float;not null"`
	Y float64 `json:"y" gorm:"type:float;not null"`
	Z float64 `json:"z" gorm:"type:float;not null"`
}

type GPSTelemetry struct {
	BaseEntity
	Latitude  float64 `json:"latitude" gorm:"type:float;not null"`
	Longitude float64 `json:"longitude" gorm:"type:float;not null"`
}

func (GPSTelemetry) TableName() string {
	return "gps_telemetry"
}

func (GyroscopeTelemetry) TableName() string {
	return "gyroscope_telemetry"
}

func BuildGyroscopeTelemetry(deviceId string, createdAt time.Time, x, y, z float64) *GyroscopeTelemetry {
	return &GyroscopeTelemetry{
		BaseEntity: BaseEntity{
			DeviceID:   deviceId,
			CreatedAt:  createdAt,
			ReceivedAt: time.Now(),
		},
		X: x,
		Y: y,
		Z: z,
	}
}

func BuildGPSTelemetry(deviceId string, createdAt time.Time, latitude, longitude float64) *GPSTelemetry {
	return &GPSTelemetry{
		BaseEntity: BaseEntity{
			DeviceID:   deviceId,
			CreatedAt:  createdAt,
			ReceivedAt: time.Now(),
		},
		Latitude: latitude,
	}

}

func (g *GyroscopeTelemetry) Validate() error {

	if g.DeviceID == "" {
		return errors.New("o ID do dispositivo não pode estar vazio")
	}

	if g.CreatedAt.IsZero() {
		return errors.New("a data de criação não pode ser zero")
	}

	if g.ReceivedAt.IsZero() {
		return errors.New("a data de recebimento não pode ser zero")
	}

	return nil
}

func (g *GPSTelemetry) Validate() error {

	if g.DeviceID == "" {
		return errors.New("o ID do dispositivo não pode estar vazio")
	}

	if g.CreatedAt.IsZero() {
		return errors.New("a data de criação não pode ser zero")
	}

	if g.ReceivedAt.IsZero() {
		return errors.New("a data de recebimento não pode ser zero")
	}

	if g.Latitude < -90 || g.Latitude > 90 {
		return errors.New("a latitude deve estar entre -90 e 90")
	}
	return nil
}
