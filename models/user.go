package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type User struct {
	Email        string
	PasswordHash string
	Name         string
}

//create new user
func CreateNewUser(db gorm.DB, email string, password_hash string, name string) error {
	res := db.Create(&User{
		Email:        email,
		PasswordHash: password_hash,
		Name:         name,
	})

	if res.Error != nil {
		err_message := fmt.Errorf("Could not create a new user - %v", res.Error)
		log.Fatal(err_message)
		return err_message
	}

	return nil
}

//Fetch all the users
func GetAllUser(db *gorm.DB) *gorm.DB {
	var users []User

	res := db.Find(&users)

	return res
}
