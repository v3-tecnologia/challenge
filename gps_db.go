package main

import (
	"database/sql"
	"errors"
	"log"
)

func createGpsTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS gps (
			unique_id TEXT NOT NULL,
			timestamp TEXT NOT NULL,
			latitude REAL NOT NULL,
			longitude REAL NOT NULL
		);
	`)
	if err != nil {
		log.Print("could not create table")
		log.Fatal(err)
	}
}

func (ctx *AppContext) insertGps(data GpsData) (sql.Result, error) {
	validation := ctx.GpsSchema.ValidateStruct(data)
	if !validation.IsValid() {
		log.Print("validation error")
		return nil, errors.New("validation failed")
	}

	result, err := ctx.DB.Exec(`INSERT INTO gps (unique_id, timestamp, latitude, longitude) VALUES (?, ?, ?, ?)`, data.UniqueId, data.Timestamp, data.Latitude, data.Longitude)
	if err != nil {
		log.Printf("error inserting gps")
		return nil, err
	}

	return result, nil
}
