package main
import (
	"encoding/json"
	"encoding/csv"
	"fmt"
	"os"
)

// Write a program in `golang` which parses the file path passed in command-line argument as `csv`
// and prints its data in `json` format.

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}
	filePath:=os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
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
}