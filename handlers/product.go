package handlers

import (
	"errors"
	"learnfiber/database"
	"learnfiber/models"

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

func HandleCreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Create(&product)
	response := CreateResponseProduct(product)
	return c.Status(200).JSON(response)
}

const (
	Admin = iota + 1
	Internal
)

func HandleDeleteProduct(c *fiber.Ctx) error {

	var product models.Product
	var user models.User

	userId, err := c.ParamsInt(("userId"))
	id, err := c.ParamsInt("id")

	if err := FindUser(userId, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if user.Role != Admin {
		return c.Status(401).JSON("You are not authorized to delete products")
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err != nil {
		return c.Status((400)).JSON("no product matching that id")
	}

	database.DB.Where("id = ?", id).Delete(&product)
	response := CreateResponseProduct(product)
	return c.Status(200).JSON(response)
}

func HandleGetProductById(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}
	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := CreateResponseProduct(product)
	return c.Status(200).JSON(response)

}

func GetAllProducts() ([]Product, error) {

	var p []Product

	if err := database.DB.Find(&p).Error; err != nil {
		return p, errors.New("no users found")
	}
	return p, nil

}

func HandleAllProducts(c *fiber.Ctx) error {

	p, err := GetAllProducts()

	if err != nil {
		return c.Status(400).JSON("No products available")
	}

	return c.Status(200).JSON(&p)
}
