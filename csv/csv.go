
package csv

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)
// CsvtoJson function takes the path of the csv file and returns the data in json format
// It takes the path of the csv file as a string and returns the data in json format as a string
// It returns an error if the file is not found or the data is not in the correct format
// It returns an error if the data cannot be converted to json format
// It returns an error if the data cannot be written to the response body
func CsvtoJson(csvData string) (jsonData []byte, err error) {
	// open the csv file and return an error if the file is not found
	file, err := os.Open(csvData)
	if err != nil {
		log.Println(err)
		http.Error(nil, err.Error(), http.StatusBadRequest)
		return
	}
	// defer the closing of the file until the function returns
	defer file.Close()

	reader := csv.NewReader(file)

	// read the csv file and return an error if the data is not in the correct format
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		log.Println(err)
		http.Error(nil, err.Error(), http.StatusInternalServerError)
		return
	}

	// create new variable to store the data in the json format
	var data []map[string]string
	// store the header of the csv file in a variable
	header := rawCSVdata[0]

	// loop through the data in the csv file onwards the header
	for _, row := range rawCSVdata[1:] {
		// create a new map to store the data in the json format
		mp := make(map[string]string)
		// loop through the rows of the csv file and store the data in the map
		for j := range row {
			mp[header[j]] = row[j]
		}
		// append the data to the data variable
		data = append(data, mp)
	}

	// convert the data to json Marshal encoding and return an error
	//if the data cannot be converted to json format
	jsonData, err = json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(nil, err.Error(), http.StatusNotAcceptable)
		return
	}
	// return the data in json format or an error if the data cannot be written to the response body
	return jsonData, nil

}
