package main

import (
	"fmt"
	"net/http"

	"github.com/abidkhan03/go_training/handler"
	"github.com/go-chi/chi/v5"
)

const PORT = ":8000"

func main() {
	r := chi.NewRouter()
	r.Post("/hello", handler.HelloHandler)

	fmt.Println("Server is running on port ", PORT)
	fmt.Println(http.ListenAndServe(PORT, r))
}
