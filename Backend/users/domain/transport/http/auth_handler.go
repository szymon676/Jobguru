package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/szymon676/job-guru/users/domain/models"
	"github.com/szymon676/job-guru/users/domain/repository"
	"github.com/szymon676/job-guru/users/domain/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	listenaddr string
	jwtservice services.JwtService
}

func NewApiServer(listenaddr string) *AuthHandler {
	jwtservice := services.NewJwtService()
	return &AuthHandler{
		listenaddr: listenaddr,
		jwtservice: *jwtservice,
	}
}

func (ah AuthHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/register", MakeHTTPHandleFunc(ah.handleRegisterUser)).Methods("POST")
	router.HandleFunc("/login", MakeHTTPHandleFunc(ah.handleLoginUser)).Methods("POST")
	router.HandleFunc("/users/{id}", MakeHTTPHandleFunc(ah.handleGetUserByID)).Methods("GET")
	router.HandleFunc("/user", MakeHTTPHandleFunc(ah.handleGetUserInfo)).Methods("GET")

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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(bindRegisterUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	if err := repository.CreateUser(bindRegisterUser.Name, string(hashedPassword), bindRegisterUser.Email); err != nil {
		return err
	}

	return WriteJSON(w, 202, "User registration done successfully")
}

func (ah AuthHandler) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
	var loginUser models.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		return err
	}

	if err := services.VerifyLogin(loginUser); err != nil {
		w.Header().Set("WWW-Authenticate", `Bearer realm="Restricted"`)
		return WriteJSON(w, http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	token, err := ah.jwtservice.CreateJWT(&loginUser)
	if err != nil {
		return err
	}

	if err := services.CreateCookie(w, token); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, "user logged in successfully")
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

func (ah AuthHandler) handleGetUserInfo(w http.ResponseWriter, r *http.Request) error {
	token, err := ah.jwtservice.ParseJWT(r)
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid JWT claims")
	}

	userID, ok := claims["accountID"].(string)
	if !ok {
		return fmt.Errorf("invalid JWT claims")
	}

	id, _ := strconv.Atoi(userID)

	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, 200, user)
}
