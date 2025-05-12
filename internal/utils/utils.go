package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDSN() string {
	if url := os.Getenv("DATABASE_URL"); url != "" {
		fmt.Printf("Using DATABASE_URL: %s\n", url)
		return url
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if host == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Missing required database environment variables: DB_HOST, DB_USER, DB_PASSWORD, DB_NAME")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", host, user, password, dbname)
	fmt.Printf("Using DSN from environment: %s\n", dsn)
	return dsn
}

func ConnectDB(dsn string) (*gorm.DB, error) {
	fmt.Printf("Connecting to database with DSN: %s\n", dsn)
	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := db.DB()
			if err == nil {
				if err = sqlDB.Ping(); err == nil {
					fmt.Println("Database connection successful")
					return db, nil
				}
			}
		}
		fmt.Printf("Failed to connect to database: %v. Retrying in 5s (attempt %d/10)...\n", err, i+1)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to database after 10 retries: %w", err)
}
