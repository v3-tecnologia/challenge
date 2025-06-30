package main

import (
	"database/sql"
	"errors"
	"log"
)

func createPhotoTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS photo (
			unique_id TEXT NOT NULL,
			timestamp TEXT NOT NULL,
			photo TEXT NOT NULL
	
		);
	`)
	if err != nil {
		log.Print("failed to create photo table")
		log.Fatal(err)
	}
}

func (ctx *AppContext) insertPhoto(data PhotoData) (sql.Result, error) {
	validation := ctx.PhotoSchema.ValidateStruct(data)
	if !validation.IsValid() {
		log.Print("validation error")
		return nil, errors.New("validation failed")
	}

	result, err := ctx.DB.Exec(`INSERT INTO photo (unique_id, timestamp, photo) VALUES (?, ?, ?)`, data.UniqueId, data.Timestamp, data.Photo)
	if err != nil {
		log.Print("error inserting gyro")
		return nil, err
	}

	return result, nil
}
