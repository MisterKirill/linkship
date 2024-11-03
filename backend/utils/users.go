package utils

import (
	"os"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateUsername(username string) (string, error) {
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

func GenerateToken(username string) (string, error) {
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
