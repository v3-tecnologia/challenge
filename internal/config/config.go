package config

import (
	"log"
	"os"
)

var (
	API_PORT = getEnv("API_PORT")
	DbUrl    = getEnv("DB_URL")
)

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("No env variable %s set.", key)
	}
	return value
}
