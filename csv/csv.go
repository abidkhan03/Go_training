
package csv

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

// f/training-43 function to convert csv data to json data
func CsvtoJson(csvData string) (jsonData string, err error) {
	file, err := os.Open(csvData)
	if err != nil {
		return
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		return
	}
	jsondata, err := json.Marshal(rawCSVdata)
	if err != nil {
		return
	}
	return string(jsondata), nil
}
