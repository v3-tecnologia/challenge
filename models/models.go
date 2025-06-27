package models

import (
	"errors"
	"time"
)

type GyroscopeData struct {
	DeviceID  string    `json:"device_id"`
	X         *float64  `json:"x"`
	Y         *float64  `json:"y"`
	Z         *float64  `json:"z"`
	Timestamp time.Time `json:"timestamp"`
}

func (g *GyroscopeData) Validate() error {
	if g.DeviceID == "" {
		return errors.New("campo obrigatório ausente: device_id")
	}
	if g.Timestamp.IsZero() {
		return errors.New("campo obrigatório ausente: timestamp")
	}
	if g.X == nil {
		return errors.New("campo obrigatório ausente: x")
	}
	if g.Y == nil {
		return errors.New("campo obrigatório ausente: y")
	}
	if g.Z == nil {
		return errors.New("campo obrigatório ausente: z")
	}
	return nil
}

type GPSData struct {
	DeviceID  string    `json:"device_id"`
	Latitude  *float64  `json:"latitude"`
	Longitude *float64  `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}

func (gps *GPSData) Validate() error {
	if gps.DeviceID == "" {
		return errors.New("campo obrigatório ausente: device_id")
	}
	if gps.Timestamp.IsZero() {
		return errors.New("campo obrigatório ausente: timestamp")
	}
	if gps.Latitude == nil {
		return errors.New("campo obrigatório ausente: latitude")
	}
	if gps.Longitude == nil {
		return errors.New("campo obrigatório ausente: longitude")
	}
	return nil
}

type PhotoRequest struct {
	DeviceID  string    `json:"device_id"`
	Photo     string    `json:"photo"`
	Timestamp time.Time `json:"timestamp"`
}

type PhotoData struct {
	DeviceID   string    `json:"device_id"`
	Photo      string    `json:"photo"`
	Timestamp  time.Time `json:"timestamp"`
	Recognized bool      `json:"recognized"`
}

type AuditEvent struct {
	Actor   string                 `json:"actor"`
	Action  string                 `json:"action"`
	Details map[string]interface{} `json:"details"`
}

func (p *PhotoData) Validate() error {
	if p.DeviceID == "" {
		return errors.New("campo obrigatório ausente: device_id")
	}
	if p.Timestamp.IsZero() {
		return errors.New("campo obrigatório ausente: timestamp")
	}
	if p.Photo == "" {
		return errors.New("campo obrigatório ausente: photo")
	}
	return nil
}

type ErrorResponse struct {
	Message string `json:"message"`
}
