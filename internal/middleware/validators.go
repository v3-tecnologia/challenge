package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/challenge/internal/common"
	"github.com/ricardoraposo/challenge/internal/dto"
)

func ValidateInsertGyroscopeReadingParams(c *fiber.Ctx) error {
	params := dto.InsertGryoscopeReadingsDto{}
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	validationErrors := common.Validate(params)

	if len(validationErrors) > 0 {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"errors": validationErrors})
	}

	return c.Next()
}
