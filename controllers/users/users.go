package users

import (
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	userhandlers "github.com/sofisalmanarif/lms/handlers/users"
	usermodel "github.com/sofisalmanarif/lms/models/users"
	utilities "github.com/sofisalmanarif/lms/utils"
)

func AllUsers(c *fiber.Ctx) error {
	fmt.Println("hitted", c.Locals("userId"))
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
	validate := validator.New()
	var user usermodel.Users
	fmt.Println("hitted")
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid request body",
			})
		}
		var errors []string
		// Validation errors
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("'%s' failed type is '%s'", err.Field(), err.Tag()))
			fmt.Printf("Validation error: Field '%s' failed on '%s'\n", err.Field(), err.Tag())
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": errors,
		})
	}

	err = userhandlers.CreateUser(&user)
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

func GetuserDetails(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	fmt.Printf("userid %T", userId)
	user, err := userhandlers.GetUserDetails(userId.(int))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})

}
