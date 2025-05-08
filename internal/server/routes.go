package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ricardoraposo/challenge/internal/dto"
	"github.com/ricardoraposo/challenge/internal/handlers"
	"github.com/ricardoraposo/challenge/internal/middleware"
	"github.com/ricardoraposo/challenge/internal/usecases"
)

func (s *FiberServer) RegisterRoutes() {
	s.Use(logger.New())
	s.Use(cors.New())

	s.App.Get("/health", health)

	gyroscopeUseCase := usecases.NewGyroscopeUseCase(s.Database.Query)
	gpsUseCase := usecases.NewGPSUseCase(s.Database.Query)

	gyroscopeHandler := handlers.NewGyroscopeHandler(gyroscopeUseCase)
	gpsHandler := handlers.NewGPSHandler(gpsUseCase)

	telemetryApi := s.App.Group("/telemetry")

	telemetryApi.Post("/gyroscope", middleware.ValidateStruct[dto.InsertGryoscopeReadingsDto], gyroscopeHandler.CreateGyroscopeReadings)
	telemetryApi.Post("/gps", middleware.ValidateStruct[dto.InsertGPSReadingsDto], gpsHandler.CreateGPSReadings)
}

func health(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "ok"})
}
