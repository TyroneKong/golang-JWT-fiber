package handlers


import (
	"errors"
	"learnfiber/database"
	"learnfiber/helpers"
	"learnfiber/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	Role     int    `json:"role"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, Name: userModel.Name, Username: userModel.Username, Email: userModel.Email, Password: userModel.Password, Role: userModel.Role}
}

func HandleCreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.DB.Create(&user)
	response := CreateResponseUser(user)

	return c.Status(200).JSON(response)
}

func HandleSetRole(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := helpers.AuthUser(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err != nil {
		return c.Status(400).JSON("no user matching that id")
	}
	type UpdateUser struct {
		Role int `json:"role"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	user.Role = updateData.Role

	database.DB.Save(&user)

	return c.Status(200).JSON(&updateData)
}

func GetAllUsers() ([]User, error) {

	var u []User

	if err := database.DB.Find(&u).Error; err != nil {
		return u, errors.New("no users found")
	}
	return u, nil

}

func HandleAllUsers(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := helpers.AuthUser(cookie)

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

func FindUser(id int, user *models.User) error {

	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		return errors.New("user does not exist")
	}

	return nil
}

func HandleGetUserById(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := helpers.AuthUser(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(user)

}

func HandleUpdateUser(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := helpers.AuthUser(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}
	if err := FindUser(id, &user); err != nil {
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
	user.Username = updateData.Username
	user.Email = updateData.Email

	database.DB.Save(&user)

	return c.Status(200).JSON(&updateData)
}
