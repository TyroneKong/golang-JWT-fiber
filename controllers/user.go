package controllers

import (
	"learnfiber/models"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `gorm:"size:255;not null;unique" json:"username"`
	Password  string `gorm:"size:255;not null;" json:"password"`
}

// func CreateResponseUser(userModel models.User) User {
// 	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName, Username: userModel.FirstName, Password: userModel.FirstName}
// }

// func CreateUser(c *fiber.Ctx) error {
// 	var user models.User

// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(400).JSON(err.Error())
// 	}
// 	// database.DB.Create(&user)
// 	response := CreateResponseUser(user)

// 	return c.Status(200).JSON(response)
// }

func GetAllUsers() ([]User, error) {
	var u []User

	// if err := database.DB.Find(&u).Error; err != nil {
	// 	return u, errors.New("no users found")
	// }
	return u, nil

}

func AllUsers(c *fiber.Ctx) error {

	u, err := GetAllUsers()

	if err != nil {
		return c.Status(400).JSON("No users available")
	}

	return c.Status(200).JSON(&u)
}

func findUser(id int, user *models.User) error {

	// if err := database.DB.Find(&user, "id = ?", id).Error; err != nil {
	// 	return errors.New("user does not exist")
	// }

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
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
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
