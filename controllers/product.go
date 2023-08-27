package controllers

import (
	"errors"
	"learnfiber/database"
	"learnfiber/models"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func FindProduct(id int, product *models.Product) error {
	database.DB.Find(&product, "id = ?", id)

	if product.ID == 0 {
		return errors.New("Product does not exist")
	}
	return nil

}

func CreateResponseProduct(product models.Product) Product {
	return Product{ID: product.ID, Name: product.Name, SerialNumber: product.SerialNumber}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Create(&product)
	response := CreateResponseProduct(product)
	return c.Status(200).JSON(response)
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}
	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err != nil {
		return c.Status(400).JSON("no user matching that id")
	}

	return c.Status(200).JSON(product)

}

func GetAllProducts() ([]Product, error) {

	var p []Product

	if err := database.DB.Find(&p).Error; err != nil {
		return p, errors.New("no users found")
	}
	return p, nil

}

func AllProducts(c *fiber.Ctx) error {

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

	p, err := GetAllProducts()

	if err != nil {
		return c.Status(400).JSON("No products available")
	}

	return c.Status(200).JSON(&p)
}
