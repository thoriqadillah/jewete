package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Reponse struct {
	*fiber.Ctx
}

func (r *Reponse) Error(err error) error {
	return r.Ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"status":  "error",
		"message": err.Error(),
	})
}

func (r *Reponse) Created(data interface{}) error {
	return r.Ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func (r *Reponse) Success(data interface{}) error {
	return r.Ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func (r *Reponse) NotFound() error {
	return r.Ctx.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": "data not found",
	})
}
