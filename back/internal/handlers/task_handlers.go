package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"taskmanager/internal/database"
	"taskmanager/internal/models"
	"taskmanager/pkg/middleware"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
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
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	// Obtener el UserID del contexto
	userID, ok := r.Context().Value(middleware.UserIDKey).(uuid.UUID)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// Convertir el ID de la tarea al tipo adecuado si es necesario
	taskID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid Task ID format", http.StatusBadRequest)
		return
	}

	var task models.Task
	result := database.DB.Where("user_id = ? AND id = ?", userID, taskID).First(&task)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
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

	vars := mux.Vars(r)
	id := vars["id"]

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
	vars := mux.Vars(r)
	id := vars["id"]

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
	vars := mux.Vars(r)
	id := vars["id"]

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
