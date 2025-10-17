package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := "postgresql://neondb_owner:npg_g7HJw4zNKOqE@ep-long-block-adabnpad-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"
	DB, err := sql.Open("postgre", connStr)
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
