package main

import (
	//"encoding/json"
	"github.com/abidkhan03/go_training/handler"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

const PORT = ":8000"
func main() {
	r := chi.NewRouter()
	r.Post("/hello", handler.HelloHandler)
	r.Post("/handler", handler.CsvHandler)
	log.Println("Server is running on port " + PORT)
	log.Println(http.ListenAndServe(PORT, r))
}
