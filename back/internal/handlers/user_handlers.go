package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"taskmanager/internal/database"
	"taskmanager/internal/models"
	"taskmanager/pkg/auth"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// RegisterUser maneja el registro de nuevos usuarios
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.RegisterUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := user.HashPassword(); err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LoginUser maneja el login de usuarios y genera un token JWT
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginRequest models.LoginUser
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.RegisterUser
	result := database.DB.Where("email = ?", loginRequest.Email).First(&user)
	if result.Error != nil {
		log.Println("User not found:", result.Error)
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	} else {
		log.Println("User found:", user)
	}

	if user.CheckPassword(loginRequest.Password) {
		log.Println("Password mismatch")
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	} else {
		log.Println("Password match")
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// obtener un usuario por id
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusBadRequest)
		return
	}

	var user models.RegisterUser
	result := database.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

//modificar un usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr := vars["id"]

    if idStr == "" {
        http.Error(w, "User ID is required", http.StatusBadRequest)
        return
    }

    id, err := uuid.Parse(idStr)
    if err != nil {
        http.Error(w, "Invalid User ID format", http.StatusBadRequest)
        return
    }

    var userToUpdate models.RegisterUser
    result := database.DB.First(&userToUpdate, "id = ?", id)
    if result.Error != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Decode the updated user data from the request body
    var updateRequest models.UpdateUserRequest
    err = json.NewDecoder(r.Body).Decode(&updateRequest)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Update only the Name field
    userToUpdate.Name = updateRequest.Name

    // Save the updated user back to the database
    result = database.DB.Save(&userToUpdate)
    if result.Error != nil {
        http.Error(w, "Failed to update user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(userToUpdate)
}

