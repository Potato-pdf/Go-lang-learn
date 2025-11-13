package user

import (
	"fiber/src/interfaces/structures/users"

	"github.com/gofiber/fiber/v2"
)

func CreateUsers(c *fiber.Ctx) error {
	user := users.User{
		Firstname: "Cesar",
		Lastname : "Bernal",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
