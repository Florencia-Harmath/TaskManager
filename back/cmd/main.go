package main

import (
	"log"
	"net/http"

	"taskmanager/internal/database"
	"taskmanager/internal/models"
	"taskmanager/internal/router"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDatabase()
	database.SetupExtension()
	models.MigrateTasks()
	models.MigrateUsers()

	log.Println("Server running on port 3000")
	r := router.InitializeRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Origin"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":3000", handler))
}
