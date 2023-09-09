package entity

import (
	"errors"
	"time"
)

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

type JobReq struct {
	UserID      int      `json:"user_id"`
	Title       string   `json:"title"`
	Company     string   `json:"company"`
	Skills      []string `json:"skills"`
	Salary      int      `json:"salary"`
	Description string   `json:"description"`
	Currency    string   `json:"currency"`
	Date        string   `json:"date"`
	Location    string   `json:"location"`
}

func VerifyJobReq(req JobReq) error {
	if req.UserID == 0 {
		return errors.New("userID must be greater than 0")
	}
	if len(req.Skills) < 1 {
		return errors.New("at least one skill is required")
	}
	if len(req.Company) < 3 {
		return errors.New("company name must be at least 3 characters")
	}
	if len(req.Title) < 3 {
		return errors.New("job title must be at least 3 characters")
	}
	if len(req.Description) < 10 {
		return errors.New("job description must be at least 10 characters")
	}
	if req.Salary < 10 {
		return errors.New("salary must be at least 10")
	}
	if req.Currency == "" {
		return errors.New("currency must be specified")
	}
	if len(req.Date) < 5 {
		return errors.New("date must be specified")
	}
	if len(req.Location) < 5 {
		return errors.New("location must be specified")
	}

	return nil
}
