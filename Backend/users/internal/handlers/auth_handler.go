package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/szymon676/job-guru/users/internal/database"
	"github.com/szymon676/job-guru/users/internal/models"
	"github.com/szymon676/job-guru/users/internal/utils"
	"github.com/szymon676/job-guru/users/internal/validation"
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
	router.HandleFunc("/login", utils.MakeHTTPHandleFunc(ah.handleLoginUser)).Methods("POST")
	router.HandleFunc("/users/{id}", utils.MakeHTTPHandleFunc(ah.handleGetUserByID)).Methods("GET")

	fmt.Println("server listening on port:", ah.listenaddr)
	http.ListenAndServe(ah.listenaddr, router)
}

func (AuthHandler) handleRegisterUser(w http.ResponseWriter, r *http.Request) error {
	var bindRegisterUser models.RegisterUser

	if err := json.NewDecoder(r.Body).Decode(&bindRegisterUser); err != nil {
		return err
	}

	if err := validation.VerifyRegister(bindRegisterUser); err != nil {
		return err
	}

	if err := database.CreateUser(bindRegisterUser.Name, bindRegisterUser.Password, bindRegisterUser.Email); err != nil {
		return err
	}

	return utils.WriteJSON(w, 202, "User registration done successfully")
}

func (AuthHandler) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
	var loginUser models.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		return err
	}

	if err := validation.ValidateUser(loginUser); err != nil {
		return err
	}

	return nil
}

func (AuthHandler) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	user, err := database.GetUserByID(id)
	if err != nil {
		return err
	}

	return utils.WriteJSON(w, 200, user)
}
