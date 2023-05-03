package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fatih/color"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := color.New(color.FgRed).SprintFunc()
		path := color.New(color.FgGreen).SprintFunc()
		log.Printf("req %s on %s", method(r.Method), path(r.URL.Path))

		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v ...any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
