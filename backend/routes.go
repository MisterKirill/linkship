package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

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

	usernameLength := utf8.RuneCountInString(body.Username)

	if usernameLength < 3 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, "Username is too short")
		return
	}

	if usernameLength > 40 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, "Username is too long")
		return
	}

	usernameValid, err := regexp.Match("^[a-zA-Z0-9_]+$", []byte(body.Username))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !usernameValid {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, "Username can contain only English letters, numbers and underscores")
		return
	}

	var username string
	DB.QueryRow("SELECT username FROM users WHERE username = $1", body.Username).Scan(&username)
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

	DB.Query("INSERT INTO users (username, password) VALUES ($1, $2)", body.Username, string(password_hash))
	fmt.Fprint(w, "User successfully regisered!")
}
