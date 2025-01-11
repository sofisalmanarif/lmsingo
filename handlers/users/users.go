package userhandlers

import (
	"log"

	database "github.com/sofisalmanarif/lms/db/postgresql"
	usermodel "github.com/sofisalmanarif/lms/models/users"
)

func CreateUser(user *usermodel.Users) (id int64,err error){
	db, err := database.GetPostgressClient()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	result :=db.Create(user)
	if result.Error  !=nil{
		log.Fatal(result.Error)
		return 0, err
	}
	return result.RowsAffected, nil


}
