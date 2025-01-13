package users

import (
	"github.com/gofiber/fiber/v2"
	userController "github.com/sofisalmanarif/lms/controllers/users"
	middleware "github.com/sofisalmanarif/lms/middlewares"
)

func UserRouter(app *fiber.App) {
	us :=userController.NewUsersHandler()
	router := app.Group("api/users")
	router.Get("/", middleware.IsUserAuthenticated, us.AllUsers)
	router.Post("/", us.RegisterUser)
	router.Post("/login", us.Login)
	router.Get("/me", middleware.IsUserAuthenticated,us.GetUserDetails)

}
