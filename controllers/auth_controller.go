package controllers

import (
	"errors"
	"jewete/entities"
	"jewete/handler"
	"jewete/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	var request entities.User
	response := handler.NewResponse(c)

	if err := c.BodyParser(&request); err != nil {
		return response.Error(err)
	}

	record, err := services.CreateUser(&request)
	if err != nil {
		return response.Error(err)
	}

	return response.Created(record)
}

func Login(c *fiber.Ctx) error {
	var request entities.User
	response := handler.NewResponse(c)

	if err := c.BodyParser(&request); err != nil {
		return response.Error(err)
	}

	user, err := services.GetUser(&request)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NotFound()
	}

	token, expiredTime, err := services.NewToken(user.ID)
	if err != nil {
		return response.ServerError(err)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    *token,
		Expires:  expiredTime,
		HTTPOnly: true,
	}

	jwt := entities.Claim{
		User:   user,
		Cookie: &cookie,
	}

	return response.SuccessWithCookie(&jwt)
}
