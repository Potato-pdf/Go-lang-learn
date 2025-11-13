package user

import (
	"fiber/src/interfaces/structures/users"

	"github.com/gofiber/fiber/v2"
)

func CreateUsers(c *fiber.Ctx) error {
	user := users.User{}
	if err := c.BodyParser(&user); err !=nil{
		return  err
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
