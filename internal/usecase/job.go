package usecase

import (
	"github.com/szymon676/jobguru/internal/entity"
)

type JobUsecase struct {
	repo JobRepo
}

func NewJobUsecase(repo JobRepo) *JobUsecase {
	return &JobUsecase{
		repo: repo,
	}
}

func (js *JobUsecase) CreateJob(req *entity.JobReq) error {
	job, err := entity.VerifyJobReq(req)
	if err != nil {
		return err
	}

	err = js.repo.CreateJob(job)
	if err != nil {
		return err
	}

	return nil
}

func (js *JobUsecase) GetJobs() ([]entity.Job, error) {
	jobs, err := js.repo.GetJobs()
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (js *JobUsecase) GetJobsByUser(userid int) ([]entity.Job, error) {
	jobs, err := js.repo.GetJobsByUserID(userid)
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func (js *JobUsecase) UpdateJobByID(id int, req *entity.JobReq) error {
	job, err := entity.VerifyJobReq(req)
	if err != nil {
		return err
	}
	err = js.repo.UpdateJobByID(id, job)
	if err != nil {
		return err
	}
	return nil
}

func (js *JobUsecase) DeleteJobByID(jobID int) error {
	err := js.repo.DeleteJobByID(jobID)
	if err != nil {
		return err
	}
	return nil
}
