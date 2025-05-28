package main

import (
	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/rekognition"
	"github.com/KaiRibeiro/challenge/internal/routes"
	"github.com/KaiRibeiro/challenge/internal/s3"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	s3.InitS3()
	rekognition.InitRekognition()
	defer db.DB.Close()

	gyroscopeService := services.NewGyroscopeDBService(db.DB)
	gyroscopeHandler := routes.NewGyroscopeHandler(gyroscopeService)

	gpsService := services.NewGPSDBService(db.DB)
	gpsHandler := routes.NewGpsHandler(gpsService)

	photoService := services.NewPhotoDBService(db.DB)
	photoHandler := routes.NewPhotoHandler(photoService)

	router := gin.Default()
	api := router.Group("/telemetry/")
	api.POST("/gps", gpsHandler.SaveGps)
	api.POST("/gyroscope", gyroscopeHandler.SaveGyroscope)
	api.POST("/photo", photoHandler.SavePhoto)

	router.Run(":" + config.API_PORT)
}
