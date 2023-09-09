package entity

import "errors"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegisterUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func VerifyRegisterReq(req RegisterUser) error {
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
