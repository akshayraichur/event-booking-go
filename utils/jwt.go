package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretKey"

func GenerateToken(email, userID string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	return jwtToken.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	// verify token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // type checking syntax

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("Invalid token")
	}

	if !parsedToken.Valid {
		return errors.New("Invalid token")
	}

	// claims, ok := parsedToken.Claims.(jwt.MapClaims) // type assertion syntax

	// if !ok{
	// 	return errors.New("Invalid token")
	// }

	// email := claims["email"].(string)
	// userID := claims["userId"].(int64)

	return nil

}