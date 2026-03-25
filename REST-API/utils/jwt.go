package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var keySecrete string = "jjwsjwiu982797"

func Generatetoken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(keySecrete))
}

func ValidateToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid 1 token")
		}
		return []byte(keySecrete), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, errors.New("invalid 2 token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid 3 token")
	}
	idFloat, ok := claims["id"].(float64)
	if !ok {
		return 0, errors.New("invalid id type")
	}

	id := int64(idFloat)
	return id, nil
}
