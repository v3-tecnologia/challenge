package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"telemetry-api/internal/handlers"
	"telemetry-api/internal/middleware"
	"telemetry-api/internal/services"
)

func SetupRouter(db *gorm.DB, nc *nats.Conn, logger *zap.Logger) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.GlobalRateLimitMiddleware())

	telemetryService := services.NewTelemetryService(db, logger)
	authService := services.NewAuthService(db, logger)

	telemetryHandler := handlers.NewTelemetryHandler(telemetryService, nc, logger)
	authHandler := handlers.NewAuthHandler(authService, logger)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Public auth routes
	authRoutes := r.Group("v1/auth")
	authRoutes.Use(middleware.AuthRateLimitMiddleware())
	{
		authRoutes.POST("/register", authHandler.CreateUser)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/refresh", authHandler.RefreshToken)
	}

	// Protected telemetry routes
	telemetryRoutes := r.Group("v1/telemetry")
	telemetryRoutes.Use(middleware.TelemetryRateLimitMiddleware())
	telemetryRoutes.Use(middleware.JWTAuth())
	telemetryRoutes.Use(middleware.RequireRole("admin"))
	{
		// POST routes (criar dados)
		telemetryRoutes.POST("/gyroscope", telemetryHandler.CreateGyroscope)
		telemetryRoutes.POST("/gps", telemetryHandler.CreateGPS)
		telemetryRoutes.POST("/photo", telemetryHandler.CreateTelemetryPhoto)

		// GET routes (consultar dados)
		telemetryRoutes.GET("/gyroscope", telemetryHandler.GetGyroscopeData)
		telemetryRoutes.GET("/gps", telemetryHandler.GetGPSData)
		telemetryRoutes.GET("/photo", telemetryHandler.GetPhotoData)
		telemetryRoutes.GET("/devices", telemetryHandler.GetDevices)
	}

	return r
}
