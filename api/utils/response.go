package utils

import "github.com/gofiber/fiber/v2"

type HTTPErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type HTTPSuccessResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data,omitempty"`
}

func ResError(c *fiber.Ctx, statusCode int, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}

func ResSuccess(c *fiber.Ctx, statusCode int, data any) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
