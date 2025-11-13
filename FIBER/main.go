package main

import (
	routes "fiber/src/router"

	"github.com/gofiber/fiber/v2"
)


func main (){
	app := fiber.New()
	routes.UserRoutes(app)
	app.Listen(":3000")

}
