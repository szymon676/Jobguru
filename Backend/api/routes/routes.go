package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/szymon676/jobguru/api/handlers"
)

func SetupRoutes(listenaddr string, ah *handlers.AuthHandler, jh *handlers.JobsHandler) error {
	router := mux.NewRouter()

	u := router.PathPrefix("/users").Subrouter()
	u.HandleFunc("/register", makeHTTPHandleFunc(ah.HandleRegisterUser)).Methods("POST")
	u.HandleFunc("/login", makeHTTPHandleFunc(ah.HandleLoginUser)).Methods("POST")
	u.HandleFunc("/user", makeHTTPHandleFunc(ah.HandleGetUserInfo)).Methods("GET")
	u.HandleFunc("/users/{id}", makeHTTPHandleFunc(ah.HandleGetUserByID)).Methods("GET")

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
