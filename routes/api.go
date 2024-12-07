package routes

import (
	"github.com/gofiber/fiber/v2"
	"aquamansi-backend/handlers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Example endpoints
	api.Get("/info", handlers.GetInfo)       // Fetch static info
	api.Post("/submit", handlers.SubmitData) // Handle submitted data
}