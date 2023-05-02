package service

import (
	"github.com/szymon676/jobguru/api/auth"
	"github.com/szymon676/jobguru/storage"
	"github.com/szymon676/jobguru/types"
	"github.com/szymon676/jobguru/validators"

	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(types.RegisterUser) error
	LoginUser(types.LoginUser) (string, error)
	GetUserByID(id int) (*types.User, error)
}

type UserServie struct {
	storage storage.UserStorager
}

func NewUserService(storage storage.UserStorager) *UserServie {
	return &UserServie{
		storage: storage,
	}
}

func (us *UserServie) CreateUser(req types.RegisterUser) error {
	if err := validators.VerifyRegisterReq(req); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)

	if err := us.storage.CreateUser(req); err != nil {
		return err
	}
	return nil
}

func (us *UserServie) LoginUser(req types.LoginUser) (string, error) {
	err := validators.VerifyLogin(req, us.storage)
	if err != nil {
		return "", err
	}
	token, err := auth.CreateJWT(&req, us.storage)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *UserServie) GetUserByID(id int) (*types.User, error) {
	user, err := us.storage.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
