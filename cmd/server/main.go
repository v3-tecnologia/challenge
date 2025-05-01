package main

import (
	"github.com/KaiRibeiro/challenge/internal/config"
	"github.com/KaiRibeiro/challenge/internal/db"
	"github.com/KaiRibeiro/challenge/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	db.SetupDb()
	defer db.DB.Close()
	router := gin.Default()
	api := router.Group("/telemetry/")
	api.POST("/gps", routes.SaveGps)
	api.POST("/gyroscope", routes.SaveGyroscope)
	api.POST("/photo", routes.SavePhoto)

	router.Run(":" + config.API_PORT)
}
