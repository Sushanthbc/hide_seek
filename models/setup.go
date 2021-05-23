package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDatabase() {
	dsn := "host=localhost user=sushanth password=basic dbname=hide_seek port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(&User{})
}
