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

	result, error := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID) // passing in the actual values for the placeholders in the query
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

func GetAllEvents() ([]Event, error) {
	// get all events from database

	query := `SELECT * FROM events;`
	rows, err := db.DB.Query(query) // query the database and get all the events from the events table
	if err != nil {
		return nil , err
	}

	defer rows.Close() // close the rows after the function ends to free up resources

	events := []Event{}

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) // scan the rows and assign the values to the event struct, the order of the values should match the order of the columns in the query & pointer to the struct fields
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}