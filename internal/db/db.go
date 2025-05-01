package db

import (
	"database/sql"
	"fmt"
	"github.com/KaiRibeiro/challenge/internal/config"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDb() {
	var err error

	DB, err = sql.Open("postgres", config.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
}

func SetupDb() {
	fmt.Println("Creating tables.")

	gpsTable := `
	CREATE TABLE IF NOT EXISTS gps (
		id SERIAL PRIMARY KEY,
		latitude DOUBLE PRECISION NOT NULL,
		longitude DOUBLE PRECISION NOT NULL,
		mac TEXT NOT NULL,
		timestamp TIMESTAMPTZ NOT NULL
	);`

	gyroTable := `
	CREATE TABLE IF NOT EXISTS gyroscope (
		id SERIAL PRIMARY KEY,
		x DOUBLE PRECISION NOT NULL,
		y DOUBLE PRECISION NOT NULL,
		z DOUBLE PRECISION NOT NULL,
		mac TEXT NOT NULL,
		timestamp TIMESTAMPTZ NOT NULL
	);`

	photoTable := `
	CREATE TABLE IF NOT EXISTS photo (
		id SERIAL PRIMARY KEY,
		image_base64 TEXT NOT NULL,
		mac TEXT NOT NULL,
		timestamp TIMESTAMPTZ NOT NULL
	);`

	statements := []string{gpsTable, gyroTable, photoTable}

	for _, stmt := range statements {
		if _, err := DB.Exec(stmt); err != nil {
			log.Fatalf("Error creating table: %v\n", err)
		}
	}

	fmt.Println("Tables created or already exist.")
}
