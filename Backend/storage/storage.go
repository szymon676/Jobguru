package storage

import "database/sql"

func NewPostgreStorage(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
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
	return db, nil
}
