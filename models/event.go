package models

import (
	"time"

	"akshayraichur.com/event-booking-go/db"
)

type Event struct {
	ID          int64       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date"`
	UserID      int       `json:"userId"`
}

var events = []Event{}


func (event Event)Save() error {
	// save event to database

	query := `
	INSERT INTO events (name, description, location, date, user_id) 
	VALUES (?, ?, ?, ?, ?);` // ? is a placeholder for the actual values that we will pass in later

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close() // close the statement after the function ends to free up resources 

	result, error := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID) // passing in the actual values
	if error != nil {
		return error
	}
	id , err := result.LastInsertId()
	if err != nil {
		return err
	}
	event.ID = id

	return nil
}

func GetAllEvents() []Event {
	// get all events from database
	return events
}