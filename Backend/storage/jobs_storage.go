package storage

import (
	"github.com/szymon676/jobguru/types"
)

type JobStorager interface {
	CreateJob(userid uint, title, company string, skills []string, salary int, description, currency, dateStr, location string) error
	GetJobsByUser(userid uint) ([]types.Job, error)
	GetJobs() ([]types.Job, error)
	UpdateJob(ID int, userid uint, title, company string, skills []string, salary int, description, currency, dateStr, location string) error
	DeleteJob(ID string) error
}
