// write parseCsv function

package csv

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
)
// f/training-55 function to parse csv file and return json data
func ParseCsv(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// f/training-43 function to convert csv data to json data
func CsvtoJson(csvData string) (jsonData []byte, err error) {

	file, err := os.Open(csvData)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	jsondata, err := json.Marshal(rawCSVdata)
	if err != nil {
		return nil, err
	}
	return jsondata, nil
}

// f/training-43 function to read json data from CsvtoJson function and write it to a file
func CsvHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var filePath struct {
		Path string `json:"path"`
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	// decode the json request to filePath
	err := json.NewDecoder(r.Body).Decode(&filePath)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// retrieve the data from the csv file
	data, err := CsvtoJson(filePath.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// write the data to the response body as json data type and return the response
	_, err = w.Write(data)
	if err != nil {
		return
	}
}