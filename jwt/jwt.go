package jwt

import (
	"encoding/json"
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
	TokenAuth = &jwtauth.JWTAuth{}
}

func signUp(w http.ResponseWriter, r *http.Request) {
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


// func Authenticator(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token, claims, err := jwtauth.FromContext(r.Context())
// 		if err != nil {
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		// set the token in the contex

// 		t.PrivateClaims()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusUnauthorized)
// 			return
// 		}

// 		userRole := claims["role"].(string)
// 		if userRole != "admin" {
// 			http.Error(w, "Forbidden", 403)
// 			return
// 		}
// 		next.ServeHTTP(w, r) // call next handler
// 	})

// }
