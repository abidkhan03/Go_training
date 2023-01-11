package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	response := &Response{
		Code:      200,
		Message:   "Welcome " + req.Name + "!",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := chi.NewRouter()
	r.Post("/hello", HelloHandler)
	http.ListenAndServe(":8000", r)
}
