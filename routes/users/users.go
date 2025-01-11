package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sofisalmanarif/lms/controllers/users"
)

func UserRouter(app *fiber.App) {
	router:=app.Group("api/users")
	router.Get("/", users.GetUsers)
	router.Post("/",users.RegisterUser)
	router.Post("/login",users.Login)

}
