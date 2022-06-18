package handler

import (
	"jewete/entities"

	"github.com/gofiber/fiber/v2"
)

type Reponse struct {
	Ctx *fiber.Ctx
}

func NewResponse(c *fiber.Ctx) *Reponse {
	return &Reponse{Ctx: c}
}

func (r *Reponse) Error(err error) error {
	return r.Ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "error",
		"message": err.Error(),
	})
}

func (r *Reponse) Created(data interface{}) error {
	return r.Ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func (r *Reponse) Success(data interface{}) error {
	return r.Ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func (r *Reponse) SuccessWithCookie(data interface{}) error {
	claim := data.(*entities.Claim)

	r.Ctx.Cookie(claim.Cookie)
	return r.Ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   claim.User,
	})
}

func (r *Reponse) NotFound() error {
	return r.Ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "data not found",
	})
}

func (r *Reponse) ServerError(err error) error {
	return r.Ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"status":  "error",
		"message": err.Error(),
	})
}

func (r *Reponse) Unauthorized() error {
	return r.Ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "unauthenticated",
	})
}
