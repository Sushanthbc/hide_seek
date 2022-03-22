package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("HIDE_SEEK_DATABASE_HOST"),
		os.Getenv("HIDE_SEEK_DATABASE_USER"),
		os.Getenv("HIDE_SEEK_DATABASE_NAME"),
		os.Getenv("HIDE_SEEK_DATABASE_NAME"),
		os.Getenv("HIDE_SEEK_DATABASE_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	err = database.AutoMigrate(&User{})

	if err != nil {
		log.Fatalf("Auto migrations did not work - %v", err)
	}

	DB = database
}
