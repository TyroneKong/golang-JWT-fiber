package controllers

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

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// database.DB.Create(&product)
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