package http

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wellmtx/challenge/docs"
	"github.com/wellmtx/challenge/internal/http/controller"
)

type Router struct {
	gyroscopeController   *controller.GyroscopeController
	geolocationController *controller.GeolocationController
	photoController       *controller.PhotoController
}

func NewRouter(
	gyroscopeController *controller.GyroscopeController,
	geolocationController *controller.GeolocationController,
	photoController *controller.PhotoController,
) *Router {
	return &Router{
		gyroscopeController:   gyroscopeController,
		geolocationController: geolocationController,
		photoController:       photoController,
	}
}

func (router *Router) Init() {
	r := gin.Default()
	r.Use(gin.Recovery())

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	docs.SwaggerInfo.BasePath = "/"

	telemetry := r.Group("/telemetry")
	{
		telemetry.POST("/gyroscope", router.gyroscopeController.SaveData)
		telemetry.POST("/gps", router.geolocationController.SaveData)
		telemetry.POST("/photo", router.photoController.RecognizePhoto)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
