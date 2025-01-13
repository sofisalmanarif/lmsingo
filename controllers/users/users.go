package userController

import (
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	usermodel "github.com/sofisalmanarif/lms/models/users"
	userServices "github.com/sofisalmanarif/lms/services/users"
	utilities "github.com/sofisalmanarif/lms/utils"
)

type Controller interface {
	AllUsers(c *fiber.Ctx) error
	RegisterUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetUserDetails(c *fiber.Ctx) error
}

type UsersHandler struct {
	Validator *validator.Validate
	Service   userServices.UserService
}

func NewUsersHandler() Controller {
	return &UsersHandler{
		Validator: validator.New(),
		Service:   userServices.NewUserServices(),
	}
}

func (h *UsersHandler) AllUsers(c *fiber.Ctx) error {
	fmt.Println("hitted", c.Locals("userId"))
	users, err := h.Service.AllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "All Users",
		"data":    users,
	})
}

func (h *UsersHandler) RegisterUser(c *fiber.Ctx) error {
	var user usermodel.Users
	fmt.Println("hitted")
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	err := h.Validator.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Invalid request body",
			})
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("'%s' failed type is '%s'", err.Field(), err.Tag()))
			fmt.Printf("Validation error: Field '%s' failed on '%s'\n", err.Field(), err.Tag())
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": errors,
		})
	}

	err = h.Service.CreateUser(&user)
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

func (h *UsersHandler) Login(c *fiber.Ctx) error {
	fmt.Println("hitted")
	var user usermodel.Users
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}
	fmt.Println(user.Email)
	id, err := h.Service.Login(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	token, err := utilities.GenerateJWTToken(id)
	if err != nil {
		log.Fatal("something went wrong")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "something went wrong",
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

func (h *UsersHandler) GetUserDetails(c *fiber.Ctx) error {
	userId := c.Locals("userId")
	fmt.Printf("userid %T", userId)
	user, err := h.Service.GetUserDetails(userId.(int))

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
