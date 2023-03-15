package storage

import (
	"database/sql"
	"fmt"

	"github.com/szymon676/job-guru/jobs/types"
)

var DB *sql.DB

type Storager interface {
	GetJobs() ([]types.Job, error)
	UpdateJob(ID int, title, company string, skills []string, salary int, description, currency, dateStr, location string) error
	DeleteJob(ID string) error
	CreateJob(title, company string, skills []string, salary int, description, currency, dateStr, location string) error
}

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
	DB = db
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
