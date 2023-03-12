package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/szymon676/job-guru/users/internal/utils"
)

type AuthHandler struct {
	listenaddr string
}

func NewApiServer(listenaddr string) *AuthHandler {
	return &AuthHandler{
		listenaddr: listenaddr,
	}
}

func (ah AuthHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/register", utils.MakeHTTPHandleFunc(ah.handleRegisterUser)).Methods("POST")

	fmt.Println("server listening on port:", ah.listenaddr)
	http.ListenAndServe(ah.listenaddr, router)
}

func (AuthHandler) handleRegisterUser(w http.ResponseWriter, r *http.Request) error {
	return utils.WriteJSON(w, 202, "User registration done successfully")
}

func (AuthHandler) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
	return nil
}
