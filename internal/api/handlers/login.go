package handlers

import (
	"go-boilerplate-api/shared/helpers"

	"github.com/gofiber/fiber/v2"
)

type LoginParams struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Password string `json:"password" validate:"required,min=6"`
}

func LoginHandler(c *fiber.Ctx) error {
	var params LoginParams

	// Parse JSON body
	if err := c.BodyParser(&params); err != nil {
		return helpers.SendError(c, fiber.StatusBadRequest, "invalid_request", "Invalid request body")
	}

	// Validate input
	if err := helpers.ValidateStruct(&params); err != nil {
		return helpers.SendError(c, fiber.StatusBadRequest, "validation_error", err.Error())
	}

	// - Check username/password against database
	// - Generate JWT token
	// - Return token

	return helpers.SendError(c, fiber.StatusNotImplemented, "not_implemented", "Login functionality not yet implemented")
}
