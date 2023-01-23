package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Create a struct for the request body that stores the name of the user and path of the csv file
type HelloRequest struct {
	Name string `json:"name"`
	Path string
}

// Create a struct for the response body
type HelloResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// HelloHandler function takes the name of the user and returns a response
// It decodes the request body and encodes the response body
// It returns the response body as json data type
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// set the content type of the response body to json
	w.Header().Set("Content-Type", "application/json")

	var req HelloRequest
	// decode the json request to req variable
	err := json.NewDecoder(r.Body).Decode(&req)
	// return an error if the request body cannot be decoded
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	// create a new response body and store the data in it
	response := &HelloResponse{
		Code:      http.StatusOK,
		Message:   fmt.Sprintf("Welcome %s!", req.Name),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	// encode the response body and return the response
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	// return the response
	w.WriteHeader(http.StatusOK)
}
