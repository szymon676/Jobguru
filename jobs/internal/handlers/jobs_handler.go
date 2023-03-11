package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/szymon676/job-guru/jobs/internal/database"
	"github.com/szymon676/job-guru/jobs/internal/models"
	"github.com/szymon676/job-guru/jobs/internal/utils"
)

func (jh JobsHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/jobs", utils.MakeHTTPHandleFunc(jh.handleGetUser)).Methods("GET")
	router.HandleFunc("/jobs", utils.MakeHTTPHandleFunc(jh.handleCreateUser)).Methods("POST")

	log.Println("server running on port:", jh.listenAddr)
	http.ListenAndServe(jh.listenAddr, router)
}

type JobsHandler struct {
	db         *sql.DB
	listenAddr string
}

func NewJobsHandler(db *sql.DB, listenAddr string) *JobsHandler {
	return &JobsHandler{
		db:         db,
		listenAddr: listenAddr,
	}
}

func (jh JobsHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	var bindJob models.BindJob

	if err := json.NewDecoder(r.Body).Decode(&bindJob); err != nil {
		return err
	}

	err := utils.VerifyJSON(bindJob)
	if err != nil {
		return err
	}

	err = database.CreateJob(bindJob.Title, bindJob.Category, bindJob.Description, bindJob.Skills)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusAccepted, "ok")
}

func (jh JobsHandler) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	jobs, err := database.GetJobs()
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, jobs)
}

func (jh JobsHandler) handleUpdateUser() error {
	return nil
}

func (jh JobsHandler) handleDeleteUser() error {
	return nil
}
