package usecase

import (
	"github.com/szymon676/jobguru/internal/entity"
)

type (
	Job interface {
		CreateJob(*entity.JobReq) error
		GetJobs() ([]entity.Job, error)
		GetJobsByUser(userid int) ([]entity.Job, error)
		UpdateJob(int, *entity.JobReq) error
		DeleteJob(ID int) error
	}
	JobRepo interface {
		CreateJob(*entity.Job) error
		GetJobs() ([]entity.Job, error)
		GetJobsByUserID(userid int) ([]entity.Job, error)
		UpdateJob(int, *entity.Job) error
		DeleteJob(int) error
	}
	User interface {
		CreateUser(entity.RegisterUser) error
		LoginUser(entity.LoginUser) (int, error)
	}
	UserRepo interface {
		CreateUser(entity.RegisterUser) error
		GetUserByCriterion(criterion, value string) (*entity.User, error)
	}
)
