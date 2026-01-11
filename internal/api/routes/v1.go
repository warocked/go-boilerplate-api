package routes

import (
	"go-boilerplate-api/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupV1Routes configures v1 API routes
func SetupV1Routes(v1 fiber.Router) {
	v1.Post("/login", handlers.LoginHandler)
}
