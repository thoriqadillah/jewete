package entities

import "github.com/gofiber/fiber/v2"

type Claim struct {
	User   *User
	Cookie *fiber.Cookie
}
