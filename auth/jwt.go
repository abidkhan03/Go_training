package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	abc "github.com/dgrijalva/jwt-go"
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

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user User
	// receive user oject in request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims := &abc.MapClaims{
		"username": user.Username,
		"exec":     time.Now().Add(time.Hour),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		
		if claims["username"] == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"error": "invalid credentials"}`)
			return
		}
		if claims["username"] == "Ali" {
			next.ServeHTTP(w, r)
			return
		}
	})
}
