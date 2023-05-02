package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/szymon676/jobguru/storage"
	"github.com/szymon676/jobguru/types"

	"github.com/szymon676/jobguru/api/validators"
)

type JobsHandler struct {
	storage storage.JobStorager
}

func NewJobHandler(storage storage.JobStorager) *JobsHandler {
	return &JobsHandler{
		storage: storage,
	}
}

func (jh JobsHandler) HandleCreateJob(w http.ResponseWriter, r *http.Request) error {
	var jobReq types.JobReq

	if err := json.NewDecoder(r.Body).Decode(&jobReq); err != nil {
		return err
	}

	err := validators.VerifyJobReq(jobReq)
	if err != nil {
		return err
	}

	err = jh.storage.CreateJob(uint(jobReq.UserID), jobReq.Title, jobReq.Company, jobReq.Skills, jobReq.Salary, jobReq.Description, jobReq.Currency, jobReq.Date, jobReq.Location)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, "job created successfully")
}

func (jh JobsHandler) HandleGetJobs(w http.ResponseWriter, r *http.Request) error {
	jobs, err := jh.storage.GetJobs()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, jobs)
}

func (jh JobsHandler) HandleGetJobsByUser(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	userid, _ := strconv.Atoi(path["userid"])

	jobs, err := jh.storage.GetJobsByUser(uint(userid))
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, jobs)
}

func (jh JobsHandler) HandleUpdateJob(w http.ResponseWriter, r *http.Request) error {
	var jobReq types.JobReq
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	if err := json.NewDecoder(r.Body).Decode(&jobReq); err != nil {
		return err
	}

	err := validators.VerifyJobReq(jobReq)
	if err != nil {
		return err
	}

	err = jh.storage.UpdateJob(id, jobReq.UserID, jobReq.Title, jobReq.Company, jobReq.Skills, jobReq.Salary, jobReq.Description, jobReq.Currency, jobReq.Date, jobReq.Location)
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, "job updated successfully")
}

func (jh JobsHandler) HandleDeleteJob(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id := path["id"]

	if err := jh.storage.DeleteJob(id); err != nil {
		return err
	}

	return WriteJSON(w, 204, "job deleted successfully")
}

func ParseDate(bj *types.JobReq) (time.Time, error) {
	dateStr := bj.Date
	dateLayout := "2006-01-02"

	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date")
	}

	return date, nil
}
