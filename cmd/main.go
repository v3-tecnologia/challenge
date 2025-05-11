package main

import (
	"fmt"
	"log"
	"v3/internal/infra/api"
	"v3/internal/repository/gps"
	"v3/internal/repository/gyroscope"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	usecase "v3/internal/usecase"
)

func main() {
	// Initialize database
	dsn := "host=postgres user=meuusuario password=minhasenha dbname=meubanco port=5432 sslmode=disable"

	// Conectar ao banco de dados
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	fmt.Println("Conectado ao banco de dados com sucesso!")

	// Initialize repositories
	gyroRepo := gyroscope.NewGyroscopeRepository(db)
	gpsRepo := gps.NewGPSRepository(db)
	// photoRepo := photo.NewPhotoRepository(db)

	// Initialize AWS Rekognition client (simplified)
	//rekogClient := initializeRekognitionClient() // Custom function

	// Initialize use cases
	createGyroUseCase := usecase.NewCreateGyroscopeUseCase(gyroRepo)
	createGPSUseCase := usecase.NewCreateGPSUseCase(gpsRepo)
	// createPhotoUseCase := usecase.NewCreatePhotoUseCase(photoRepo, rekogClient)

	// Initialize handlers
	gyroHandlers := api.NewGyroscopeHandlers(createGyroUseCase)
	gpsHandlers := api.NewGPSHandlers(createGPSUseCase)
	// photoHandlers := api.NewPhotoHandlers(createPhotoUseCase)

	// Setup router
	router := api.SetupRouter(gyroHandlers, gpsHandlers)

	// Start server
	router.Run(":8080")
}
