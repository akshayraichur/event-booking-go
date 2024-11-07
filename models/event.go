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
	UserID      int64       `json:"userId"`
}

func (event *Event) Save() error {
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

	// TODO: attach userId here

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

func GetEventById(id int64) (*Event, error) {
	// get event by id from database

	query := `SELECT * FROM events WHERE id = ?;` // ? is a placeholder for the actual values that we will pass in later, this is to prevent SQL injection attacks
	row := db.DB.QueryRow(query, id) // query the database and get the event with the id from the events table
	
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID) // scan the row and assign the values to the event struct, the order of the values should match the order of the columns in the query & pointer to the struct fields
	
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEvent() error {
	// update event in database

	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, date = ?, user_id = ?
	WHERE id = ?;` // ? is a placeholder for the actual values that we will pass in later

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close() // close the statement after the function ends to free up resources 

	_, error := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID) // passing in the actual values for the placeholders in the query
	if error != nil {
		return error
	}

	return nil
}

func (event Event) DeleteEvent() error {
	// delete event from database

	query := `DELETE FROM events WHERE id = ?;` // ? is a placeholder for the actual values that we will pass in later

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close() // close the statement after the function ends to free up resources 

	_, error := stmt.Exec(event.ID) // passing in the actual values for the placeholders in the query
	if error != nil {
		return error
	}

	return nil


}