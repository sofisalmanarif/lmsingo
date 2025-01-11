package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	userhandlers "github.com/sofisalmanarif/lms/handlers/users"
	usermodel "github.com/sofisalmanarif/lms/models/users"
)

func GetUsers(c *fiber.Ctx) error {
	fmt.Println("hitted")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"sucess":  true,
		"message": "Users retrieved successfully",
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
	_, err := userhandlers.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create user",
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
	fmt.Println(user.Name)
	 err := userhandlers.Login(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message":err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User logged in successfully",
		})
		

}
