package validation

import (
	"fmt"

	"github.com/szymon676/job-guru/jobs/internal/models"
)

func VerifyJSON(bj models.BindJob) error {
	if bj.Category == "" || bj.Title == "" || bj.Description == "" || bj.Skills == nil {
		return fmt.Errorf("error binding job")
	}

	if len(bj.Skills) < 1 || len(bj.Category) < 3 || len(bj.Title) < 3 || len(bj.Description) < 10 {
		return fmt.Errorf("error binding job")
	}
	return nil
}