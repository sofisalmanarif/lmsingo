package users

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	userhandlers "github.com/sofisalmanarif/lms/handlers/users"
	usermodel "github.com/sofisalmanarif/lms/models/users"
	utilities "github.com/sofisalmanarif/lms/utils"
)

func AllUsers(c *fiber.Ctx) error {
	fmt.Println("hitted",c.Locals("userId"))
	users, err := userhandlers.AllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"sucess":  true,
		"message": "All Users",
		"data":    users,
	})
}

func RegisterUser(c *fiber.Ctx) error {
	fmt.Println("hitted")
	var user usermodel.Users
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}
	fmt.Println(user.Name)
	err := userhandlers.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
	})
}

func Login(c *fiber.Ctx) error {
	fmt.Println("hitted")
	var user usermodel.Users
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}
	fmt.Println(user.Email)
	id, err := userhandlers.Login(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	token, err := utilities.GenerateJWTToken(id)
	if err != nil {
		log.Fatal("someting went wrong")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "someting went wrong",
		})

	}

	cookie := fiber.Cookie{
		Name:     "auth-token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Login successfully",
		"token":   token,
	})

}
