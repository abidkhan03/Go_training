package handler

import (
	"net/http"

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
		// check expiration time
		// var expire_time = 300
		// if claims["exec"].(int64) < int64(expire_time) {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// 	return
		// }

		if claims["username"] == "Ali" {
			next.ServeHTTP(w, r)
		}
	})
}
