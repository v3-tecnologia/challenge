package env

import (
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalln("Error loading .env file, server cannot be initialized:", err)
	}
}
