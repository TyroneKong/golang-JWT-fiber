package controllers

import (
	"learnfiber/database"
	"learnfiber/models"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
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
