package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%v password=basic dbname=hide_seek port=5432 sslmode=disable",
		os.Getenv("HIDE_SEEK_DATABASE_USER"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(&User{})

	return db
}
