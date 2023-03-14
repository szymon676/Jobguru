package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/szymon676/job-guru/users/domain/models"
)

const secret = "secret"

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

func (JwtService) CreateJWT(account *models.LoginUser) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
		"accountID": account.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (JwtService) ParseJWT(r *http.Request) (*jwt.Token, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return nil, fmt.Errorf("missing JWT cookie")
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid JWT: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid JWT: token is not valid")
	}

	return token, nil
}
