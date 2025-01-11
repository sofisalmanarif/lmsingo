package userhandlers

import (
	"errors"
	"fmt"

	database "github.com/sofisalmanarif/lms/db/postgresql"
	usermodel "github.com/sofisalmanarif/lms/models/users"
)

func CreateUser(user *usermodel.Users) (id int64, err error) {
	db, err := database.GetPostgressClient()
	if err != nil {
		return 0, err
	}
	result := db.Create(user)
	if result.Error != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

func Login(user *usermodel.Users) (err error) {
	db, err := database.GetPostgressClient()
	var registeredUser usermodel.Users
	if err != nil {
		return err
	}
	result := db.Where("email = ?", user.Email).First(&registeredUser)
	if result.Error != nil {
		fmt.Println("result", result.Error)
		return errors.New("invalid credentials")
	}
	fmt.Println("registerd user ", registeredUser)
	err = registeredUser.IsPasswordCorrect(user.Password)
	if err != nil {
		fmt.Println("incorrect password")
		return errors.New("invalid credentials")
	}

	return nil
}
