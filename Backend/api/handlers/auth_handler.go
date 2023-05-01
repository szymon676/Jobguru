package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/szymon676/jobguru/storage"
	"github.com/szymon676/jobguru/types"

	"github.com/szymon676/jobguru/api/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	jwtservice utils.JwtService
	storage    storage.UserStorager
	verifier   utils.VerifyService
}

func NewAuthHandler(storage storage.UserStorager) *AuthHandler {
	jwtservice := utils.NewJwtService(storage)
	verifier := utils.NewVerifier(storage)

	return &AuthHandler{
		jwtservice: *jwtservice,
		storage:    storage,
		verifier:   verifier,
	}
}

func (ah *AuthHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) error {
	var bindRegisterUser types.RegisterUser

	if err := json.NewDecoder(r.Body).Decode(&bindRegisterUser); err != nil {
		return err
	}

	if err := ah.verifier.VerifyRegister(bindRegisterUser); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bindRegisterUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := ah.storage.CreateUser(bindRegisterUser.Name, string(hashedPassword), bindRegisterUser.Email); err != nil {
		return err
	}

	return WriteJSON(w, 202, "User registration done successfully")
}

func (ah *AuthHandler) HandleLoginUser(w http.ResponseWriter, r *http.Request) error {
	var loginUser types.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		return err
	}

	if err := ah.verifier.VerifyLogin(loginUser); err != nil {
		w.Header().Set("WWW-Authenticate", `Bearer realm="Restricted"`)
		return WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	token, err := ah.jwtservice.CreateJWT(&loginUser)
	if err != nil {
		return err
	}

	if err := utils.CreateCookie(w, token); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, "user logged in successfully")
}

func (ah *AuthHandler) HandleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	user, err := ah.storage.GetUserByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, user)
}
func (ah *AuthHandler) HandleGetUserInfo(w http.ResponseWriter, r *http.Request) error {
	token, err := ah.jwtservice.ParseJWT(r)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid JWT claims %v", claims)
	}

	userID, ok := claims["accountID"].(float64)
	if !ok {
		return fmt.Errorf("missing or invalid accountID claim")
	}

	id := int(userID)
	user, err := ah.storage.GetUserByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, user)
}
