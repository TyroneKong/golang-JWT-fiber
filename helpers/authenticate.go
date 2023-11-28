package helpers

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func AuthUser(cookie string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil

	})
}
