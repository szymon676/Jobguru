package usecase

import (
	"errors"
	"log"

	"github.com/szymon676/jobguru/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (us *UserUsecase) CreateUser(req entity.RegisterUser) error {
	if err := entity.VerifyRegisterReq(req); err != nil {
		return err
	}
	log.Println(req.Email, req.Password, req.Name)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPassword)

	if err := us.repo.CreateUser(req); err != nil {
		return err
	}

	return nil
}

func (us *UserUsecase) LoginUser(req entity.LoginUser) (string, error) {
	account, err := us.repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("wrong password!")
	}

	return "123", nil
}

func (us *UserUsecase) GetUserByID(id int) (*entity.User, error) {
	user, err := us.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
