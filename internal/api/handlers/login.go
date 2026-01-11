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

	if err := c.BodyParser(&params); err != nil {
		return helpers.SendBadRequest(c, "invalid_request", "Invalid request body")
	}

	if err := helpers.ValidateStruct(&params); err != nil {
		return helpers.SendBadRequest(c, "validation_error", err.Error())
	}

	// TODO: Implement login logic
	return helpers.SendError(c, fiber.StatusNotImplemented, "not_implemented", "Login functionality not yet implemented")
}
