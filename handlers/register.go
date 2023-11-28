package handlers


import (
	"errors"
	"fmt"
	"learnfiber/database"
	"learnfiber/models"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password map[string]string, cost int) ([]byte, error) {
	// this ensures whats passed in actually is a password
	if _, ok := password["password"]; !ok {
		return nil, errors.New("password key not found in map")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password["password"]), cost)

	if err != nil {
		fmt.Println("Error:", err)
	}
	return hashed, nil
}

func compareHashpassword(password map[string]string, user models.User) error {
	// this ensures whats passed in actually is a password
	if _, ok := password["password"]; !ok {
		return errors.New("password key not found in map")
	}
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password["password"]))
	if err != nil {
		return fmt.Errorf("Password comparisson has failed")
	}
	return nil
}

func HandleRegister(c *fiber.Ctx) error {
	var data map[string]string
	//we map the request to data using BodyParser and check for errors

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// we hash the password usin bcrypt
	// needed to convert password to byte array as func does not accept string
	password, _ := hashPassword(data, 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Username: data["username"],
		Password: password,
	}
	database.DB.Create(&user)
	return c.JSON(user)
}

func HandleLogin(c *fiber.Ctx) error {
	var data map[string]string
	//we map the request to data using BodyParser
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}
	// we compare the hashed password
	if err := compareHashpassword(data, user); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON((fiber.Map{
			"message": "incorrect password",
		}))
	}
	// we create claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	// create a token
	token, err := claims.SignedString([]byte(os.Getenv("API_SECRET")))

	if err != nil {
		return err
	}
	// set the token in a http only cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Path:     "/",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 48),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func HandleCurrentUser(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {

		// checking that the jwt was signed with the correct method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected sign in method")
		}

		return []byte(os.Getenv("API_SECRET")), nil

	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func HandleLogout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie((&cookie))

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
