package controllers

import (
	"errors"
	"jewete/entities"
	"jewete/handler"
	"jewete/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := entities.JWTClaim{
		User: user,
		StandardClaims: &jwt.StandardClaims{
			Audience:  string(c.Request().Header.Referer()),
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Issuer:    strconv.Itoa(int(user.ID)),
		},
	}

	token, err := services.NewToken(&claims)
	if err != nil {
		return response.ServerError(err)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    *token,
		Expires:  expirationTime,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return response.Success(token)
}

func User(c *fiber.Ctx) error {
	token := c.Cookies("jwt")
	response := handler.NewResponse(c)

	claims, err := services.ParseJwt(&token)
	if err != nil {
		return response.Unauthorized()
	}

	user := services.GetUserById(&claims.Issuer)
	return response.Success(user)
}

func Logout(c *fiber.Ctx) error {
	response := handler.NewResponse(c)
	cookie := fiber.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	c.Cookie(&cookie)
	return response.Success(nil)
}
