package database

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/szymon676/job-guru/jobs/internal/models"
)

func CreateJob(title, category, description string, skills []string) error {
	query := "INSERT INTO jobs (title, category, skills, description) VALUES($1, $2, $3, $4)"
	convskills := pq.Array(skills)

	_, err := DB.Query(query, title, category, convskills, description)
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

		err := rows.Scan(&job.ID, &job.Title, (*pq.StringArray)(&job.Skills), &job.Category, &job.Description)
		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	return jobs, nil
}

func UpdateJob(ID int, title string, skills []string, category, description string) error {
	var count int

	if err := DB.QueryRow("SELECT COUNT(*) FROM jobs WHERE id = $1", ID).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("job with id %v does not exist", ID)
	}

	query := "UPDATE jobs SET title = $1, skills = $2, category = $3, description = $4 WHERE id = $5;"
	_, err := DB.Exec(query, title, pq.Array(skills), category, description, ID)
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
