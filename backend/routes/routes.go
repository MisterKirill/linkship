package routes

import (
	"linkship/backend/middleware"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", Register)
	mux.HandleFunc("POST /users/login", Login)
	mux.HandleFunc("GET /users/{username}", GetUser)

	mux.Handle("POST /links", middleware.Authentication(http.HandlerFunc(CreateLink)))
	mux.Handle("DELETE /links/{id}", middleware.Authentication(http.HandlerFunc(DeleteLink)))
}