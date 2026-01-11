package routes

import (
	"go-boilerplate-api/internal/api/handlers"
	"go-boilerplate-api/shared/helpers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Root API endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return helpers.SendOK(c, nil, "Service is operational")
	})

	// API routes
	api := app.Group("/api")

	// Health check route
	api.Get("/health", handlers.Health)

	// v1 API routes
	SetupV1Routes(api)

	// WebSocket routes
	SetupWebSocketRoutes(app)
}
