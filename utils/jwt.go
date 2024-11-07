package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretKey"

func GenerateToken(email string, userID int64) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	return jwtToken.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	// verify token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // type checking syntax

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Invalid token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims) // type assertion syntax

	if !ok{
		return 0, errors.New("Invalid token")
	}

	// email := claims["email"].(string)
	userID := int64(claims["userId"].(float64))

	if err != nil {
		fmt.Println("Error converting str to int64")
		return 0, errors.New("Err converting str to int")
	}
	fmt.Println("User id: ", userID)

	return userID, nil

}