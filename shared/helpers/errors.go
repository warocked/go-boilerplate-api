package helpers

import "github.com/gofiber/fiber/v2"

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse represents a standardized success response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SendError sends a standardized error response
func SendError(c *fiber.Ctx, statusCode int, errorCode string, message string) error {
	return c.Status(statusCode).JSON(ErrorResponse{
		Error:   errorCode,
		Message: message,
	})
}

// SendSuccess sends a standardized success response
func SendSuccess(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}
