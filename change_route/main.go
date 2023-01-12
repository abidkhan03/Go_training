package main

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
)

const port = "8000"
func main() {
	r := chi.NewRouter()
	r.Post("/hello", HelloHandler)
	fmt.Println("Server is running on port ", port)
	http.ListenAndServe(port, r)
}
