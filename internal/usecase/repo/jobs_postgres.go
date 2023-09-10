package repo

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/szymon676/jobguru/internal/entity"
)

type JobRepo struct {
	db *sql.DB
}

func NewJobRepo(db *sql.DB) *JobRepo {
	return &JobRepo{
		db: db,
	}
}

func (jr *JobRepo) CreateJob(job *entity.Job) error {
	query := `
		INSERT INTO jobs (user_id, title, company, skills, salary, description, currency, date, location)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	convskills := pq.Array(job.Skills)

	_, err := jr.db.Exec(query, job.UserID, job.Title, job.Company, convskills, job.Salary, job.Description, job.Currency, job.Date, job.Location)
	if err != nil {
		return fmt.Errorf("insert into jobs failed: %v", err)
	}

	return nil
}

func (jr *JobRepo) getJobsByQuery(query string, args ...interface{}) ([]entity.Job, error) {
	rows, err := jr.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []entity.Job
	for rows.Next() {
		job, err := scanJob(rows)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, *job)
	}

	return jobs, nil
}

func (jr *JobRepo) GetJobs() ([]entity.Job, error) {
	query := "SELECT * FROM jobs"
	return jr.getJobsByQuery(query)
}

func (jr *JobRepo) GetJobsByUserID(userID int) ([]entity.Job, error) {
	query := "SELECT * FROM jobs WHERE user_id = $1"
	return jr.getJobsByQuery(query, userID)
}

func (jr *JobRepo) UpdateJob(ID int, job *entity.Job) error {
	count := 0

	if err := jr.db.QueryRow("SELECT COUNT(*) FROM jobs WHERE id = $1", ID).Scan(&count); err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("job with ID %v does not exist", ID)
	}

	convskills := pq.Array(job.Skills)

	query := `
		UPDATE jobs SET user_id = $1, title = $2, company = $3, skills = $4, salary = $5, 
		description = $6, currency = $7, date = $8, location = $9 WHERE id = $10
	`
	_, err := jr.db.Exec(query, job.UserID, job.Title, job.Company, convskills, job.Salary, job.Description, job.Currency, job.Date, job.Location, ID)
	if err != nil {
		return err
	}

	return nil
}

func (jr *JobRepo) DeleteJob(ID int) error {
	query := "DELETE FROM jobs WHERE id = $1"
	_, err := jr.db.Exec(query, ID)
	if err != nil {
		return err
	}
	return nil
}

func scanJob(rows *sql.Rows) (*entity.Job, error) {
	job := new(entity.Job)
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
