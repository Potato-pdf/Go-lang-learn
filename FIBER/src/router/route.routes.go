package routes

import (
	"fiber/src/services/user"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App){
	UserGroup := app.Group("/user")	
	UserGroup.Get("/get", user.GetUsers)
	UserGroup.Post("/post", user.CreateUsers)
}