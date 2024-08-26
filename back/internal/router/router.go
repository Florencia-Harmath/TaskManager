package router

import (
    "taskmanager/internal/handlers"
    "taskmanager/pkg/middleware"
    "github.com/gorilla/mux"
    gorillahandlers "github.com/gorilla/handlers"
)

func InitializeRouter() *mux.Router {
    r := mux.NewRouter()

    // Configura CORS para permitir solicitudes desde http://localhost:5173
    corsMiddleware := gorillahandlers.CORS(
        gorillahandlers.AllowedOrigins([]string{"http://localhost:5173"}),
        gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        gorillahandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    // Rutas públicas (sin autenticación)
    r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
    r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

    // Rutas privadas con autenticación y CORS
    s := r.PathPrefix("/api").Subrouter()
    s.Use(middleware.JwtAuthentication)
    s.Use(corsMiddleware) // Aplica el middleware CORS a las rutas privadas

    s.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
    s.HandleFunc("/tasks/{id}", handlers.GetTaskByID).Methods("GET")
    s.HandleFunc("/createTasks", handlers.CreateTask).Methods("POST")
    s.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
    s.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
    s.HandleFunc("/tasks/{id}/complete", handlers.CompleteTask).Methods("PUT")
    s.HandleFunc("/tasks/{id}/uncomplete", handlers.UncompleteTask).Methods("PUT")

    return r
}
