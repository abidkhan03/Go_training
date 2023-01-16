
package handler

import (
	"encoding/json"
	"github.com/abidkhan03/go_training/csv"
	"net/http"
)

func CsvHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var filePath struct {
		Path string `json:"path"`
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}
	// decode the json request to filePath
	err := json.NewDecoder(r.Body).Decode(&filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// retrieve the data from the csv file
	data, err := csv.CsvtoJson(filePath.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// write the data to the response body as json data type and return the response
	_, err = w.Write([]byte(data))
	if err != nil {
		return
	}
}
