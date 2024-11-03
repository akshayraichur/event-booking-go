package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
func InitDB(){
	var err error
	
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10) // how many open connects we can have at a time to this database
	DB.SetMaxIdleConns(5) // how many idle connections we can have to the database

	createTables()
}

func createTables(){
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`

	_, userTableErr := DB.Exec(createUserTable)
	if userTableErr != nil {
		panic("could not create users table")
	}


	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT, 
		name TEXT NOT NULL, 
		description TEXT NOT NULL, 
		location TEXT NOT NULL, 
		date DATETIME NOT NULL, 
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create events table")
		fmt.Print(err)
	}

}