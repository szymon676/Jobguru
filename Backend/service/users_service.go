package service

import (
	"log"

	jwt "github.com/szymon676/jobguru/api/jwt"
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

type UserService struct {
	storage storage.UserStorager
}

func NewUserService(storage storage.UserStorager) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (us *UserService) CreateUser(req types.RegisterUser) error {
	if err := validators.VerifyRegisterReq(req); err != nil {
		return err
	}
	log.Println(req.Email, req.Password, req.Name)
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

func (us *UserService) LoginUser(req types.LoginUser) (string, error) {
	err := validators.VerifyLogin(req, us.storage)
	if err != nil {
		return "", err
	}
	token, err := jwt.CreateJWT(&req, us.storage)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *UserService) GetUserByID(id int) (*types.User, error) {
	user, err := us.storage.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
