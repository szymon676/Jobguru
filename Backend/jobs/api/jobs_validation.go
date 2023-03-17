package api

import (
	"errors"

	"github.com/szymon676/job-guru/jobs/types"
)

func VerifyJSON(bj types.BindJob) error {
	if bj.UserID == 0 || len(bj.Skills) < 1 || len(bj.Company) < 3 || len(bj.Title) < 3 || len(bj.Description) < 10 || bj.Salary < 10 || bj.Currency == "" || len(bj.Date) < 5 || len(bj.Location) < 5 {
		return errors.New("error binding job")
	}

	return nil
}
