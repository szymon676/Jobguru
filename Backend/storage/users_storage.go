package storage

import (
	"github.com/szymon676/jobguru/types"
)

type UserStorager interface {
	CreateUser(name, password, email string) error
	GetUserByID(id int) (*types.User, error)
	GetUserByEmail(email string) (*types.User, error)
}
