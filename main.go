package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	database "github.com/sofisalmanarif/lms/db/postgresql"
	"github.com/sofisalmanarif/lms/routes/users"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	app := fiber.New()
	err = database.InitilizeDb()
	if err != nil {
		log.Fatalln("Database Connection failed")
		panic(err)
	}

	users.UserRouter(app)

	app.Listen(":3000")

}
