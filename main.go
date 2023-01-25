package main

import (
	"log"
	"net/http"
	"time"

	"github.com/abidkhan03/go_training/auth"
	"github.com/abidkhan03/go_training/handler"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	// "github.com/go-chi/chi/v5/middleware"
)

const PORT = ":8000"

func main() {
	log.Println("Server is running on port " + PORT)
	log.Println(http.ListenAndServe(PORT, getRouter()))
}

func getRouter() *chi.Mux {
	r := chi.NewRouter()
	// Config
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin", "Cache-Control"},
		ExposedHeaders:   []string{"Content-Type", "JWT-Token", "Content-Disposition"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Routes
	r.Group(publicRoutes)
	r.Group(protectedRoutes)
	return r
}

func publicRoutes(r chi.Router) {
	r.Post("/hello", handler.HelloHandler)
	r.Post("/auth/signin", auth.SignIn)
}

func protectedRoutes(r chi.Router) {
	r.Use(jwtauth.Verifier(auth.TokenAuth))
	r.Use(auth.ValidateToken)
	r.Post("/parse", handler.Csv)
}
