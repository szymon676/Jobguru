package middlewares

import (
	"log"
	"net/http"
)

func Log(http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("req %s on %s", r.Method, r.URL.Path)
	}
}
