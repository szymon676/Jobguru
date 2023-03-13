package validation

import (
	"fmt"

	"github.com/szymon676/job-guru/users/internal/models"
)

func VerifyRegister(v models.RegisterUser) error {
	if len(v.Name) <= 3 || len(v.Email) <= 6 || len(v.Password) <= 4 {
		return fmt.Errorf("error binding registration")
	}

	return nil
}
