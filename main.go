package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sofisalmanarif/lms/routes/users"
)

func main() {
	app := fiber.New()

	users.UserRouter(app)

	app.Listen(":3000")

}
