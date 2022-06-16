package main

import (
	"jewete/database"
	"jewete/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
	defer database.Close()
}
