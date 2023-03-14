package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/szymon676/job-guru/users/domain/models"
	"github.com/szymon676/job-guru/users/domain/repository"
	"github.com/szymon676/job-guru/users/domain/services"
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

	router.HandleFunc("/register", MakeHTTPHandleFunc(ah.handleRegisterUser)).Methods("POST")
	router.HandleFunc("/login", MakeHTTPHandleFunc(ah.handleLoginUser)).Methods("POST")
	router.HandleFunc("/users/{id}", MakeHTTPHandleFunc(ah.handleGetUserByID)).Methods("GET")

	fmt.Println("server listening on port:", ah.listenaddr)
	http.ListenAndServe(ah.listenaddr, router)
}

func (AuthHandler) handleRegisterUser(w http.ResponseWriter, r *http.Request) error {
	var bindRegisterUser models.RegisterUser

	if err := json.NewDecoder(r.Body).Decode(&bindRegisterUser); err != nil {
		return err
	}

	if err := services.VerifyRegister(bindRegisterUser); err != nil {
		return err
	}

	if err := repository.CreateUser(bindRegisterUser.Name, bindRegisterUser.Password, bindRegisterUser.Email); err != nil {
		return err
	}

	return WriteJSON(w, 202, "User registration done successfully")
}

func (AuthHandler) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
	var loginUser models.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		return err
	}

	if err := services.ValidateUser(loginUser); err != nil {
		return err
	}

	token, err := services.CreateJWT(&loginUser)
	if err != nil {
		return err
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)

	return WriteJSON(w, 200, "user logged in")
}

func (AuthHandler) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, user)
}
