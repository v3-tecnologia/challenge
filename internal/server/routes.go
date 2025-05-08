package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ricardoraposo/challenge/internal/handlers"
	"github.com/ricardoraposo/challenge/internal/middleware"
)

func (s *FiberServer) RegisterRoutes() {
	s.Use(logger.New())
	s.Use(cors.New())

	s.App.Get("/health", health)

	gyroscopeHandler := handlers.NewGyroscopeHandler(s.Database)

	telemetryApi := s.App.Group("/telemetry")

	telemetryApi.Post("/gyroscope", middleware.ValidateInsertGyroscopeReadingParams, gyroscopeHandler.CreateGyroscopeReadings)
}

func health(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "ok"})
}
