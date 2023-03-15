package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/szymon676/job-guru/jobs/storage"
	"github.com/szymon676/job-guru/jobs/types"
)

func init() {
	storage.NewDatabase("postgres", "host=localhost user=postgres password=1234 dbname=jobguru-jobs port=5432 sslmode=disable")
}

func TestHandleCreateJob(t *testing.T) {
	jh := NewApiServer("")
	srv := httptest.NewServer(http.HandlerFunc(MakeHTTPHandleFunc(jh.handleCreateJob)))
	defer srv.Close()

	job := types.BindJob{
		Title:       "Software Engineer",
		Company:     "Acme Inc",
		Skills:      []string{"Go", "Java", "Python"},
		Salary:      100000,
		Description: "We're looking for a software engineer to join our team",
		Currency:    "USD",
		Date:        "2022-01-01",
		Location:    "San Francisco",
	}

	jobJSON, err := json.Marshal(job)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", srv.URL+"/jobs", bytes.NewReader(jobJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("expected status %d; got %d", http.StatusAccepted, resp.StatusCode)
	}
}

func TestHandleGetJobs(t *testing.T) {
	jh := NewApiServer("")
	srv := httptest.NewServer(http.HandlerFunc(MakeHTTPHandleFunc(jh.handleGetJobs)))
	defer srv.Close()

	req, err := http.NewRequest("GET", srv.URL+"/jobs", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d; got %d", http.StatusOK, resp.StatusCode)
	}
}
