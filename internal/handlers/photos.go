package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ricardoraposo/challenge/internal/usecases"
)

type PhotosHandler struct {
	useCase usecases.PhotosUseCase
}

func NewPhotosHandler(useCase usecases.PhotosUseCase) *PhotosHandler {
	return &PhotosHandler{
		useCase: useCase,
	}
}

func (h *PhotosHandler) CreatePhoto(c *fiber.Ctx) error {
	deviceId := c.FormValue("deviceId")
	collectedAt := c.FormValue("collectedAt")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	parsedCollectedAt, err := time.Parse(time.RFC3339, collectedAt)
	if err != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid timestamp format at field 'collectedAt'"})
	}

	var timestamp pgtype.Timestamp
	err = timestamp.Scan(parsedCollectedAt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error converting timestamp")
	}

	params := usecases.CreatePhotoParams{
		DeviceID:    deviceId,
		File:        f,
		Key:         file.Filename,
		CollectedAt: timestamp,
	}

	photo, err := h.useCase.CreatePhoto(c.Context(), params)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(photo)
}
