package MiddleWare

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

func HandleError(c *fiber.Ctx, err error, message []string) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"status":  code,
		"message": message,
		"error":   err,
	})
}
