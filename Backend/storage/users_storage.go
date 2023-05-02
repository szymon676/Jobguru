package storage

import (
	"github.com/szymon676/jobguru/types"
)

type UserStorager interface {
	CreateUser(types.RegisterUser) error
	GetUserByID(id int) (*types.User, error)
	GetUserByEmail(email string) (*types.User, error)
}
