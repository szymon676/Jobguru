package storage

import (
	"github.com/szymon676/jobguru/types"
)

type JobStorager interface {
	CreateJob(types.JobReq) error
	GetJobs() ([]types.Job, error)
	GetJobsByUserID(userid int) ([]types.Job, error)
	UpdateJobByID(int, types.JobReq) error
	DeleteJobByID(int) error
}
