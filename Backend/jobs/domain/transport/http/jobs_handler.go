package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/szymon676/job-guru/jobs/domain/models"
	"github.com/szymon676/job-guru/jobs/domain/repository"
	"github.com/szymon676/job-guru/jobs/domain/services"
)

type JobsHandler struct {
	listenAddr string
}

func (jh JobsHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/jobs", MakeHTTPHandleFunc(jh.handleGetJobs)).Methods("GET")
	router.HandleFunc("/jobs", MakeHTTPHandleFunc(jh.handleCreateJob)).Methods("POST")
	router.HandleFunc("/jobs/{id}", MakeHTTPHandleFunc(jh.handleUpdateJob)).Methods("PUT")
	router.HandleFunc("/jobs/{id}", MakeHTTPHandleFunc(jh.handleDeleteJob)).Methods("DELETE")

	log.Println("server running on port:", jh.listenAddr)
	http.ListenAndServe(jh.listenAddr, router)
}

func NewApiServer(listenAddr string) *JobsHandler {
	return &JobsHandler{
		listenAddr: listenAddr,
	}
}

func (jh JobsHandler) handleCreateJob(w http.ResponseWriter, r *http.Request) error {
	var bindJob models.BindJob

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := services.VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = repository.CreateJob(bindJob.Title, bindJob.Category, bindJob.Description, bindJob.Skills)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, "job created successfully")
}

func (jh JobsHandler) handleGetJobs(w http.ResponseWriter, r *http.Request) error {
	jobs, err := repository.GetJobs()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, jobs)
}

func (jh JobsHandler) handleUpdateJob(w http.ResponseWriter, r *http.Request) error {
	var bindJob models.BindJob
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := services.VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = repository.UpdateJob(id, bindJob.Title, bindJob.Skills, bindJob.Category, bindJob.Description)
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, "job updated successfully")
}

func (jh JobsHandler) handleDeleteJob(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id := path["id"]

	if err := repository.DeleteJob(id); err != nil {
		return err
	}

	return WriteJSON(w, 204, "job deleted successfully")
}
