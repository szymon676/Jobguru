package repository

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/szymon676/job-guru/jobs/domain/models"
)

func CreateJob(title, company string, skills []string, salary int, description, currency string) error {
	query := "INSERT INTO jobs (title, company, skills, salary, description, currency) VALUES($1, $2, $3, $4, $5, $6)"
	convskills := pq.Array(skills)

	_, err := DB.Query(query, title, company, convskills, salary, description, currency)
	if err != nil {
		return fmt.Errorf("insert into jobs failed: %v", err)
	}

	return nil
}

func GetJobs() ([]models.Job, error) {
	query := "SELECT * FROM jobs;"

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job

		err := rows.Scan(&job.ID, &job.Title, &job.Company, (*pq.StringArray)(&job.Skills), &job.Salary, &job.Description, &job.Currency)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func UpdateJob(ID int, title, company string, skills []string, salary int, description, currency string) error {
	var count int

	if err := DB.QueryRow("SELECT COUNT(*) FROM jobs WHERE id = $1", ID).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("job with id %v does not exist", ID)
	}

	query := "UPDATE jobs SET title = $1, company = $2, skills = $3, salary = $4, description = $5, currency = $6 WHERE id = $7;"
	_, err := DB.Exec(query, title, company, pq.Array(skills), salary, description, currency, ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteJob(ID string) error {
	query := "DELETE FROM jobs WHERE id = $1"

	_, err := DB.Exec(query, ID)

	if err != nil {
		return err
	}
	return nil
}
