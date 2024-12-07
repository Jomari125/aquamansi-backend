package utils

import "github.com/gofiber/fiber/v2"

// JSONResponse standardizes JSON responses for the API.
func JSONResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"message": message,
		"data":    data,
	})
}
