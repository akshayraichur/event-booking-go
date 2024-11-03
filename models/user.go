package models

import (
	"akshayraichur.com/event-booking-go/db"
	"akshayraichur.com/event-booking-go/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user User) Save() error {
	// save user to database

	query := `
	INSERT INTO users (name, email, password)
	VALUES (?, ?, ?);`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, error := stmt.Exec(user.Name, user.Email, hashedPassword)

	if error != nil {
		return error
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = userId
	return nil

}
