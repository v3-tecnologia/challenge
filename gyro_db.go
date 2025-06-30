package main

import (
	"database/sql"
	"errors"
	"log"
)

func createGyroTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS gyro (
			unique_id TEXT NOT NULL,
			timestamp TEXT NOT NULL,
			x REAL NOT NULL,
			y REAL NOT NULL,
			z REAL NOT NULL
		);
	`)
	if err != nil {
		log.Print("error creating gyro table")
		log.Fatal(err)
	}
}

func (ctx *AppContext) insertGyro(data GyroData) (sql.Result, error) {
	validation := ctx.GyroSchema.ValidateStruct(data)
	if !validation.IsValid() {
		log.Print("validation error")
		return nil, errors.New("validation failed")
	}

	result, err := ctx.DB.Exec(`INSERT INTO gyro (unique_id, timestamp, x, y, z) VALUES (?, ?, ?, ?, ?)`, data.UniqueId, data.Timestamp, data.X, data.Y, data.Z)
	if err != nil {
		log.Printf("error inserting gyro")
		return nil, err
	}

	return result, nil
}
