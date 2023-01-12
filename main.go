package main

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/abidkhan03/go_training/change_route"
)

const PORT = "8000"
func main() {
	r := chi.NewRouter()
	r.Post("/hello", change_route.HelloHandler)
	fmt.Println("Server is running on port ", PORT)
	http.ListenAndServe(PORT, r)
}
