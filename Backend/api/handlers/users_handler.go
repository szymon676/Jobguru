package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/szymon676/jobguru/service"
	"github.com/szymon676/jobguru/types"

	"github.com/szymon676/jobguru/api/auth"
)

type UsersHandler struct {
	service service.IUserService
}

func NewUsersHandler(service service.IUserService) *UsersHandler {
	return &UsersHandler{
		service: service,
	}
}

func (uh *UsersHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) error {
	var req types.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	if err := uh.service.CreateUser(req); err != nil {
		return err
	}
	return WriteJSON(w, 202, "User registration done successfully")
}

func (uh *UsersHandler) HandleLoginUser(w http.ResponseWriter, r *http.Request) error {
	var req types.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	token, err := uh.service.LoginUser(req)
	if err != nil {
		return err
	}

	if err := auth.CreateCookie(w, token); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, "user logged in successfully")
}

func (uh *UsersHandler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	user, err := uh.service.GetUserByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, user)
}
