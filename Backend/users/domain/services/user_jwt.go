package services

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/szymon676/job-guru/users/domain/models"
)

const secret = "secret"

func CreateJWT(account *models.LoginUser) (string, error) {
	claims := &jwt.MapClaims{
		"expiresAt":     time.Now().Add(time.Hour * 24).Unix(),
		"accountNumber": account.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
