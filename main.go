package main

import (
	"encoding/json"
	"github.com/abidkhan03/go_training/handler"
	"github.com/abidkhan03/go_training/mycsv"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)
func csvHandler(w http.ResponseWriter, r *http.Request) {
	var filePath struct {
		Path string `json:"path"`
	}
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&filePath)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	data, err := csv.CsvtoJson(filePath.Path)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		return
	}
}
const PORT = ":8000"
func main() {
	r := chi.NewRouter()
	r.Post("/hello", handler.HelloHandler)
	r.Post("/parse", csvHandler)
	log.Println("Server is running on port " + PORT)
	log.Println(http.ListenAndServe(PORT, r))
}
