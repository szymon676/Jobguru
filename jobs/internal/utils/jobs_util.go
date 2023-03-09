package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/szymon676/job-guru/jobs/internal/models"
)

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

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func VerifyJSON(bj models.BindJob) error {
	if bj.Category == "" || bj.Title == "" || bj.Description == "" || bj.Skills == nil {
		return fmt.Errorf("error binding job")
	}

	if len(bj.Skills) < 1 || len(bj.Category) < 3 || len(bj.Title) < 3 || len(bj.Description) < 10 {
		return fmt.Errorf("error binding job")
	}
	return nil
}
