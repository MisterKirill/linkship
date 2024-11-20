package routes

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", Register)
	mux.HandleFunc("POST /users/login", Login)
	mux.HandleFunc("GET /users/{username}", GetUser)
}
