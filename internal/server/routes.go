package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ricardoraposo/challenge/internal/handlers"
	"github.com/ricardoraposo/challenge/internal/handlers/dto"
	"github.com/ricardoraposo/challenge/internal/middleware"
	"github.com/ricardoraposo/challenge/internal/services"
	"github.com/ricardoraposo/challenge/internal/usecases"
)

func (s *FiberServer) RegisterRoutes() {
	s.Use(logger.New())
	s.Use(cors.New())

	s.App.Get("/health", health)

	gyroscopeUseCase := usecases.NewGyroscopeUseCase(s.Database.Query)
	gpsUseCase := usecases.NewGPSUseCase(s.Database.Query)
	photosUseCase := usecases.NewPhotosUseCase(
		s.Database.Query,
		services.NewS3Uploader(),
		services.NewRekognitionClient(),
	)

	gyroscopeHandler := handlers.NewGyroscopeHandler(gyroscopeUseCase)
	gpsHandler := handlers.NewGPSHandler(gpsUseCase)
	photosHandler := handlers.NewPhotosHandler(photosUseCase)

	telemetryApi := s.App.Group("/telemetry")

	telemetryApi.Post("/gyroscope", middleware.ValidateJSONBodyStruct[dto.InsertGryoscopeReadingsDto], gyroscopeHandler.CreateGyroscopeReadings)
	telemetryApi.Post("/gps", middleware.ValidateJSONBodyStruct[dto.InsertGPSReadingsDto], gpsHandler.CreateGPSReadings)
	telemetryApi.Post("/photo", middleware.ValidateCreatePhotoParams, photosHandler.CreatePhoto)

	//NOTE I'm leaving this route for demo purposes, it should not be used in production
	demoApi := s.App.Group("/demo")
	demoApi.Post("collection", func(c *fiber.Ctx) error {
		client := services.NewRekognitionClient()
		res, err := client.CreateCollection()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "error"})
		}

		return c.Status(http.StatusOK).JSON(res)
	})
}

func health(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "ok"})
}
