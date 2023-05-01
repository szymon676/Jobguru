package utils

import (
	"github.com/pkg/errors"
	"github.com/szymon676/jobguru/storage"
	"github.com/szymon676/jobguru/types"
	"golang.org/x/crypto/bcrypt"
)

type VerifyService interface {
	VerifyRegister(v types.RegisterUser) error
	VerifyLogin(v types.LoginUser) error
}

type Verifier struct {
	storage storage.UserStorager
}

func NewVerifier(storage storage.UserStorager) *Verifier {
	return &Verifier{
		storage: storage,
	}
}

func (vs Verifier) VerifyRegister(v types.RegisterUser) error {
	if len(v.Name) < 3 || len(v.Email) < 6 || len(v.Password) < 4 {
		return errors.New("eror binding registration!")
	}

	return nil
}

func (vs Verifier) VerifyLogin(v types.LoginUser) error {
	account, _ := vs.storage.GetUserByEmail(v.Email)

	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(v.Password))
	if err != nil {
		return errors.New("wrong password!")
	}

	return nil
}
