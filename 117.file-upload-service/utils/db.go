package utils

import (
	"file-upload-service/models"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	fmt.Println("Connecting to DB...")

	var err error
	DB, err = gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect db - %s\n", err.Error())
	}

	DB.AutoMigrate(&models.User{}, &models.File{})

	fmt.Println("Connected to db successfully!")
}
