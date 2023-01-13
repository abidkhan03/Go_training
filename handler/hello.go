package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return 
	}

	response := &Response{
		Code:      http.StatusOK,
		Message:   fmt.Sprintf("Welcome %s!", req.Name),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return 
	}
	w.WriteHeader(http.StatusOK)
}
