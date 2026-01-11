package routes

import (
	"go-boilerplate-api/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// API routes
	api := app.Group("/api")

	// Health check route
	api.Get("/health", handlers.Health)

	// v1 API routes
	SetupV1Routes(api)

	// WebSocket routes
	SetupWebSocketRoutes(app)
}
