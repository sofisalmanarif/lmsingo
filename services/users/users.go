package userServices

import (
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgconn"
	database "github.com/sofisalmanarif/lms/db/postgresql"
	usermodel "github.com/sofisalmanarif/lms/models/users"
)

type UserService interface{
	CreateUser(user *usermodel.Users) (err error)
	Login(user *usermodel.Users) (id int, err error)
	AllUsers() (users []usermodel.Users, err error)
	GetUserDetails(id int) (user *usermodel.Users, err error)
}

type userServices struct{}

func NewUserServices() UserService{
	return &userServices{}
}

func (u *userServices) CreateUser(user *usermodel.Users) (err error) {
	db, err := database.GetPostgressClient()
	if err != nil {
		return err
	}
	result := db.Create(user)
	if result.Error != nil {
		if pgErr, ok := result.Error.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505":
				return fmt.Errorf("email already registered")
			default:
				return err
			}
		}
		return result.Error
	}
	return nil
}

func (u *userServices) Login(user *usermodel.Users) (id int, err error) {
	db, err := database.GetPostgressClient()
	var registeredUser usermodel.Users
	if err != nil {
		return 0, err
	}
	result := db.Where("email = ?", user.Email).First(&registeredUser)
	if result.Error != nil {
		fmt.Println("result", result.Error)
		return 0, fmt.Errorf("invalid credentials")
	}
	fmt.Println("registerd user ", registeredUser)
	err = registeredUser.IsPasswordCorrect(user.Password)
	if err != nil {
		fmt.Println("incorrect password")
		return 0, fmt.Errorf("invalid credentials")
	}

	return registeredUser.ID, nil
}

func (u *userServices) AllUsers() (users []usermodel.Users, err error) {
	var allUsers []usermodel.Users
	db, err := database.GetPostgressClient()
	if err != nil {
		return nil, err
	}
	result := db.Select("ID, Name, Email").Find(&allUsers)
	if result.Error != nil {
		slog.Error(result.Error.Error())
		return nil, fmt.Errorf("failed to fetch users")
	}
	return allUsers, nil

}

func (u *userServices) GetUserDetails(id int) (user *usermodel.Users, err error) {
	db, err := database.GetPostgressClient()
	if err != nil {
		return nil, err
	}
	var searchedUser usermodel.Users
	result := db.Select("ID, Name, Email").Where("id =?", id).First(&searchedUser)

	if result.Error != nil {
		fmt.Println("result", result.Error)
		return nil, fmt.Errorf("user Not Found")
	}
	return &searchedUser, nil

}
