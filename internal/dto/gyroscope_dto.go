package dto

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type InsertGryoscopeReadingsDto struct {
	DeviceID    string           `json:"deviceId" validate:"required"`
	X           float64          `json:"x" validate:"required,gt=0"`
	Y           float64          `json:"y" validate:"required,gt=0"`
	Z           float64          `json:"z" validate:"required,gt=0"`
	CollectedAt pgtype.Timestamp `json:"collectedAt" validate:"required"`
}
