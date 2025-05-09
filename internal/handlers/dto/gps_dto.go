package dto

import "github.com/jackc/pgx/v5/pgtype"

type InsertGPSReadingsDto struct {
	DeviceID    string           `json:"deviceId,valida" validate:"required"`
	Latitude    pgtype.Numeric   `json:"latitude" validate:"required"`
	Longitude   pgtype.Numeric   `json:"longitude" validate:"required"`
	CollectedAt pgtype.Timestamp `json:"collectedAt" validate:"required"`
}
