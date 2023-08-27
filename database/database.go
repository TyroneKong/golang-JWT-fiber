package database

import (
	"learnfiber/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connection, err := gorm.Open(mysql.Open("root:root@/golang"), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	} else {
		println("We are connected to the database")
	}
	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Product{})
	connection.AutoMigrate(&models.Order{})
}
