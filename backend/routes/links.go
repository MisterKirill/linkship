package routes

import (
	"database/sql"
	"encoding/json"
	"linkship/backend/database"
	"linkship/backend/middleware"
	"net/http"
	"unicode/utf8"
)

func CreateLink(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserKey).(database.User)

	var body struct {
		Name string
		Url  string
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if utf8.RuneCountInString(body.Name) > 50 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Name is too long",
		})
		return
	}

	if utf8.RuneCountInString(body.Url) > 100 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "URL is too long",
		})
		return
	}

	database.DB.Query(
		"INSERT INTO links (user_id, name, url) VALUES ($1, $2, $3)",
		user.Id,
		body.Name,
		body.Url,
	)
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	linkId := r.PathValue("id")
	user := r.Context().Value(middleware.UserKey).(database.User)

	err := database.DB.QueryRow("SELECT * FROM links WHERE id = $1 AND user_id = $2", linkId, user.Id).Scan()
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	database.DB.Query("DELETE FROM links WHERE id = $1 AND user_id = $2", linkId, user.Id)
}
