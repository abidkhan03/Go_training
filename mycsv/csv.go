package csv

import (
	"encoding/csv"
	"os"
)

func CsvtoJson(csvData string) (jsonData [][]string, err error) {
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

	return rawCSVdata, nil
}
