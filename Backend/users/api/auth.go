package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/szymon676/job-guru/users/storage"
	"github.com/szymon676/job-guru/users/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	listenaddr string
	jwtservice JwtService
	storage    storage.Storager
	verifier   VerifyService
}

func NewApiServer(listenaddr string, storage storage.Storager) *AuthHandler {
	jwtservice := NewJwtService()
	verifier := NewVerifier(storage)

	return &AuthHandler{
		listenaddr: listenaddr,
		jwtservice: *jwtservice,
		storage:    storage,
		verifier:   verifier,
	}
}

func (ah AuthHandler) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/register", MakeHTTPHandleFunc(ah.handleRegisterUser)).Methods("POST")
	router.HandleFunc("/login", MakeHTTPHandleFunc(ah.handleLoginUser)).Methods("POST")
	router.HandleFunc("/user", MakeHTTPHandleFunc(ah.handleGetUserInfo)).Methods("GET")
	router.HandleFunc("/users/{id}", MakeHTTPHandleFunc(ah.handleGetUserByID)).Methods("GET")

	fmt.Println("server listening on port:", ah.listenaddr)
	http.ListenAndServe(ah.listenaddr, router)
}

func (ah AuthHandler) handleRegisterUser(w http.ResponseWriter, r *http.Request) error {
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

func (ah AuthHandler) handleLoginUser(w http.ResponseWriter, r *http.Request) error {
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

	if err := CreateCookie(w, token); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusAccepted, "user logged in successfully")
}

func (ah AuthHandler) handleGetUserByID(w http.ResponseWriter, r *http.Request) error {
	path := mux.Vars(r)
	id, _ := strconv.Atoi(path["id"])

	user, err := ah.storage.GetUserByID(id)
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

func WriteJSON(w http.ResponseWriter, status int, v ...any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
