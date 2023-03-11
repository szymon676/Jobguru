package database

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/szymon676/job-guru/jobs/internal/models"
)

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

		err := rows.Scan(&job.ID, &job.Title, (*pq.StringArray)(&job.Skills), &job.Category, &job.Description)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func CreateJob(title, category, description string, skills []string) error {
	query := "INSERT INTO jobs (title, category, skills, description) VALUES($1, $2, $3, $4)"
	convskills := pq.Array(skills)

	_, err := DB.Query(query, title, category, convskills, description)
	if err != nil {
		return fmt.Errorf("insert into jobs failed: %v", err)
	}

	return nil
}
