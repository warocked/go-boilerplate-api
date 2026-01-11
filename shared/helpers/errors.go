package helpers

import (
	"github.com/gofiber/fiber/v2"
)

// APIResponse represents a standardized API response format
type APIResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
	Error      *ErrorInfo  `json:"error,omitempty"`
}

// ErrorInfo represents error details in API responses
type ErrorInfo struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// SendSuccess sends a successful response with data
func SendSuccess(c *fiber.Ctx, statusCode int, data interface{}, message string) error {
	return c.Status(statusCode).JSON(APIResponse{
		StatusCode: statusCode,
		Data:       data,
		Message:    message,
	})
}

// SendError sends an error response
func SendError(c *fiber.Ctx, statusCode int, errorCode string, message string) error {
	return SendErrorWithDetails(c, statusCode, errorCode, message, nil)
}

// SendErrorWithDetails sends an error response with additional details
func SendErrorWithDetails(c *fiber.Ctx, statusCode int, errorCode string, message string, details interface{}) error {
	return c.Status(statusCode).JSON(APIResponse{
		StatusCode: statusCode,
		Error: &ErrorInfo{
			Code:    errorCode,
			Message: message,
			Details: details,
		},
	})
}

// SendOK sends a 200 OK response
func SendOK(c *fiber.Ctx, data interface{}, message string) error {
	return SendSuccess(c, fiber.StatusOK, data, message)
}

// SendCreated sends a 201 Created response
func SendCreated(c *fiber.Ctx, data interface{}, message string) error {
	return SendSuccess(c, fiber.StatusCreated, data, message)
}

// SendBadRequest sends a 400 Bad Request error
func SendBadRequest(c *fiber.Ctx, errorCode string, message string) error {
	return SendError(c, fiber.StatusBadRequest, errorCode, message)
}

// SendUnauthorized sends a 401 Unauthorized error
func SendUnauthorized(c *fiber.Ctx, message string) error {
	return SendError(c, fiber.StatusUnauthorized, "unauthorized", message)
}

// SendForbidden sends a 403 Forbidden error
func SendForbidden(c *fiber.Ctx, message string) error {
	return SendError(c, fiber.StatusForbidden, "forbidden", message)
}

// SendNotFound sends a 404 Not Found error
func SendNotFound(c *fiber.Ctx, message string) error {
	return SendError(c, fiber.StatusNotFound, "not_found", message)
}

// SendInternalServerError sends a 500 Internal Server Error
func SendInternalServerError(c *fiber.Ctx, message string) error {
	return SendError(c, fiber.StatusInternalServerError, "internal_error", message)
}
