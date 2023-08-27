package controllers

import (
	"learnfiber/database"
	"learnfiber/models"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
}

func createResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product}
}

func CreateOrder(c *fiber.Ctx) error {

	var order models.Order

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

func GetAllOrders(c *fiber.Ctx) error {

	orders := []models.Order{}
	database.DB.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		database.DB.Find(&user, "id = ?", order.UserRefer)
		database.DB.Find(&product, "id = ?", order.ProductRefer)
		responseOrder := createResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)

	}
	return c.Status(200).JSON(responseOrders)
}
