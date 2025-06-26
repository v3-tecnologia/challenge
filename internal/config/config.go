package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	ServerPort string
	NatsURL    string
}

func LoadConfig() (*Config, error) {
	if os.Getenv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	return &Config{
		DBHost:     getEnv("DB_HOST", "postgres"),
		DBUser:     getEnv("DB_USER", "telemetry_user"),
		DBPassword: getEnv("DB_PASSWORD", "telemetry_P0stgr3s_2025!"),
		DBName:     getEnv("DB_NAME", "telemetry_db"),
		DBPort:     getEnv("DB_PORT", "5432"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		NatsURL:    getEnv("NATS_URL", "nats://localhost:4222"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
