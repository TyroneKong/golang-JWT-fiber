package helpers

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

//TODO: might be better to add this as middleware
func AuthUser(cookie string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil

	})
}
