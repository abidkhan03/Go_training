package csv

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// CsvtoJson function takes the path of the csv file and returns the data in json format
// It takes the path of the csv file as a string and returns the data in json format as a string
// It returns an error if the file is not found or the data is not in the correct format
// It returns an error if the data cannot be converted to json format
// It returns an error if the data cannot be written to the response body

func CsvtoJson(csvData string, hasHeaders bool) (jsonData []byte, err error) {

	// open the csv file and return an error if the file is not found
	file, err := os.Open(csvData)
	if err != nil {
		return nil, err
	}
	// defer the closing of the file until the function returns
	defer file.Close()

	reader := csv.NewReader(file)

	// read the csv file and return an error if the data is not in the correct format
	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// if file is empty return an error
	if len(rawCSVdata) == 0 {
		return nil, errors.New("File is empty")
	}

	// create new variable to store the data in the json format
	var data []map[string]string
	// store the header of the csv file in a variable
	
	columnName := "column"

	// if the csv file has headers, store the data in a map with the header as the key
	if hasHeaders  {
		// header to store the headers of the csv file
		header := rawCSVdata[0]
		// loop through the data and store the data in a map with the header as the key
		for _, row := range rawCSVdata[1:] {
			// create a new map for each row of data
			mp := make(map[string]string)
			for j := range row {
				// store the data in the map with the header as the key and the data as the value
				mp[header[j]] = row[j]

			}
			data = append(data, mp)
		}
		// if the csv file does not have headers, store the data in a map with the column name as the key
	} else {
		header := make([]string, len(rawCSVdata[0]))
		for i := range header {
			header[i] = fmt.Sprintf("%s %d", columnName, i+1)
		}
		for _, row := range rawCSVdata {
			mp := make(map[string]string)
			for j := range row {
				mp[header[j]] = row[j]

			}
			data = append(data, mp)
		}
	}

	// convert the data to json Marshal encoding and return an error
	//if the data cannot be converted to json format
	jsonData, err = json.Marshal(data)
	if err != nil {
		return nil, err
	}
	// return the data in json format or an error if the data cannot be written to the response body
	return jsonData, nil

}
