package middleware

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func IsUserAuthenticated(c *fiber.Ctx) error {
	authToken := c.Cookies("auth-token")
	if authToken == "" {
		fmt.Println("hay")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token, err := jwt.ParseWithClaims(authToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("dhhfjshdjhs"), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "invalid auth-token",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)
	userId, err := strconv.Atoi(claims.Issuer)
	if err != nil {
		fmt.Println("Invalid user ID in token:", claims.Issuer)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}
	c.Locals("userId", userId)
	return c.Next()

}
