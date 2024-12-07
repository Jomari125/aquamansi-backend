package handlers

import (
	"github.com/gofiber/fiber/v2"
	"aquamansi-backend/utils"
)

// GetInfo serves static application information.
func GetInfo(c *fiber.Ctx) error {
	info := fiber.Map{
		"appName":    "Aquamansi",
		"version":    "1.0.0",
		"description": "Aquamansi is an app for water management and irrigation control.",
	}
	return utils.JSONResponse(c, 200, "Application info retrieved", info)
}

// SubmitData handles POST requests to simulate data processing.
func SubmitData(c *fiber.Ctx) error {
	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return utils.JSONResponse(c, 400, "Invalid request body", nil)
	}

	// Simulate processing the received data
	return utils.JSONResponse(c, 200, "Data received successfully", data)
}
