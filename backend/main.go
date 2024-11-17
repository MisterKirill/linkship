package main

import (
	"linkship/backend/database"
	"linkship/backend/middleware"
	"linkship/backend/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: failed to load .env file")
	}

	database.ConnectDatabase()

	http.HandleFunc("/users", middleware.CorsMiddleware(routes.UserRegister))
	http.HandleFunc("/users/login", middleware.CorsMiddleware(routes.UserLogin))
	http.HandleFunc("/users/", middleware.CorsMiddleware(routes.GetUser))

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
