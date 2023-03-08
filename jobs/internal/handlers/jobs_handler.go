package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/szymon676/job-guru/jobs/internal/util"
)

func (jh JobsHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/jobs", util.MakeHTTPHandleFunc(jh.handleGetUser))
	log.Println("server running on port: ", jh.listenAddr)

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

func (jh JobsHandler) handleGetUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (jh JobsHandler) handleCreateUser() error {
	return nil
}

func (jh JobsHandler) handleUpdateUser() error {
	return nil
}

func (jh JobsHandler) handleDeleteUser() error {
	return nil
}
