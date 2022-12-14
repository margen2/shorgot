package middlewares

import (
	"log"
	"net/http"

	"github.com/margen2/shorgot/src/answers"
	"github.com/margen2/shorgot/src/auth"
)

// Logger adds a logger for each request
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// Authenticate authenticates the user based on the given JWT token
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			answers.Error(w, http.StatusUnauthorized, err)
			return
		}
		nextFunc(w, r)
	}
}
