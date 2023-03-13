package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/szymon676/job-guru/jobs/internal/database"
	"github.com/szymon676/job-guru/jobs/internal/models"
	"github.com/szymon676/job-guru/jobs/internal/utils"
	"github.com/szymon676/job-guru/jobs/internal/validation"
)

type JobsHandler struct {
	listenAddr string
}

func (jh JobsHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/jobs", utils.MakeHTTPHandleFunc(jh.handleGetUser)).Methods("GET")
	router.HandleFunc("/jobs", utils.MakeHTTPHandleFunc(jh.handleCreateUser)).Methods("POST")
	router.HandleFunc("/jobs/{id}", utils.MakeHTTPHandleFunc(jh.handleUpdateUser)).Methods("PUT")
	router.HandleFunc("/jobs/{id}", utils.MakeHTTPHandleFunc(jh.handleDeleteUser)).Methods("DELETE")

	log.Println("server running on port:", jh.listenAddr)
	http.ListenAndServe(jh.listenAddr, router)
}

func NewApiServer(listenAddr string) *JobsHandler {
	return &JobsHandler{
		listenAddr: listenAddr,
	}
}

func (jh JobsHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var bindJob models.BindJob

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := validation.VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = database.CreateJob(bindJob.Title, bindJob.Category, bindJob.Description, bindJob.Skills)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusAccepted, "job created successfully")
}

func (jh JobsHandler) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	jobs, err := database.GetJobs()
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, jobs)
}

func (jh JobsHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {
	var bindJob models.BindJob
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := validation.VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = database.UpdateJob(id, bindJob.Title, bindJob.Skills, bindJob.Category, bindJob.Description)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, 200, "job updated successfully")
}

func (jh JobsHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id := path["id"]

	if err := database.DeleteJob(id); err != nil {
		return err
	}

	return utils.WriteJSON(w, 204, "job deleted successfully")
}
