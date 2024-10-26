package main

import (
	"log"
	"net/http"
)

func main() {
	ConnectDatabase()

	http.HandleFunc("/users", UserRegister)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
