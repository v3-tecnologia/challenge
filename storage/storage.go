package storage

import (
	"challenge-v3/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"

	_ "github.com/lib/pq"
)

type Storage interface {
	SaveGyroscope(data *models.GyroscopeData) error
	SaveGPS(data *models.GPSData) error
	SavePhoto(data *models.PhotoData) error
	LogAuditEvent(event models.AuditEvent) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("não foi possível conectar ao postgres: %w", err)
	}
	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) InitTables() error {
	gyroscopeTable := `
	CREATE TABLE IF NOT EXISTS gyroscope (
		id SERIAL PRIMARY KEY,
		device_id TEXT NOT NULL,
		x REAL NOT NULL,
		y REAL NOT NULL,
		z REAL NOT NULL,
		timestamp TIMESTAMP NOT NULL
	);`

	gpsTable := `
	CREATE TABLE IF NOT EXISTS gps (
		id SERIAL PRIMARY KEY,
		device_id TEXT NOT NULL,
		latitude REAL NOT NULL,
		longitude REAL NOT NULL,
		timestamp TIMESTAMP NOT NULL
	);`

	photoTable := `
	CREATE TABLE IF NOT EXISTS photo (
		id SERIAL PRIMARY KEY,
		device_id TEXT NOT NULL,
		photo TEXT NOT NULL,
		timestamp TIMESTAMP NOT NULL,
		recognized BOOLEAN NOT NULL DEFAULT FALSE
	);`

	auditTable := `
    CREATE TABLE IF NOT EXISTS audit_log (
        id SERIAL PRIMARY KEY,
        timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        actor TEXT NOT NULL,
        action TEXT NOT NULL,
        details JSONB
    );`

	tables := []string{gyroscopeTable, gpsTable, photoTable, auditTable}
	for _, tableSQL := range tables {
		if _, err := s.db.Exec(tableSQL); err != nil {
			return err
		}
	}
	slog.Info("Tabelas do banco de dados inicializadas com sucesso")
	return nil
}

func (s *PostgresStorage) LogAuditEvent(event models.AuditEvent) error {
	query := "INSERT INTO audit_log(actor, action, details) VALUES($1, $2, $3)"

	detailsJSON, err := json.Marshal(event.Details)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(query, event.Actor, event.Action, detailsJSON)
	return err
}

func (s *PostgresStorage) SaveGyroscope(data *models.GyroscopeData) error {
	query := "INSERT INTO gyroscope(device_id, x, y, z, timestamp) VALUES($1, $2, $3, $4, $5)"
	_, err := s.db.Exec(query, data.DeviceID, *data.X, *data.Y, *data.Z, data.Timestamp)
	return err
}

func (s *PostgresStorage) SaveGPS(data *models.GPSData) error {
	query := "INSERT INTO gps(device_id, latitude, longitude, timestamp) VALUES($1, $2, $3, $4)"
	_, err := s.db.Exec(query, data.DeviceID, *data.Latitude, *data.Longitude, data.Timestamp)
	return err
}

func (s *PostgresStorage) SavePhoto(data *models.PhotoData) error {
	query := "INSERT INTO photo(device_id, photo, timestamp, recognized) VALUES($1, $2, $3, $4)"
	_, err := s.db.Exec(query, data.DeviceID, data.Photo, data.Timestamp, data.Recognized)
	return err
}
