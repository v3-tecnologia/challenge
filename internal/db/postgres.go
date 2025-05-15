package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	createTables(db)

	return db
}

func createTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS gyroscope (id SERIAL PRIMARY KEY, x FLOAT NOT NULL, y FLOAT NOT NULL, z FLOAT NOT NULL)`,
		`CREATE TABLE IF NOT EXISTS gps (id SERIAL PRIMARY KEY, latitude FLOAT NOT NULL, longitude FLOAT NOT NULL, altitude FLOAT NOT NULL)`,
		`CREATE TABLE IF NOT EXISTS photos (id SERIAL PRIMARY KEY, filename TEXT NOT NULL, data TEXT NOT NULL)`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}
}
