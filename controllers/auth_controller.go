package controllers

import (
	"jewete/entities"
	"jewete/handler"
	"jewete/services"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user entities.User
	response := handler.Reponse{Ctx: c}

	if err := c.BodyParser(&user); err != nil {
		return response.Error(err)
	}

	record, err := services.CreateUser(&user)
	if err != nil {
		return response.Error(err)
	}

	return response.Created(record)
}
