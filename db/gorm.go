package db

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Object struct {
	ID int `json:"id" db:"primary_key"`
	Name string `json:"name"`
	Description string `json:"description"`
}


func (h handler)GetAllObjects(w http.ResponseWriter, r *http.Request) {
	var objects []Object
	if result := h.DB.Find(&objects); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(objects)
}

func (h handler) GetObjectByID(w http.ResponseWriter, r *http.Request) {
	var object Object
	vars := mux.Vars(r)
	id := vars["id"]
	if result := h.DB.First(&object, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(object)
}

func (h handler) CreateObject(w http.ResponseWriter, r *http.Request) {
	//defer r.Body.Close()
	var object Object
	//json.NewDecoder(r.Body).Decode(&object)
	if result := h.DB.Create(&object); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(object)
}

func (h handler) UpdateObjectByID(w http.ResponseWriter, r *http.Request) {
	// Read the id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body


	var updatedObject Object
	//json.Unmarshal(body, &updatedObject)

	var object Object

	if result := h.DB.First(&object, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	object.ID = updatedObject.ID
	object.Name = updatedObject.Name
	object.Description = updatedObject.Description

	h.DB.Save(&object)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}

func (h handler) DeleteObjectByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var object Object
	if result := h.DB.First(&object, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	h.DB.Delete(&object)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
