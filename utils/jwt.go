package utils

import (
	"github.com/ecea-nitt/ecea-server/config"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Email     string `json:"email"`
	ContactNo string `json:"contact"`
	jwt.StandardClaims
}

func CreateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JwtSecret))
}
