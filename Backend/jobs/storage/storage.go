package storage

import (
	"database/sql"
	"fmt"

	"github.com/szymon676/job-guru/jobs/types"
)

var DB *sql.DB

type Storager interface {
	CreateJob(userid uint, title, company string, skills []string, salary int, description, currency, dateStr, location string) error
	GetJobsByUser(userid uint) ([]types.Job, error)
	GetJobs() ([]types.Job, error)
	UpdateJob(ID int, userid uint, title, company string, skills []string, salary int, description, currency, dateStr, location string) error
	DeleteJob(ID string) error
}

func NewPostgresDatabase(driverName, dsn string) (*sql.DB, error) {
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

	DB = db

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
