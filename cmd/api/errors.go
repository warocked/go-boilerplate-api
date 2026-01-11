package main

import (
	"github.com/gofiber/fiber/v2"
)

// errorHandler handles errors and returns standardized responses
func errorHandler(c *fiber.Ctx, err error) error {
	// Default status code
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// Check if it's a Fiber error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	// Don't expose internal error details in production
	errorCode := "internal_error"
	if code >= 400 && code < 500 {
		errorCode = "client_error"
	}

	return c.Status(code).JSON(fiber.Map{
		"error":   errorCode,
		"message": message,
	})
}
