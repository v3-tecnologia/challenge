package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/challenge/internal/repository"
	"github.com/ricardoraposo/challenge/internal/usecases"
)

type GyroscopeHandler struct {
	useCase usecases.GyroscopeUseCase
}

func NewGyroscopeHandler(useCase usecases.GyroscopeUseCase) *GyroscopeHandler {
	return &GyroscopeHandler{
		useCase: useCase,
	}
}

func (h *GyroscopeHandler) CreateGyroscopeReadings(c *fiber.Ctx) error {
	params := repository.InsertGyroscopeReadingParams{}
	if err := c.BodyParser(&params); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	reading, err := h.useCase.CreateGyroscopeReading(c.Context(), params)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(reading)
}
