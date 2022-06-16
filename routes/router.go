package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Route("/auth", authRoute)
}
