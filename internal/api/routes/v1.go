package routes

import (
	"go-boilerplate-api/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupV1Routes configures v1 API routes
func SetupV1Routes(api fiber.Router) {
	v1 := api.Group("/v1")

	v1.Post("/login", handlers.LoginHandler)
}
