
package handler

import (
	"encoding/json"
	"github.com/abidkhan03/go_training/csv"
	"log"
	"net/http"
)

// CsvRequest Create a new struct to store the path of the csv file
var CsvRequest struct{
	Path string `json:"path"`
}

// Csv function takes the path of the csv file and converts it into json
// It decodes the request body and encodes the response body
// It writes the response body as json data type
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
	// return an error if the csv file cannot be read
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write the response body
	w.WriteHeader(http.StatusOK)

	// write the data to the response body
	_, err = w.Write([]byte(data))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
