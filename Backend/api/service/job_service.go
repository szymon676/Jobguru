package service

import (
	"github.com/szymon676/jobguru/api/validators"
	"github.com/szymon676/jobguru/storage"
	"github.com/szymon676/jobguru/types"
)

type IJobService interface {
	CreateJob(types.JobReq) error
	GetJobs() ([]types.Job, error)
	GetJobsByUser(userid int) ([]types.Job, error)
	UpdateJobByID(int, types.JobReq) error
	DeleteJobByID(ID int) error
}

type JobService struct {
	storage storage.JobStorager
}

func NewJobService(storage storage.JobStorager) *JobService {
	return &JobService{
		storage: storage,
	}
}

func (js *JobService) CreateJob(req types.JobReq) error {
	err := validators.VerifyJobReq(req)
	if err != nil {
		return err
	}

	err = js.storage.CreateJob(req)
	if err != nil {
		return err
	}

	return nil
}

func (js *JobService) GetJobs() ([]types.Job, error) {
	jobs, err := js.storage.GetJobs()
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (js *JobService) GetJobsByUser(userid int) ([]types.Job, error) {
	jobs, err := js.storage.GetJobsByUserID(userid)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (js *JobService) UpdateJobByID(userID int, req types.JobReq) error {
	err := validators.VerifyJobReq(req)
	if err != nil {
		return err
	}
	err = js.storage.UpdateJobByID(userID, req)
	if err != nil {
		return err
	}
	return nil
}

func (js *JobService) DeleteJobByID(jobID int) error {
	err := js.storage.DeleteJobByID(jobID)
	if err != nil {
		return err
	}
	return nil
}
