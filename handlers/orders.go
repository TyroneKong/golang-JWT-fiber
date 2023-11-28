package handlers


import (
	"errors"
	"learnfiber/database"
	"learnfiber/helpers"
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

func HandleCreateOrder(c *fiber.Ctx) error {

	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.JSON(err.Error())
	}

	var user models.User
	if err := FindUser(order.UserRefer, &user); err != nil {
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

func HandleGetAllOrders(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := helpers.AuthUser(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

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

func findOrder(id int, order *models.Order) error {
	database.DB.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil

}

func HandleGetOrderById(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := helpers.AuthUser(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := FindUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := FindProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	response := createResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(response)

}

// Add a new endpoint to get orders by user ID
func HandleGetOrdersByUser(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := helpers.AuthUser(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(400).JSON("Please ensure user_id is an integer")
	}

	// Find orders by user ID
	orders := []models.Order{}
	database.DB.Find(&orders, "user_refer = ?", userID)

	if len(orders) == 0 {
		return c.Status(404).JSON("No orders found for the given user ID")
	}

	responseOrders := []Order{}

	// Retrieve user and product information for each order
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

func HandleDeleteOrder(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := helpers.AuthUser(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status((400)).JSON("no order matcing that id")
	}

	if err := findOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Where("id = ?", id).Delete(&order)
	return c.Status(200).JSON(order)
}
