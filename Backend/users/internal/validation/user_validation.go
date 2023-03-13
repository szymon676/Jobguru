package validation

import (
	"fmt"
	"strconv"

	"github.com/szymon676/job-guru/users/internal/database"
	"github.com/szymon676/job-guru/users/internal/models"
)

func VerifyRegister(v models.RegisterUser) error {
	if len(v.Name) <= 3 || len(v.Email) <= 6 || len(v.Password) <= 4 {
		return fmt.Errorf("error binding registration")
	}

	return nil
}

func ValidateUser(v models.LoginUser) error {
	id, _ := strconv.Atoi(v.ID)
	acc, _ := database.GetUserByID(id)

	if v.Password != acc.Password {
		return fmt.Errorf("wrong password!")
	}

	return nil
}
