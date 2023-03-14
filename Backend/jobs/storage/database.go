package storage

import (
	"database/sql"
	"fmt"

	"github.com/szymon676/job-guru/jobs/types"
)

var DB *sql.DB

type Storager interface {
	GetJobs() ([]types.Job, error)
	UpdateJob(ID int, title, company string, skills []string, salary int, description, currency string) error
	DeleteJob(ID string) error
	CreateJob(title, company string, skills []string, salary int, description, currency string) error
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
            title varchar(255),
			company varchar(50),
			skills text[],
			salary integer,
			description text,
			currency varchar(5)
		);
    `)
	DB = db
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
