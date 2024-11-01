package routes

import (
	"linkship/backend/database"
	"linkship/backend/middleware"
	"log"
	"net/http"
)

func SecuredRoute(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.ContextUser).(database.User)
	log.Print(user.Username)
}
