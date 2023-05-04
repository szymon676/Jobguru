package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/szymon676/jobguru/types"

	"github.com/szymon676/jobguru/service"
)

type JobsHandler struct {
	service service.IJobService
}

func NewJobHandler(service service.IJobService) *JobsHandler {
	return &JobsHandler{
		service: service,
	}
}

func (jh JobsHandler) HandleCreateJob(w http.ResponseWriter, r *http.Request) error {
	var jobReq types.JobReq

	if err := json.NewDecoder(r.Body).Decode(&jobReq); err != nil {
		return err
	}

	if err := jh.service.CreateJob(jobReq); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, "job created successfully")
}

func (jh JobsHandler) HandleGetJobs(w http.ResponseWriter, r *http.Request) error {
	jobs, err := jh.service.GetJobs()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, jobs)
}

func (jh JobsHandler) HandleGetJobsByUser(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	userid, _ := strconv.Atoi(path["userid"])

	jobs, err := jh.service.GetJobsByUser(userid)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, jobs)
}

func (jh JobsHandler) HandleUpdateJob(w http.ResponseWriter, r *http.Request) error {
	var jobReq types.JobReq
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	err := json.NewDecoder(r.Body).Decode(&jobReq)
	if err != nil {
		return err
	}

	err = jh.service.UpdateJobByID(id, jobReq)
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, "job updated successfully")
}

func (jh JobsHandler) HandleDeleteJob(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	if err := jh.service.DeleteJobByID(id); err != nil {
		return err
	}

	return WriteJSON(w, 204, "job deleted successfully")
}
