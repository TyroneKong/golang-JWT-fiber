package controllers

import (
	"errors"
	"learnfiber/database"
	"learnfiber/models"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID           uint    `json:"id" gorm:"primaryKey"`
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
}

func createResponseOrder(order Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product}
}

func CreateOrder(c *fiber.Ctx) error {

	var order Order

	if err := c.BodyParser(&order); err != nil {
		return c.JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := FindProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	response := createResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(response)
}

func GetAllOrders() ([]Order, error) {

	var o []Order

	if err := database.DB.Find(&o).Error; err != nil {
		return o, errors.New("no users found")
	}
	return o, nil

}

func AllOrders(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil

	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	o, err := GetAllOrders()

	if err != nil {
		return c.Status(400).JSON("No orders available")
	}

	return c.Status(200).JSON(&o)

}
