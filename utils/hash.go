package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}



	return string(bcryptPassword), err
}


func ComparePasswords(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return nil == err
}