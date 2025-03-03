package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		fmt.Printf("Could not connect to database: %v", err)
	}

	if DB == nil {
		panic("Database connection is nil")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	if DB == nil {
		panic("Database connection is nil, Cannot create table")
	}
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL UNIQUE );`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		fmt.Printf("Could not create user table: %s", err)
		panic("Could not create event table")
	}
	println("Users table Created Successfully!")

	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	);`

	_, err = DB.Exec(createEventTable)

	if err != nil {
		fmt.Printf("Could not create event table: %s", err)
		panic("Could not create event table")
	}
	println("Events table Created Successfully!")

	createRegistrationsTable := `CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id),
	FOREIGN KEY(event_id) REFERENCES events(id));`

	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		fmt.Printf("Could not create registrations table: %s", err)
		panic("Could not create registrations table")
	}
	println("Registerations table Created Successfully!")
}
