package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()
	err := database.InitilizeDb()
	if err != nil {
		log.Fatalln("Database Connection failed")
		panic(err)
	}

	db, err := database.GetPostgressClient()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Printf("migrations failed")
		log.Fatal(err)
	}
	users.UserRouter(app)

	app.Listen(":3000")

}
