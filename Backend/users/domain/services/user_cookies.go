package services

import (
	"net/http"
	"time"
)

func CreateCookie(w http.ResponseWriter, token string) error {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: expiration,
		Path:    "/",
	}

	http.SetCookie(w, &cookie)
	return nil
}
