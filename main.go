package main

import (
	"log"
	"net/http"

	"github.com/abidkhan03/go_training/handler"
	"github.com/go-chi/chi/v5"
)

const PORT = ":8000"

func main() {
	r := chi.NewRouter()
	r.Post("/hello", handler.HelloHandler)

	log.Println("Server is running on port ", PORT)
	log.Println(http.ListenAndServe(PORT, r))
}
