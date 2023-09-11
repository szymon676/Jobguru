package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/jobguru/internal/entity"
	"github.com/szymon676/jobguru/internal/usecase"
)

type JobRoutes struct {
	UseCase usecase.Job
}

func newJobRoutes(router *fiber.App, jobUseCase usecase.Job) {
	jobRoutes := JobRoutes{UseCase: jobUseCase}

	jobGroup := router.Group("/jobs")
	jobGroup.Post("/", jobRoutes.CreateJob)
	jobGroup.Get("/", jobRoutes.GetAllJobs)
	jobGroup.Get("/:userid", jobRoutes.GetJobsByUser)
	jobGroup.Put("/:id", jobRoutes.UpdateJob)
	jobGroup.Delete("/:id", jobRoutes.DeleteJob)
}

func (jr *JobRoutes) CreateJob(c *fiber.Ctx) error {
	var req *entity.JobReq
	c.BodyParser(&req)

	if err := jr.UseCase.CreateJob(req); err != nil {
		return err
	}

	return c.JSON("Job created successfully")
}

func (jr *JobRoutes) GetAllJobs(c *fiber.Ctx) error {
	jobs, err := jr.UseCase.GetJobs()
	if err != nil {
		return err
	}

	return c.JSON(jobs)
}

func (jr *JobRoutes) GetJobsByUser(c *fiber.Ctx) error {
	userID := c.Params("userid")
	id, _ := strconv.Atoi(userID)

	jobs, err := jr.UseCase.GetJobsByUser(id)
	if err != nil {
		return err
	}

	return c.JSON(jobs)
}

func (jr *JobRoutes) UpdateJob(c *fiber.Ctx) error {
	var req *entity.JobReq
	jobID := c.Params("id")
	id, _ := strconv.Atoi(jobID)

	c.BodyParser(&req)

	if err := jr.UseCase.UpdateJob(id, req); err != nil {
		return err
	}

	return c.JSON("Job updated successfully")
}

func (jr *JobRoutes) DeleteJob(c *fiber.Ctx) error {
	jobID := c.Params("id")
	id, _ := strconv.Atoi(jobID)

	if err := jr.UseCase.DeleteJob(id); err != nil {
		return err
	}

	return c.JSON("Job deleted successfully")
}
