package main

import (
	"challenge-cloud/internal/config"
	controller "challenge-cloud/internal/controllers"
	repository "challenge-cloud/internal/repositories/gorm"
	"challenge-cloud/internal/router"
	service "challenge-cloud/internal/services"
	"fmt"
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

	authController := controller.NewAuthController(db)

	c := router.Controllers{
		Gyro:  gyroscopeController,
		GPS:   gpsController,
		Photo: photoController,
		Auth:  authController,
	}
	r := router.LoadRouter(c)
	port := fmt.Sprintf(":%d", config.Port)
	log.Printf("🚀 Servidor rodando em http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, r))

}
