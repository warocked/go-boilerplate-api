package routes

import (
	"go-boilerplate-api/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	// Health check route
	v1.Get("/health", handlers.Health)

	// v1 API routes
	SetupV1Routes(v1)

	// WebSocket routes
	SetupWebSocketRoutes(app)
}
