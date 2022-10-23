package database

import (
	"fmt"
	"mygram/config"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() error {
	conn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=%d sslmode=disable", config.HOST, config.USERNAME, config.PASSWORD, config.DB_NAME, config.PORT)
	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Successfully Connected to Database: ", config.DB_NAME)
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	return nil
}

func GetDB() *gorm.DB {
	return db
}
