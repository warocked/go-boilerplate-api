package routes

import (
	"go-boilerplate-api/internal/api/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// SetupWebSocketRoutes configures WebSocket routes
func SetupWebSocketRoutes(app *fiber.App) {
	app.Get("/ws", websocket.New(handlers.HandleWebSocket))
}
