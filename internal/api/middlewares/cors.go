package middlewares

import "github.com/gofiber/fiber/v2/middleware/cors"

func SetupCorsConfig() cors.Config {
	return cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Origin,  Content-Type,  Accept,  Accept-Language,  Content-Length",
	}
}
