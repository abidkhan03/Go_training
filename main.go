package main

import (
	"github.com/abidkhan03/go_training/jwt"
	"github.com/go-chi/jwtauth"
	"log"
	"net/http"

	"github.com/abidkhan03/go_training/handler"
	"github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
)

const PORT = ":8000"

// main function starts the server and listens to the port
// It creates a new router and adds the handler functions to the router
// It logs the message that the server is running on the port
// It logs the error if the server cannot be started


func main() {
	r := chi.NewRouter()
	r.Post("/hello", handler.HelloHandler)
	r.Post("/auth/signin", handler.SignIn)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(jwt.TokenAuth))
		r.Use(jwt.Authenticator)

		r.Post("/parse", handler.Csv)
	})

	log.Println("Server is running on port " + PORT)
	log.Println(http.ListenAndServe(PORT, r))
}
