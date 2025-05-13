package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/challenge/internal/repository"
	"github.com/ricardoraposo/challenge/internal/usecases"
)

type GPSHandler struct {
	useCase usecases.GpsUseCase
}

func NewGPSHandler(useCase usecases.GpsUseCase) *GPSHandler {
	return &GPSHandler{
		useCase: useCase,
	}
}

func (h *GPSHandler) CreateGPSReadings(c *fiber.Ctx) error {
	params := repository.InsertGPSReadingParams{}
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	reading, err := h.useCase.CreateGPSReading(c.Context(), params)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(reading)
}
