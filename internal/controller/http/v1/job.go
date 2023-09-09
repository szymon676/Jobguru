package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/szymon676/jobguru/internal/entity"
	"github.com/szymon676/jobguru/internal/usecase"
)

type jobRoutes struct {
	j usecase.Job
}

func newJobRoutes(router *fiber.App, jobusecase usecase.Job) {
	r := jobRoutes{j: jobusecase}

	jr := router.Group("/jobs")
	jr.Post("/", r.createJob)
	jr.Get("/", r.getAllJobs)
	jr.Get("/{userid}", r.getJobsByUser)
	jr.Put("/:id", r.updateJob)
	jr.Delete("/:id", r.deleteJob)
}

func (jh *jobRoutes) createJob(c *fiber.Ctx) error {
	var req entity.JobReq

	c.BodyParser(&req)

	if err := jh.j.CreateJob(req); err != nil {
		return err
	}

	return c.JSON("job created successfully")
}

func (jh *jobRoutes) getAllJobs(c *fiber.Ctx) error {
	jobs, err := jh.j.GetJobs()
	if err != nil {
		return err
	}

	return c.JSON(jobs)
}

func (jh *jobRoutes) getJobsByUser(c *fiber.Ctx) error {
	userid := c.Params("id")
	id, _ := strconv.Atoi(userid)

	jobs, err := jh.j.GetJobsByUser(id)
	if err != nil {
		return err
	}

	return c.JSON(jobs)
}

func (jh *jobRoutes) updateJob(c *fiber.Ctx) error {
	var req entity.JobReq
	userid := c.Params("id")
	id, _ := strconv.Atoi(userid)

	c.BodyParser(&req)

	err := jh.j.UpdateJobByID(id, req)
	if err != nil {
		return err
	}

	return c.JSON("job updated successfully")
}

func (jh *jobRoutes) deleteJob(c *fiber.Ctx) error {
	userid := c.Params("id")
	id, _ := strconv.Atoi(userid)

	if err := jh.j.DeleteJobByID(id); err != nil {
		return err
	}

	return c.JSON("job deleted successfully")
}
