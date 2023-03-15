package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	router.HandleFunc("/jobs/{id}", MakeHTTPHandleFunc(jh.handleUpdateJob)).Methods("PUT")
	router.HandleFunc("/jobs/{id}", MakeHTTPHandleFunc(jh.handleDeleteJob)).Methods("DELETE")

	log.Println("server running on port:", jh.listenAddr)
	http.ListenAndServe(jh.listenAddr, router)
}

func NewApiServer(listenAddr string) *JobsHandler {
	postgrestorage := storage.NewPostgreStorage()
	return &JobsHandler{
		listenAddr: listenAddr,
		storage:    postgrestorage,
	}
}

func (jh JobsHandler) handleCreateJob(w http.ResponseWriter, r *http.Request) error {
	var bindJob types.BindJob

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = jh.storage.CreateJob(bindJob.Title, bindJob.Company, bindJob.Skills, bindJob.Salary, bindJob.Description, bindJob.Currency, bindJob.Date, bindJob.Location)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, "job created successfully")
}

func (jh JobsHandler) handleGetJobs(w http.ResponseWriter, r *http.Request) error {
	jobs, err := jh.storage.GetJobs()
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

	err = jh.storage.UpdateJob(id, bindJob.Title, bindJob.Company, bindJob.Skills, bindJob.Salary, bindJob.Description, bindJob.Currency, bindJob.Date, bindJob.Location)
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
