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

	r := http.NewServeMux()
	routes.RegisterRoutes(r)

	server := http.Server{
		Addr: os.Getenv("PORT"),
		Handler: middleware.Logging(
			middleware.Cors(r),
		),
	}

	log.Println("Starting server...")
	log.Fatal(server.ListenAndServe())
}
