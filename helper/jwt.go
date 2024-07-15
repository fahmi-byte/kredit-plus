package helper

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTSecret = []byte("secret_key")

type Claims struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	jwt.StandardClaims
}

func GenerateJWT(id int, name, email, phoneNumber string) (string, error) {
	claims := Claims{
		Id:          id,
		Name:        name,
		Email:       email,
		PhoneNumber: phoneNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
