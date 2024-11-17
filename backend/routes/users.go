package routes

import (
	"database/sql"
	"encoding/json"
	"linkship/backend/database"
	"linkship/backend/utils"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body struct {
		Username string
		Password string
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errorMessage, err := utils.ValidateUsername(body.Username)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if errorMessage != "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]any{
			"message": errorMessage,
		})
		return
	}

	var username string
	database.DB.QueryRow("SELECT username FROM users WHERE username = $1", body.Username).Scan(&username)
	if username != "" {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Username already taken",
		})
		return
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	database.DB.Query("INSERT INTO users (username, password) VALUES ($1, $2)", body.Username, string(password_hash))

	token, err := utils.GenerateToken(body.Username)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"token": token,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body struct {
		Username string
		Password string
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errorMessage, err := utils.ValidateUsername(body.Username)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if errorMessage != "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]any{
			"message": errorMessage,
		})
		return
	}

	var user database.AuthenticationUser
	err = database.DB.QueryRow("SELECT password FROM users WHERE username = $1", body.Username).Scan(&user.Password)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Invalid username or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]any{
			"message": "Invalid username or password",
		})
		return
	}

	token, err := utils.GenerateToken(body.Username)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"token": token,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := r.PathValue("username")

	errorMessage, err := utils.ValidateUsername(username)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if errorMessage != "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]any{
			"message": errorMessage,
		})
		return
	}

	var user database.User
	err = database.DB.QueryRow("SELECT username, display_name, bio FROM users WHERE username = $1", username).Scan(
		&user.Username,
		&user.DisplayName,
		&user.Bio,
	)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
