package main

import (
	"log"
	"net/http"
	"os"

	"linksh/backend/database"
	"linksh/backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDatabase()

	http.HandleFunc("/users", routes.UserRegister)
	http.HandleFunc("/users/login", routes.UserLogin)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
