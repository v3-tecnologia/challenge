package dto

import (
	"io"

	"github.com/jackc/pgx/v5/pgtype"
)

type InsertPhotosDto struct {
	DeviceID    string           `json:"deviceId" validate:"required"`
	File        io.Reader        `json:"file" validate:"required"`
	CollectedAt pgtype.Timestamp `json:"collectedAt" validate:"required"`
}
