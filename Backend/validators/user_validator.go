package validators

import (
	"github.com/pkg/errors"
	"github.com/szymon676/jobguru/storage"
	"github.com/szymon676/jobguru/types"
	"golang.org/x/crypto/bcrypt"
)

func VerifyRegisterReq(req types.RegisterUser) error {
	if len(req.Name) < 3 {
		return errors.New("name should be at least 3 characters long")
	}
	if len(req.Email) < 6 {
		return errors.New("email should be at least 6 characters long")
	}
	if len(req.Password) < 4 {
		return errors.New("password should be at least 4 characters long")
	}

	return nil
}

func VerifyLogin(req types.LoginUser, storage storage.UserStorager) error {
	account, err := storage.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password))
	if err != nil {
		return errors.New("wrong password!")
	}

	return nil
}
