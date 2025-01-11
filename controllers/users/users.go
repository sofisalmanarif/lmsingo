package users

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RegisterUserType struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUsers(c *fiber.Ctx) error {
	fmt.Println("hitted")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"sucess":  true,
		"message": "Users retrieved successfully",
	})
}

func RegisterUser(c *fiber.Ctx) error {
	fmt.Println("hitted")
	var user RegisterUserType
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}
	fmt.Println(user.Name)
	return nil
	// Register user logic here
}
