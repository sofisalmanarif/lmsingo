package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitilizeDb() error{
	dsn := "host=localhost user=postgres password=root dbname=lms port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
