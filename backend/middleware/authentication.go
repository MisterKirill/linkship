package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"linkship/backend/database"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

const UserKey ContextKey = "user"

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		sub, err := claims.GetSubject()
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		var user database.User
		err = database.DB.QueryRow("SELECT username, display_name, bio FROM users WHERE username = $1", sub).Scan(
			&user.Username,
			&user.DisplayName,
			&user.Bio,
		)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
