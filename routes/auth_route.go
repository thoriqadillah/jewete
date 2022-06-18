package routes

import (
	"jewete/controllers"

	"github.com/gofiber/fiber/v2"
)

func authRoute(auth fiber.Router) {
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)
	auth.Get("/user", controllers.User)

	// auth.Post("/refresh", controllers.RefreshToken)
	// auth.Post("/logout", controllers.Logout)

}
