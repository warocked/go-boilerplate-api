package main

import (
	"go-boilerplate-api/internal/api/config"
	"go-boilerplate-api/shared/helpers"

	"github.com/gofiber/fiber/v2"
)

// errorHandler handles errors and returns standardized responses
func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"
	errorCode := "internal_error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
		
		switch code {
		case fiber.StatusBadRequest:
			errorCode = "bad_request"
		case fiber.StatusUnauthorized:
			errorCode = "unauthorized"
		case fiber.StatusForbidden:
			errorCode = "forbidden"
		case fiber.StatusNotFound:
			errorCode = "not_found"
		case fiber.StatusConflict:
			errorCode = "conflict"
		case fiber.StatusUnprocessableEntity:
			errorCode = "validation_error"
		case fiber.StatusTooManyRequests:
			errorCode = "rate_limit_exceeded"
		case fiber.StatusInternalServerError:
			errorCode = "internal_error"
		case fiber.StatusServiceUnavailable:
			errorCode = "service_unavailable"
		}
	}

	// Hide internal error details in production
	if code >= 500 && config.IS_PROD {
		message = "An internal server error occurred"
	}

	return helpers.SendError(c, code, errorCode, message)
}
