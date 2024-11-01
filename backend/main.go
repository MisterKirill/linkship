package main

import (
	"log"
	"net/http"
	"os"

	"linkship/backend/database"
	"linkship/backend/middleware"
	"linkship/backend/routes"

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

	http.HandleFunc("/secured", middleware.AuthenticationMiddleware(routes.SecuredRoute))

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
