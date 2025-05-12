package main

import (
	"fmt"
	"log"
	"v3/internal/domain"
	"v3/internal/infra/api"
	"v3/internal/infra/aws"
	"v3/internal/repository/gps"
	"v3/internal/repository/gyroscope"
	"v3/internal/repository/photo"
	usecase "v3/internal/usecase"
	"v3/internal/utils"
)

func main() {
	dsn := utils.GetDSN()
	db, err := utils.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	fmt.Println("Conectado ao banco de dados com sucesso!")

	err = db.AutoMigrate(&domain.Gyroscope{}, &domain.GPS{}, &domain.Photo{})
	if err != nil {
		log.Fatalf("erro ao fazer AutoMigrate: %v", err)
	}

	awsService, err := aws.NewAWSService("telemetry-photos")
	if err != nil {
		log.Fatalf("failed to initialize AWS service: %v", err)
	}

	gyroRepo := gyroscope.NewGyroscopeRepository(db)
	gpsRepo := gps.NewGPSRepository(db)
	photoRepo := photo.NewPhotoRepository(db)

	createGyroUseCase := usecase.NewCreateGyroscopeUseCase(gyroRepo)
	createGPSUseCase := usecase.NewCreateGPSUseCase(gpsRepo)
	createPhotoUseCase := usecase.NewCreatePhotoUseCase(photoRepo, awsService)

	gyroHandlers := api.NewGyroscopeHandlers(createGyroUseCase)
	gpsHandlers := api.NewGPSHandlers(createGPSUseCase)
	photoHandlers := api.NewPhotoHandlers(createPhotoUseCase)

	router := api.SetupRouter(gyroHandlers, gpsHandlers, photoHandlers)
	router.Run(":8080")
}
