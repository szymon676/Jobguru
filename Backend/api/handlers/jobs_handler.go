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

	"github.com/szymon676/jobguru/api/utils"
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
	var bindJob types.BindJob

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := utils.VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = jh.storage.CreateJob(uint(bindJob.UserID), bindJob.Title, bindJob.Company, bindJob.Skills, bindJob.Salary, bindJob.Description, bindJob.Currency, bindJob.Date, bindJob.Location)
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
	var bindJob types.BindJob
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := utils.VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = jh.storage.UpdateJob(id, bindJob.UserID, bindJob.Title, bindJob.Company, bindJob.Skills, bindJob.Salary, bindJob.Description, bindJob.Currency, bindJob.Date, bindJob.Location)
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

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func MakeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v ...any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func ParseDate(bj *types.BindJob) (time.Time, error) {
	dateStr := bj.Date
	dateLayout := "2006-01-02"

	date, err := time.Parse(dateLayout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date")
	}

	return date, nil
}
