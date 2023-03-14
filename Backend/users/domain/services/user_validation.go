package services

import (
	"fmt"
	"strconv"

	"github.com/szymon676/job-guru/users/domain/models"
	"github.com/szymon676/job-guru/users/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

func VerifyRegister(v models.RegisterUser) error {
	if len(v.Name) <= 3 || len(v.Email) <= 6 || len(v.Password) <= 4 {
		return fmt.Errorf("error binding registration")
	}

	return nil
}

func VerifyLogin(v models.LoginUser) error {
	id, _ := strconv.Atoi(v.ID)
	account, _ := repository.GetUserByID(id)

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(v.Password))
	if err != nil {
		return fmt.Errorf("wrong password!")
	}

	return nil
}
