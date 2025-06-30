package main

import (
	"challenge-cloud/internal/config"
	controller "challenge-cloud/internal/controllers"
	repository "challenge-cloud/internal/repositories/gorm"
	"challenge-cloud/internal/router"
	service "challenge-cloud/internal/services"
	"log"
	"net/http"
)

func main() {

	config.LoadEnvGorm()
	config.InitDB()
	db := config.DB

	gyroscopeRepository := repository.NewGyroscopeRepository(db)
	gyroscopeService := service.NewGyroscopeService(gyroscopeRepository)
	gyroscopeController := controller.NewGyroscopeController(gyroscopeService)

	gpsRepository := repository.NewGPSRepository(db)
	gpsService := service.NewGPSService(gpsRepository)
	gpsController := controller.NewGPSController(gpsService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoController := controller.NewPhotoController(photoService)

	c := router.Controllers{
		Gyro:  gyroscopeController,
		GPS:   gpsController,
		Photo: photoController,
	}
	r := router.LoadRouter(c)
	log.Println("🚀 Servidor rodando em http://localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}
