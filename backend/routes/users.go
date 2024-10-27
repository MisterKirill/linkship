package routes

import (
	"encoding/json"
	"linksh/backend/database"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var body struct {
		Username string
		Password string
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errorMessage, err := validateUsername(body.Username)
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

	token, err := generateToken(body.Username)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"token": token,
	})
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	var body struct {
		Username string
		Password string
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errorMessage, err := validateUsername(body.Username)
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
	database.DB.QueryRow("SELECT password FROM users WHERE username = $1", body.Username).Scan(&user.Password)
	if user.Password == "" {
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

	token, err := generateToken(body.Username)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"token": token,
	})
}

func validateUsername(username string) (string, error) {
	usernameLength := utf8.RuneCountInString(username)

	if usernameLength < 3 {
		return "Username is too short", nil
	}

	if usernameLength > 40 {
		return "Username is too long", nil
	}

	usernameValid, err := regexp.Match("^[a-zA-Z0-9_]+$", []byte(username))
	if err != nil {
		return "", err
	}

	if !usernameValid {
		return "Username can contain only English letters, numbers and underscores", nil
	}

	return "", nil
}

func generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iat": time.Now().Unix(),
		"exp": time.Now().AddDate(0, 1, 0).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
