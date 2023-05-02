package validators

import (
	"github.com/pkg/errors"
	"github.com/szymon676/jobguru/storage"
	"github.com/szymon676/jobguru/types"
	"golang.org/x/crypto/bcrypt"
)

func VerifyRegisterReq(req types.RegisterUser) error {
	if len(req.Name) < 3 || len(req.Email) < 6 || len(req.Password) < 4 {
		return errors.New("eror binding registration!")
	}

	return nil
}

func VerifyLogin(req types.LoginUser, storage storage.UserStorager) error {
	account, _ := storage.GetUserByEmail(req.Email)

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password))
	if err != nil {
		return errors.New("wrong password!")
	}

	return nil
}
