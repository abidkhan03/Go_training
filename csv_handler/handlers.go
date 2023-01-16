// Add a new route, which receives a local file path in json body.
// Handler method should read and parse the CSV.
// And return the data in file in json format.
// Import `./csv` in handlers.go and use the method as `csv.ParseCsv`

package csvhandler

import (
	"github.com/abidkhan03/go_training/csv_handler/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/go-chi/chi/v5"
)

type Request struct {
	Path string `json:"path"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const PORT = "8000"

func main() {
	r := chi.NewRouter()
	r.Post("/hello", handlers)
	r.Post("/parse", ParseHandler)
	fmt.Println("Server is running on port ", PORT)
	http.ListenAndServe(PORT, r)
}

func handlers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	response := &Response{
		Code:    200,
		Message: "Welcome " + req.Path + "!",
	}
	json.NewEncoder(w).Encode(response)
}

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
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
}

// Path: csv/parse.go
