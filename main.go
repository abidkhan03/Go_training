package main

import (
	"log"
	"net/http"

	"github.com/abidkhan03/go_training/handler"
	"github.com/go-chi/chi/v5"
)

const PORT = ":8000"

// main function starts the server and listens to the port
// It creates a new router and adds the handler functions to the router
// It logs the message that the server is running on the port
// It logs the error if the server cannot be started
func main() {
	r := chi.NewRouter()
	r.Post("/hello", handler.HelloHandler)
	r.Post("/parse", handler.Csv)
	log.Println("Server is running on port " + PORT)
	log.Println(http.ListenAndServe(PORT, r))
}
