package storage

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/szymon676/job-guru/jobs/types"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgreStorage() *PostgresStorage {
	return &PostgresStorage{
		db: DB,
	}
}

func (ps PostgresStorage) CreateJob(userid uint, title, company string, skills []string, salary int, description, currency, dateStr, location string) error {
	query := "INSERT INTO jobs (user_id, title, company, skills, salary, description, currency, date, location) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	convskills := pq.Array(skills)

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("failed to parse date: %v", err)
	}

	_, err = ps.db.Exec(query, userid, title, company, convskills, salary, description, currency, date, location)
	if err != nil {
		return fmt.Errorf("insert into jobs failed: %v", err)
	}

	return nil
}

func (ps PostgresStorage) GetJobs() ([]types.Job, error) {
	query := "SELECT * FROM jobs;"

	rows, err := ps.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []types.Job
	for rows.Next() {
		job, err := scanUser(rows)

		if err != nil {
			return nil, err
		}

		jobs = append(jobs, *job)
	}

	return jobs, nil
}

func (ps PostgresStorage) GetJobsByUser(userid uint) ([]types.Job, error) {
	query := "SELECT * FROM jobs WHERE user_id = $1;"

	rows, err := ps.db.Query(query, userid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []types.Job
	for rows.Next() {
		job, err := scanUser(rows)

		if err != nil {
			return nil, err
		}

		jobs = append(jobs, *job)
	}

	return jobs, nil
}

func (ps PostgresStorage) UpdateJob(ID int, userid uint, title, company string, skills []string, salary int, description, currency, dateStr, location string) error {
	var count int

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("failed to parse date: %v", err)
	}

	if err := ps.db.QueryRow("SELECT COUNT(*) FROM jobs WHERE id = $1", ID).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("job with id %v does not exist", ID)
	}

	query := "UPDATE jobs SET user_id = $1, title = $2, company = $3, skills = $4, salary = $5, description = $6, currency = $7, date = $8, location = $9 WHERE id = $10;"
	_, err = ps.db.Exec(query, userid, title, company, pq.Array(skills), salary, description, currency, date, location, ID)
	if err != nil {
		return err
	}

	return nil
}

func (ps PostgresStorage) DeleteJob(ID string) error {
	query := "DELETE FROM jobs WHERE id = $1"

	_, err := ps.db.Exec(query, ID)

	if err != nil {
		return err
	}
	return nil
}

func scanUser(rows *sql.Rows) (*types.Job, error) {
	job := new(types.Job)
	err := rows.Scan(
		&job.ID,
		&job.UserID,
		&job.Title,
		&job.Company,
		pq.Array(&job.Skills),
		&job.Salary,
		&job.Description,
		&job.Currency,
		&job.Date,
		&job.Location,
	)

	return job, err
}
