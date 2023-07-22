package utils

import (
	"github.com/dgrijalva/jwt-go"
	"gormtest/domains"
	"os"
)

func GenerateJWT(user domains.User) (string, error) {

	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil

}
