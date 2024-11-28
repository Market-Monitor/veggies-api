package utils

import "github.com/gofiber/fiber/v3"

func ResError(c fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}

func ResSuccess(c fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
