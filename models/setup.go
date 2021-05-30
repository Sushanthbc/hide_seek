package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database_host     = os.Getenv("HIDE_SEEK_DATABASE_HOST")
	database_user     = os.Getenv("HIDE_SEEK_DATABASE_USER")
	database_password = os.Getenv("HIDE_SEEK_DATABASE_PASSWORD")
	database_name     = os.Getenv("HIDE_SEEK_DATABASE_NAME")
	database_port     = os.Getenv("HIDE_SEEK_DATABASE_PORT")
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		database_host,
		database_user,
		database_password,
		database_name,
		database_port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	err = database.AutoMigrate(&User{})

	if err != nil {
		log.Fatalf("Auto migrations did not work - %v", err)
	}

	DB = database
}
