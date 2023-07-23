package utils

import (
	"errors"
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
func ValidateToken(tokenString string) (float64, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}
	id := claims["user_id"].(float64)

	return id, nil
}
