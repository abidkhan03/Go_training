
package handler

import (
	"encoding/json"
	"github.com/abidkhan03/go_training/csv"
	"net/http"
)

var CsvRequest struct{
	Path string `json:"path"`
}
func Csv(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// decode the json request to filePath
	err := json.NewDecoder(r.Body).Decode(&CsvRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// retrieve the data from the csv file
	data, err := csv.CsvtoJson(CsvRequest.Path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// write the data to the response body as json data type and return the response
	_, err = w.Write([]byte(data))
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
