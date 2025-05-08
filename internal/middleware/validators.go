package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ricardoraposo/challenge/internal/common"
)

func ValidateStruct[T any](c *fiber.Ctx) error {
	var params T

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	validationErrors := common.Validate(params)
	if len(validationErrors) > 0 {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"errors": validationErrors})
	}

	// Store the parsed and validated data in Locals to be used in the handler
	c.Locals("validatedBody", params)

	return c.Next()
}
