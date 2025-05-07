package db

import (
	"database/sql"
	"fmt"
	"log"

	env "github.com/bielgennaro/v3-challenge-cloud/internal/handlers"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// GetConnection returns a connection to the PostgreSQL database
func GetConnection() (*sql.DB, error) {
	config := Config{
		Host:     env.GetEnv("DB_HOST", "localhost"),
		Port:     env.GetEnv("DB_PORT", "5432"),
		User:     env.GetEnv("DB_USER", "postgres"),
		Password: env.GetEnv("DB_PASSWORD", "postgres"),
		DBName:   env.GetEnv("DB_NAME", "postgres"),
		SSLMode:  env.GetEnv("DB_SSLMODE", "disable"),
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("could not establish connection to database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}
