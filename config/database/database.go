package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func InitializeDatabases() *gorm.DB {
	db := connectPostgresDatabase()

	return db
}

func connectPostgresDatabase() *gorm.DB {
	postgresDSN := os.Getenv("POSTGRES_DSN")

	db, err := gorm.Open(postgres.Open(postgresDSN))

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	database, err := db.DB()

	if err != nil {
		log.Fatalf("Failed to get database: %v", err)
	}

	database.SetMaxIdleConns(10)
	database.SetMaxOpenConns(100)
	database.SetConnMaxLifetime(time.Hour)
	database.SetConnMaxIdleTime(time.Minute)

	if err := database.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to PostgresSQL")

	return db
}
