package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `sql:"unqiue" json:"email"`
	PasswordHash []byte
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
