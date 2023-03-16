package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/szymon676/job-guru/jobs/storage"
	"github.com/szymon676/job-guru/jobs/types"
)

type JobsHandler struct {
	listenAddr string
	storage    storage.Storager
}

func (jh JobsHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/jobs", MakeHTTPHandleFunc(jh.handleCreateJob)).Methods("POST")
	router.HandleFunc("/jobs", MakeHTTPHandleFunc(jh.handleGetJobs)).Methods("GET")
	router.HandleFunc("/jobs/{id}", MakeHTTPHandleFunc(jh.handleGetJobsByUser)).Methods("GET")
	router.HandleFunc("/user", MakeHTTPHandleFunc(jh.handleGetUserInfo)).Methods("GET")
	router.HandleFunc("/jobs/{id}", MakeHTTPHandleFunc(jh.handleUpdateJob)).Methods("PUT")
	router.HandleFunc("/jobs/{id}", MakeHTTPHandleFunc(jh.handleDeleteJob)).Methods("DELETE")

	log.Println("server running on port:", jh.listenAddr)
	http.ListenAndServe(jh.listenAddr, router)
}

func NewApiServer(listenAddr string, storage storage.Storager) *JobsHandler {
	return &JobsHandler{
		listenAddr: listenAddr,
		storage:    storage,
	}
}

func (jh JobsHandler) handleCreateJob(w http.ResponseWriter, r *http.Request) error {
	var bindJob types.BindJob
	cookie, _ := r.Cookie("jwt")

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = jh.storage.CreateJob(uint(bindJob.UserID), bindJob.Title, bindJob.Company, bindJob.Skills, bindJob.Salary, bindJob.Description, bindJob.Currency, bindJob.Date, bindJob.Location)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, "job created successfully", cookie.Value)
}

func (jh JobsHandler) handleGetJobs(w http.ResponseWriter, r *http.Request) error {
	jobs, err := jh.storage.GetJobs()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, jobs)
}

func (jh JobsHandler) handleGetUserInfo(w http.ResponseWriter, r *http.Request) error {
	var smth interface{}

	resp, err := http.Get("http://localhost:5000/user")
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(&smth); err != nil {
		return err
	}
	return WriteJSON(w, 200, smth)
}

func (jh JobsHandler) handleGetJobsByUser(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	jobs, err := jh.storage.GetJobsByUser(uint(id))
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, jobs)
}

func (jh JobsHandler) handleUpdateJob(w http.ResponseWriter, r *http.Request) error {
	var bindJob types.BindJob
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = jh.storage.UpdateJob(id, bindJob.UserID, bindJob.Title, bindJob.Company, bindJob.Skills, bindJob.Salary, bindJob.Description, bindJob.Currency, bindJob.Date, bindJob.Location)
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, "job updated successfully")
}

func (jh JobsHandler) handleDeleteJob(w http.ResponseWriter, r *http.Request) error {
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
