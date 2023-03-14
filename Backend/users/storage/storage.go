package storage

import (
	"github.com/szymon676/job-guru/users/types"
)

type Storager interface {
	CreateUser(name, password, email string) error
	GetUserByID(id int) (*types.User, error)
}
