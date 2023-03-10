package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var byte_key = []byte("secret")

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user User
	// receive user oject in request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	claims := &jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(byte_key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("JWT-Token", tokenString)
	w.WriteHeader(http.StatusOK)
}
