package controller

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	Email string `json:"user"`
	jwt.StandardClaims
}

func CreateToken(email string) (string, error) {
	claims := claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
