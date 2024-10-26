package main

import (
	"log"
	"net/http"

	"linksh/backend/database"
	"linksh/backend/routes"
)

func main() {
	database.ConnectDatabase()

	http.HandleFunc("/users", routes.UserRegister)
	http.HandleFunc("/users/login", routes.UserLogin)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
