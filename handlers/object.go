package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/abidkhan03/go_training/db/models"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type ObjectHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) ObjectHandler {
	return ObjectHandler{db}
}

func (h ObjectHandler) GetAllObjects(w http.ResponseWriter, r *http.Request) {
	var objects []models.Object
	if result := h.DB.Find(&objects); result.Error != nil {
		// In case of error, response code should be according to the error and respond body should contain proper error message
		log.Println(result.Error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(objects)
}

func (h ObjectHandler) GetObjectByID(w http.ResponseWriter, r *http.Request) {
	var object models.Object
	id := chi.URLParam(r, "id")
	// validate the id and if it is nil, respond with a proper error message and error code
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":true, "message":"id is required"}`))
		return
	}
	if result := h.DB.First(&object, id); result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(object)
}

func (h ObjectHandler) CreateObject(w http.ResponseWriter, r *http.Request) {
	var object models.Object
	json.NewDecoder(r.Body).Decode(&object)
	// object is empty or any field in the object is empty respond error message and error code
	if object.Name == "" || object.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":true, "message":"name and description are required"}`))
		return
	}

	if result := h.DB.Create(&object); result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(object)
}

func (h ObjectHandler) UpdateObjectByID(w http.ResponseWriter, r *http.Request) {
	// Read the id parameter
	vars := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(vars)
	// validate the id
	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":true, "message":"id is required"}`))
		return
	}

	var updatedObject models.Object
	json.NewDecoder(r.Body).Decode(&updatedObject)
	// object is empty or any field in the object is empty respond error message and error code
	if updatedObject.Name == "" || updatedObject.Description == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":true, "message":"name and description are required"}`))
		return
	}

	var object models.Object
	// if object is not present in the response body then return
	if result := h.DB.First(&object, id); result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	object.ID = updatedObject.ID
	object.Name = updatedObject.Name
	object.Description = updatedObject.Description

	// check if it returns an error
	if result := h.DB.Save(&object); result.Error != nil {
		log.Println(result.Error)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func (h ObjectHandler) DeleteObjectByID(w http.ResponseWriter, r *http.Request) {
	vars := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(vars)
	// validate the id
	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":true, "message":"id is not found..."}`))
		return
	}
	var object models.Object
	// object is empty or any field in the object is empty respond error message and error code
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
