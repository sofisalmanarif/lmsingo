package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitilizeDb() error {
	db, err := gorm.Open(postgres.Open(os.Getenv("DNS")), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func GetPostgressClient() (db *gorm.DB, err error) {
	if DB == nil {
		return nil, err
	}
	return DB, nil
}
