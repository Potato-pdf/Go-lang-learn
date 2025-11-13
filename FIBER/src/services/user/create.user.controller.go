package user

import "github.com/gofiber/fiber/v2"

func CreateUsers(c *fiber.Ctx) error {
	return c.SendString("Controlador: Create products")
}x|