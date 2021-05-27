package models

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string
	PasswordHash string
	Name         string
}

//Fetch all the users
func GetAllUser(db *gorm.DB) (*gorm.DB, error) {

	created_record := db.Create(&User{
		Email:        "sushanthbc@gmail.com",
		PasswordHash: "asdasdasdasdas",
		Name:         "Sushanth",
	})

	if created_record.Error != nil {
		log.Fatal("Something went wrong in creating a record")
	}

	var users []User

	result := db.Find(&users)
	err := result.Error

	if err != nil {
		log.Fatal("Could not retrieve all User information")
	}

	return result, err
}
