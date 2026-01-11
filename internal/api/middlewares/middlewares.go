package middlewares

import "github.com/gofiber/fiber/v2"

func SetupMiddlewares(app *fiber.App) {
	// app.Use()
	SetupMiddlewaresEssentials(app)
}
