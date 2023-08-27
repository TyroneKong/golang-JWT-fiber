package controllers

import (
	"errors"
	"learnfiber/database"
	"learnfiber/models"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, Name: userModel.Name, Username: userModel.Username, Email: userModel.Email, Password: userModel.Password}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Create(&user)
	response := CreateResponseUser(user)

	return c.Status(200).JSON(response)
}

func GetAllUsers() ([]User, error) {

	var u []User

	if err := database.DB.Find(&u).Error; err != nil {
		return u, errors.New("no users found")
	}
	return u, nil

}

func AllUsers(c *fiber.Ctx) error {

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

	u, err := GetAllUsers()

	if err != nil {
		return c.Status(400).JSON("No users available")
	}

	return c.Status(200).JSON(&u)
}

func findUser(id int, user *models.User) error {

	if err := database.DB.Find(&user, "id = ?", id).Error; err != nil {
		return errors.New("user does not exist")
	}

	return nil
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err != nil {
		return c.Status(400).JSON("no user matching that id")
	}

	return c.Status(200).JSON(user)

}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err != nil {
		return c.Status(400).JSON("no user matching that id")
	}
	type UpdateUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	// user.FirstName = updateData.FirstName
	// user.LastName = updateData.LastName

	// database.DB.Save(&user)

	return c.Status(200).JSON(&updateData)
}
