package handler

import (
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	// Get the user's credentials from the request
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check the user's credentials
	if username != "admin" || password != "admin" {
		http.Error(w, "Unauthorized", 401)
		return
	}


}

