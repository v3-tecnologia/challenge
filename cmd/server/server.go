package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/wellmtx/challenge/internal/http"
	"github.com/wellmtx/challenge/internal/http/controller"
	"github.com/wellmtx/challenge/internal/infra/database"
	"github.com/wellmtx/challenge/internal/infra/providers"
	"github.com/wellmtx/challenge/internal/infra/repositories"
	"github.com/wellmtx/challenge/internal/service"
)

func init() {
	godotenv.Load()
}

// @title           Wellmtx V3 Challenge MVP
// @version         1.0
// @description     Essa é a API do MVP do desafio V3 de Wellington Saraiva.
// @contact.name  Wellington Saraiva
// @host 		localhost:8080
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	db := database.NewDatabase(
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_DB"),
		false,
	)

	if err := db.Connect(); err != nil {
		panic(err)
	}

	gyroscopeRepository := repositories.NewGyroscopeRepository(db)
	gyroscopeService := service.NewGyroscopeService(gyroscopeRepository)
	gyroscopeController := controller.NewGyroscopeController(gyroscopeService)

	geolocationRepository := repositories.NewGeolocationRepository(db)
	geolocationService := service.NewGeolocationService(geolocationRepository)
	geolocationController := controller.NewGeolocationController(geolocationService)

	recognitionProvider := providers.NewRecognitionProvider()
	photoRepository := repositories.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository, recognitionProvider)
	photoController := controller.NewPhotoController(photoService)

	httpRouter := http.NewRouter(gyroscopeController, geolocationController, photoController)
	httpRouter.Init()
}
