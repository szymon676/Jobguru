package storage

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"github.com/szymon676/job-guru/jobs/types"
)

type PostgresStorage struct {
}

func NewPostgreStorage() *PostgresStorage {
	return &PostgresStorage{}
}

func (PostgresStorage) CreateJob(title, company string, skills []string, salary int, description, currency, dateStr, location string) error {
	query := "INSERT INTO jobs (title, company, skills, salary, description, currency, date, location) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	convskills := pq.Array(skills)
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("failed to parse date: %v", err)
	}

	_, err = DB.Exec(query, title, company, convskills, salary, description, currency, date, location)
	if err != nil {
		return fmt.Errorf("insert into jobs failed: %v", err)
	}

	return nil
}

func (PostgresStorage) GetJobs() ([]types.Job, error) {
	query := "SELECT * FROM jobs;"

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []types.Job
	for rows.Next() {
		var job types.Job

		err := rows.Scan(&job.ID, &job.Title, &job.Company, (*pq.StringArray)(&job.Skills), &job.Salary, &job.Description, &job.Currency, &job.Date, &job.Location)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (PostgresStorage) UpdateJob(ID int, title, company string, skills []string, salary int, description, currency, dateStr, location string) error {
	var count int

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return fmt.Errorf("failed to parse date: %v", err)
	}

	if err := DB.QueryRow("SELECT COUNT(*) FROM jobs WHERE id = $1", ID).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("job with id %v does not exist", ID)
	}

	query := "UPDATE jobs SET title = $1, company = $2, skills = $3, salary = $4, description = $5, currency = $6, date = $7, location = $8 WHERE id = $9;"
	_, err = DB.Exec(query, title, company, pq.Array(skills), salary, description, currency, date, location, ID)
	if err != nil {
		return err
	}

	return nil
}

func (PostgresStorage) DeleteJob(ID string) error {
	query := "DELETE FROM jobs WHERE id = $1"

	_, err := DB.Exec(query, ID)

	if err != nil {
		return err
	}
	return nil
}
