package router

import (
	"taskmanager/internal/handlers"
	"taskmanager/pkg/middleware"

	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	r := mux.NewRouter()

	// Rutas públicas (sin autenticación)
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Rutas privadas con autenticación y CORS
	s := r.PathPrefix("/api").Subrouter()
	s.Use(middleware.JwtAuthentication)

	s.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	s.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	s.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")
	s.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	s.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	s.HandleFunc("/tasks/{id}/complete", handlers.CompleteTask).Methods("PUT")
	s.HandleFunc("/tasks/{id}/incomplete", handlers.UncompleteTask).Methods("PUT")

	return r
}
