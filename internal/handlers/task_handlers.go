package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	
	"taskmanager/internal/database"
	"taskmanager/internal/models"
	"taskmanager/pkg/middleware"

	"github.com/google/uuid"
)

// Función que obtiene todos los tasks de la base de datos y los devuelve en formato JSON.
func GetTasks(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	var tasks []models.Task
	database.DB.Where("user_id = ?", userID).Find(&tasks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tasks)
}

// Función que obtiene un task de la base de datos y lo devuelve en formato JSON.
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	id := r.URL.Query().Get("id")
	var task models.Task
	result := database.DB.Where("user_id = ? AND id = ?", userID, id).First(&task)

	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// Función que inserta un task en la base de datos.
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extraer el ID del usuario desde el contexto que se estableció en el middleware
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	log.Println("User ID:", userID)

	task.UserID = userID

	result := database.DB.Create(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	json.NewEncoder(w).Encode(task)
}



// Función que actualiza un task.
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	id := r.URL.Query().Get("id")

	var task models.Task
	result := database.DB.Where("user_id = ? AND id = ?", userID, id).First(&task)
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Save(&task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// Función que elimina un task de la base de datos.
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	id := r.URL.Query().Get("id")

	var task models.Task
	result := database.DB.Where("user_id = ? AND id = ?", userID, id).Delete(&task)
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// funcion quie cambia de false a true el completed

func CompleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var task models.Task
	result := database.DB.First(&task, "id = ?", id)

	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.Completed = true

	database.DB.Save(&task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// funcion que cambia de true a false el completed
func UncompleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var task models.Task
	result := database.DB.First(&task, "id = ?", id)

	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	task.Completed = false

	database.DB.Save(&task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}