package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/szymon676/jobguru/api/handlers"
)

func SetupRoutes(listenaddr string, uh *handlers.UsersHandler, jh *handlers.JobsHandler) error {
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	u := router.PathPrefix("/auth").Subrouter()
	u.HandleFunc("/register", makeHTTPHandleFunc(uh.HandleRegisterUser)).Methods("POST")
	u.HandleFunc("/login", makeHTTPHandleFunc(uh.HandleLoginUser)).Methods("POST")
	u.HandleFunc("/users/{id}", makeHTTPHandleFunc(uh.HandleGetUserByID)).Methods("GET")

	j := router.PathPrefix("/jobs").Subrouter()
	j.HandleFunc("/jobs", makeHTTPHandleFunc(jh.HandleCreateJob)).Methods("POST")
	j.HandleFunc("/jobs", makeHTTPHandleFunc(jh.HandleGetJobs)).Methods("GET")
	j.HandleFunc("/jobs/{userid}", makeHTTPHandleFunc(jh.HandleGetJobsByUser)).Methods("GET")
	j.HandleFunc("/jobs/{id}", makeHTTPHandleFunc(jh.HandleUpdateJob)).Methods("PUT")
	j.HandleFunc("/jobs/{id}", makeHTTPHandleFunc(jh.HandleDeleteJob)).Methods("DELETE")

	log.Println("server listening on port", listenaddr)
	err := http.ListenAndServe(listenaddr, router)
	if err != nil {
		return err
	}
	return nil
}
