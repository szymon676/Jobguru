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

	existingUser, err := us.repo.GetUserByCriterion("email", req.Email)
	if err == nil && existingUser != nil {
		return errors.New("user with this email already exists")
	}

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

func (us *UserUsecase) LoginUser(req entity.LoginUser) (int, error) {
	user, err := us.repo.GetUserByCriterion("email", req.Email)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return 0, errors.New("wrong password!")
	}

	return user.ID, nil
}
