package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"aquamansi-backend/routes"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// Register Routes
	routes.RegisterRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
