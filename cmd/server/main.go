package main

import (
	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/logs"
	"github.com/KaiRibeiro/challenge/internal/middlewares"
	"github.com/KaiRibeiro/challenge/internal/rekognition"
	"github.com/KaiRibeiro/challenge/internal/routes"
	"github.com/KaiRibeiro/challenge/internal/s3"
	"github.com/KaiRibeiro/challenge/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {

	logs.Init()

	db.InitDb()
	s3.InitS3()
	rekognition.InitRekognition()
	defer db.DB.Close()

	uploader := &s3.S3Uploader{}
	faceComparer := rekognition.NewRekognitionComparer(db.DB)

	gyroscopeService := services.NewGyroscopeDBService(db.DB)
	gyroscopeHandler := routes.NewGyroscopeHandler(gyroscopeService)

	gpsService := services.NewGPSDBService(db.DB)
	gpsHandler := routes.NewGpsHandler(gpsService)

	photoService := services.NewPhotoDBService(db.DB, uploader, faceComparer)
	photoHandler := routes.NewPhotoHandler(photoService)

	router := gin.Default()
	router.Use(middlewares.LoggingMiddleware())
	api := router.Group("/telemetry/")
	api.POST("/gps", gpsHandler.SaveGps)
	api.POST("/gyroscope", gyroscopeHandler.SaveGyroscope)
	api.POST("/photo", photoHandler.SavePhoto)

	router.Run(":" + config.API_PORT)
}
