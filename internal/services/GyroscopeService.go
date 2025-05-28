package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/KaiRibeiro/challenge/internal/custom_errors"
	"github.com/KaiRibeiro/challenge/internal/models"
)

type GyroscopeService interface {
	AddGyroscope(gyroscope models.GyroscopeModel) error
}

type GyroscopeDBService struct {
	DB *sql.DB
}

func NewGyroscopeDBService(dbConn *sql.DB) *GyroscopeDBService {
	return &GyroscopeDBService{DB: dbConn}
}

func (s *GyroscopeDBService) AddGyroscope(gyroscope models.GyroscopeModel) error {
	timestamp := time.UnixMilli(gyroscope.Timestamp)

	query := `INSERT INTO gyroscope (x, y, z, mac, timestamp)
    VALUES ($1, $2, $3, $4, $5)`

	_, err := s.DB.Exec(query, gyroscope.X, gyroscope.Y, gyroscope.Z, gyroscope.MAC, timestamp)

	if err != nil {
		return fmt.Errorf("failed to insert gyroscope data into database: %w", custom_errors.NewDBError(err, http.StatusInternalServerError))
	}

	return nil
}
