package migrate

import (
	"database/sql"
	"log"
)

func MigratePostgresDB(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("couldn't open db.", err)
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
	return db
}
