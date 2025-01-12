package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sofisalmanarif/lms/controllers/users"
	middleware "github.com/sofisalmanarif/lms/middlewares"
)

func UserRouter(app *fiber.App) {
	router := app.Group("api/users")
	router.Get("/", middleware.IsUserAuthenticated, users.AllUsers)
	router.Post("/", users.RegisterUser)
	router.Post("/login", users.Login)

}
