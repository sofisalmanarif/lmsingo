package userhandlers

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
	database "github.com/sofisalmanarif/lms/db/postgresql"
	usermodel "github.com/sofisalmanarif/lms/models/users"
)

func CreateUser(user *usermodel.Users) (err error) {
	db, err := database.GetPostgressClient()
	if err != nil {
		return err
	}
	result := db.Create(user)
	if result.Error != nil {
		if pgErr, ok := result.Error.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case  "23505":
				return errors.New("email already registered")
			default:
				return err
			}
		}
		return result.Error
	}
	return nil
}

func Login(user *usermodel.Users) (id int, err error) {
	db, err := database.GetPostgressClient()
	var registeredUser *usermodel.Users
	if err != nil {
		return 0, err
	}
	result := db.Where("email = ?", user.Email).First(&registeredUser)
	if result.Error != nil {
		fmt.Println("result", result.Error)
		return 0, errors.New("invalid credentials")
	}
	fmt.Println("registerd user ", registeredUser)
	err = registeredUser.IsPasswordCorrect(user.Password)
	if err != nil {
		fmt.Println("incorrect password")
		return 0, errors.New("invalid credentials")
	}

	return registeredUser.ID, nil
}


func AllUsers()(users []*usermodel.Users,err error){
	var AllUsers []*usermodel.Users
	db,err :=database.GetPostgressClient()
	if err != nil{
		return nil,err
	}
	result :=db.Select("ID, Name, Email").Find(&AllUsers)
	if result.Error != nil{
		slog.Error(result.Error.Error())
		return nil, fmt.Errorf("failed to fetch users",)
	}
	return AllUsers,nil

}