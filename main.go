package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/abidkhan03/go_training/db"
	handler "github.com/abidkhan03/go_training/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req Request
	json.NewDecoder(r.Body).Decode(&req)
	defer r.Body.Close()
	response := &Response{
		Code:      200,
		Message:   "Welcome " + req.Name + "!",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

const PORT = ":8000"

func main() {
	err := http.ListenAndServe(":8000", getRouter())
	if err != nil {
		log.Fatal()
		return
	}
	fmt.Println("Server is running on port 8000")

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
	r.Post("/auth/signin", handler.SignIn)
}

func protectedRoutes(r chi.Router) {
	r.Use(jwtauth.Verifier(handler.TokenAuth))
	r.Use(handler.ValidateToken)
	r.Post("/parse", handler.Csv)
	DB := db.DB
	h := handler.New(DB)

	r.Get("/objects", h.GetAllObjects)
	r.Get("/object/{id}", h.GetObjectByID)
	r.Post("/addObject", h.CreateObject)
	r.Patch("/updateObject/{id}", h.UpdateObjectByID)
	r.Delete("/deleteObject/{id}", h.DeleteObjectByID)
}
