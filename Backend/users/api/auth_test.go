package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/szymon676/job-guru/users/storage"
	"github.com/szymon676/job-guru/users/types"
)

func TestLogin(t *testing.T) {
	dsn := "host=localhost user=postgres password=1234 dbname=jobguru-users port=5432 sslmode=disable"

	_, err := storage.NewPostgresDatabase("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	storage := storage.NewPostgresStorage()

	as := NewApiServer("", storage)

	handler := http.HandlerFunc(MakeHTTPHandleFunc(as.handleLoginUser))

	loginBody := types.LoginUser{
		ID:       "1",
		Password: "12345",
	}

	loginJSON, err := json.Marshal(loginBody)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(loginJSON))
	handler.ServeHTTP(recorder, req)
	resp := recorder.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("expected status %v got %v", http.StatusAccepted, resp.StatusCode)
	}
}

func TestRegister(t *testing.T) {
	dsn := "host=localhost user=postgres password=1234 dbname=jobguru-users port=5432 sslmode=disable"

	_, err := storage.NewPostgresDatabase("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	storage := storage.NewPostgresStorage()

	as := NewApiServer("", storage)

	handler := http.HandlerFunc(MakeHTTPHandleFunc(as.handleRegisterUser))

	loginBody := types.RegisterUser{
		Name:     "Bob",
		Email:    "Bob_is_awesome@gmail.com",
		Password: "12345",
	}

	loginJSON, err := json.Marshal(loginBody)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/register", bytes.NewReader(loginJSON))
	handler.ServeHTTP(recorder, req)
	resp := recorder.Result()

	if resp.StatusCode != http.StatusAccepted {
		t.Errorf("expected status %v got %v", http.StatusAccepted, resp.StatusCode)
	}
}
