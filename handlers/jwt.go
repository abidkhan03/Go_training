package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if claims["username"] == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if claims["exp"].(time.Time).Before(time.Now()) {
			http.Error(w, "Unauthorized or token Expired", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

		if claims["username"] == "Ali" {
			next.ServeHTTP(w, r)
		}
	})
}
