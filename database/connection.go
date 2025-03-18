package database

import (
	"github.com/abhijeetsharan/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:1234@/yt_go_auth"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
