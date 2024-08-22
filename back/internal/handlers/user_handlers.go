package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"taskmanager/internal/database"
	"taskmanager/internal/models"
	"taskmanager/pkg/auth"
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

