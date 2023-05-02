package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/szymon676/jobguru/storage"
	"github.com/szymon676/jobguru/types"
)

const secret = "secret"

func CreateJWT(user *types.LoginUser, storage storage.UserStorager) (string, error) {
	account, _ := storage.GetUserByEmail(user.Email)

	claims := &jwt.MapClaims{
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
		"accountID": account.ID,
	}

	fmt.Println(account.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ParseJWT(r *http.Request) (*jwt.Token, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return nil, fmt.Errorf("%c", err)
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

func CreateCookie(w http.ResponseWriter, token string) error {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expiration,
		Path:     "/",
		MaxAge:   604800,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, &cookie)
	return nil
}
