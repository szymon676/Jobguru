package app

import (
	"database/sql"
	"log"
	"os"
)

var sqldb *sql.DB

func init() {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            username text UNIQUE,
			password text,
			email text UNIQUE
		);
    `)

	db.Exec(`
	CREATE TABLE IF NOT EXISTS jobs (
		id SERIAL PRIMARY KEY,
		user_id INTEGER,
		title TEXT NOT NULL,
		company TEXT NOT NULL,
		skills TEXT[] NOT NULL,
		salary INTEGER,
		description TEXT,
		currency TEXT,
		date TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
		location TEXT
		);
	`)

	sqldb = db
}
