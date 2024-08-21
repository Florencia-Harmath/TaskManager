package main

import (
	"log"
	"net/http"

	"taskmanager/internal/database"
	"taskmanager/internal/models"
	"taskmanager/internal/router"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDatabase()
	models.MigrateTasks()
	models.MigrateUsers()

	log.Println("Server running on port 3000")
	r := router.InitializeRouter()
	log.Fatal(http.ListenAndServe(":3000", r))
}