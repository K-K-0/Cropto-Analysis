package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := ""
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("database connection Error: ", err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS trades (
			id SERIAL PRIMARY KEY
			symbol TEXT
			price TEXT
			quantity TEXT
			timestamp TIMESTAMP

		)
	`)
	if err != nil {
		log.Fatal("error while creating Database trade table: ", err)
	}

	log.Println("Database is connected successfully")
}
