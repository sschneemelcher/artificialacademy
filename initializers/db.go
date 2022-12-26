package initializers

import (
	"fmt"
	"sschneemelcher/artificialacademy/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Message{})
}