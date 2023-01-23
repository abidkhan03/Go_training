package main

import (
	"encoding/json"
	"fmt"
	"github.com/abidkhan03/go_training/db"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
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
	DB := db.DB
	h := db.New(DB)
	r := chi.NewRouter()
	r.Post("/hello", HelloHandler)

	r.Get("/objects", h.GetAllObjects)
	r.Get("/object/{id}", h.GetObjectByID)
	r.Post("/addObject", h.CreateObject)
	r.Patch("/updateObject/{id}", h.UpdateObjectByID)
	r.Delete("/deleteObject/{id}", h.DeleteObjectByID)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
	fmt.Println("Server is running on port 8000")

}
