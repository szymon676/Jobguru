package entity

import (
	"errors"
	"time"
)

// Job represents a job entity.
type Job struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	Skills      []string  `json:"skills"`
	Salary      int       `json:"salary"`
	Description string    `json:"description"`
	Currency    string    `json:"currency"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
}

// JobReq represents a job request.
type JobReq struct {
	UserID      int      `json:"user_id"`
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Skills      []string `json:"skills"`
	Salary      int      `json:"salary"`
	Description string   `json:"description"`
	Currency    string   `json:"currency"`
	Location    string   `json:"location"`
}

// VerifyJobReq validates a job request and returns a Job or an error.
func VerifyJobReq(req *JobReq) (*Job, error) {
	if req.UserID <= 0 {
		return nil, errors.New("userID must be greater than 0")
	}
	if len(req.Skills) < 1 {
		return nil, errors.New("at least one skill is required")
	}
	if len(req.Company) < 3 {
		return nil, errors.New("company name must be at least 3 characters")
	}
	if len(req.Title) < 3 {
		return nil, errors.New("job title must be at least 3 characters")
	}
	if len(req.Description) < 10 {
		return nil, errors.New("job description must be at least 10 characters")
	}
	if req.Salary < 10 {
		return nil, errors.New("salary must be at least 10")
	}
	if req.Currency == "" {
		return nil, errors.New("currency must be specified")
	}
	if len(req.Location) < 5 {
		return nil, errors.New("location must be specified")
	}

	date := time.Now()

	job := &Job{
		UserID:      req.UserID,
		Skills:      req.Skills,
		Company:     req.Company,
		Title:       req.Title,
		Description: req.Description,
		Salary:      req.Salary,
		Currency:    req.Currency,
		Date:        date,
		Location:    req.Location,
	}

	return job, nil
}
