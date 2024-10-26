package routes

import (
	"encoding/json"
	"fmt"
	"linksh/backend/database"
	"log"
	"net/http"
	"regexp"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

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

func UserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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
		fmt.Fprint(w, errorMessage)
		return
	}

	var username string
	database.DB.QueryRow("SELECT username FROM users WHERE username = $1", body.Username).Scan(&username)
	if username != "" {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "Username already taken")
		return
	}

	password_hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	database.DB.Query("INSERT INTO users (username, password) VALUES ($1, $2)", body.Username, string(password_hash))
	fmt.Fprint(w, "User successfully registered!")
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

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
		fmt.Fprint(w, errorMessage)
		return
	}

	var user database.User
	database.DB.QueryRow("SELECT username, password FROM users WHERE username = $1", body.Username).Scan(&user.Username, &user.Password)
	if user.Username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid username or password")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid username or password")
		return
	}

	fmt.Fprint(w, "Successfully logged in!")
}
