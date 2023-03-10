package repository

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func NewDatabase(driverName, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS jobs (
            id SERIAL PRIMARY KEY,
            title varchar(255),
			skills text[],
			category text,
			description text
		);
    `)
	DB = db
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
