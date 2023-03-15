package api

import (
	"fmt"

	"github.com/szymon676/job-guru/jobs/types"
)

func VerifyJSON(bj types.BindJob) error {
	if len(bj.Skills) < 1 || len(bj.Company) < 3 || len(bj.Title) < 3 || len(bj.Description) < 10 || bj.Salary < 10 || len(bj.Date) < 5 || len(bj.Location) < 5 {
		return fmt.Errorf("error binding job")
	}

	return nil
}
