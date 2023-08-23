package database

import (
	"fmt"
	"learnfiber/models"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// taken from the .env to be able to use in the connection string below
	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	//connection string
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("cannot connect to database", Dbdriver)
		log.Fatal("connection error", err)
	} else {
		fmt.Println("We are connected to the database", Dbdriver)
	}
	// create tables for users, products and orders
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
}

