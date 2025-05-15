package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
}

func Load() Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	return Config{
		DatabaseURL: dbURL,
	}
}
