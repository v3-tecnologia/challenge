package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/challenge/internal/database"
	"github.com/ricardoraposo/challenge/internal/repository"
)

type GyroscopeHandler struct {
	database *database.Database
}

func NewGyroscopeHandler(database *database.Database) *GyroscopeHandler {
	return &GyroscopeHandler{
		database,
	}
}

func (h *GyroscopeHandler) CreateGyroscopeReadings(c *fiber.Ctx) error {
	params := repository.InsertGyroscopeReadingParams{}
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	reading, err := h.database.Query.InsertGyroscopeReading(c.Context(), params)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(reading)
}
