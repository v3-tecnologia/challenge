package main

import (
	"challenge-v3-backend/config/database"
	"challenge-v3-backend/internal/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.InitializeDatabases()

	postgresDatabase, _ := db.DB()

	defer func() {
		log.Println("Closing postgresDatabase connection")
		postgresDatabase.Close()
	}()

	port := os.Getenv("PORT")

	routes := router.SetupRoutes(db)

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	serverAddress := ":" + port

	log.Printf("Server is running on %s", serverAddress)

	if err := routes.Run(serverAddress); err != nil {
		log.Fatal("Failed to start server: %v", err)
	}

}
