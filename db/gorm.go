package db

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Object struct {
	ID int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Description string `json:"description" db:"description" `
}


func (h handler)GetAllObjects(w http.ResponseWriter, r *http.Request) {
	var objects []Object
	if result := h.DB.Find(&objects); result.Error != nil {
		log.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(objects)
}

func (h handler) GetObjectByID(w http.ResponseWriter, r *http.Request) {
	var object Object
	id := chi.URLParam(r, "id")
	if result := h.DB.First(&object, id); result.Error != nil {
		log.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(object)
}

func (h handler) CreateObject(w http.ResponseWriter, r *http.Request) {
	var object Object
	json.NewDecoder(r.Body).Decode(&object)
	if result := h.DB.Create(&object); result.Error != nil {
		log.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(object)
}

func (h handler) UpdateObjectByID(w http.ResponseWriter, r *http.Request) {
	// Read the id parameter
	vars := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(vars)

	var updatedObject Object
	json.NewDecoder(r.Body).Decode(&updatedObject)

	var object Object
	if result := h.DB.First(&object, id); result.Error != nil {
		log.Println(result.Error)
	}
	object.ID = updatedObject.ID
	object.Name = updatedObject.Name
	object.Description = updatedObject.Description

	// check if it returns an error
	if result := h.DB.Save(&object); result.Error!= nil {
        log.Println(result.Error)
    }

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func (h handler) DeleteObjectByID(w http.ResponseWriter, r *http.Request) {
	vars := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(vars)
	var object Object
	if result := h.DB.First(&object, id); result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusNotFound)
        return
    }
	h.DB.Delete(&object)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
