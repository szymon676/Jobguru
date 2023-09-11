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
	Location    string   `json:"location"`
}

func VerifyJobReq(req *JobReq) (*Job, error) {
	validationErrors := []error{
		validateUserID(req.UserID),
		validateSkills(req.Skills),
		validateCompanyName(req.Company),
		validateJobTitle(req.Title),
		validateJobDescription(req.Description),
		validateSalary(req.Salary),
		validateCurrency(req.Currency),
		validateLocation(req.Location),
	}

	for _, err := range validationErrors {
		if err != nil {
			return nil, err
		}
	}

	return createJobFromRequest(req), nil
}

func validateUserID(userID int) error {
	if userID <= 0 {
		return errors.New("userID must be greater than 0")
	}
	return nil
}

func validateSkills(skills []string) error {
	if len(skills) < 1 {
		return errors.New("at least one skill is required")
	}
	return nil
}

func validateCompanyName(companyName string) error {
	if len(companyName) < 3 {
		return errors.New("company name must be at least 3 characters")
	}
	return nil
}

func validateJobTitle(jobTitle string) error {
	if len(jobTitle) < 3 {
		return errors.New("job title must be at least 3 characters")
	}
	return nil
}

func validateJobDescription(jobDescription string) error {
	if len(jobDescription) < 10 {
		return errors.New("job description must be at least 10 characters")
	}
	return nil
}

func validateSalary(salary int) error {
	if salary == 0 {
		return errors.New("salary must be specified")
	}
	return nil
}

func validateCurrency(currency string) error {
	if currency == "" {
		return errors.New("currency must be specified")
	}
	return nil
}

func validateLocation(location string) error {
	if len(location) < 1 {
		return errors.New("location must be specified")
	}
	return nil
}

func createJobFromRequest(req *JobReq) *Job {
	date := time.Now()

	return &Job{
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
}
