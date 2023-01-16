// Add a new route, which receives a local file path in json body.
// Handler method should read and parse the CSV.
// And return the data in file in json format.
// Import `./csv` in handlers.go and use the method as `csv.ParseCsv`

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/abidkhan03/go_training/handler/csv"
	"net/http"
	"os"
)

type request struct {
	Path string `json:"path"`
}

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func handlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return 
	}
	defer r.Body.Close()
	response := &Response{
		Code:    200,
		Message: "Welcome " + req.Path + "!",
	}
	er := json.NewEncoder(w).Encode(response)
	if er != nil {
		return 
	}
}

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return 
	}
	defer r.Body.Close()
	records, err := csv.ParseCsv(req.Path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	jsonData, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// print the json data
	fmt.Println(string(jsonData))

	// return json data
	_, err = w.Write(jsonData)
	if err != nil {
		return
	}
}

// Path: csv/parse.go
