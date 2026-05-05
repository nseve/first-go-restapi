package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nseve/first-go-restapi/internal/response"
)

type contextKey string

const userIDKey contextKey = "userID"

func JWTAuth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {

				response.WriteError(w, http.StatusUnauthorized, "Missing auth header")
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				response.WriteError(w, http.StatusUnauthorized, "Invalid auth header")
				return
			}

			tokenStr := parts[1]

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				response.WriteError(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				response.WriteError(w, http.StatusUnauthorized, "Invalid token claims")
				return
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				response.WriteError(w, http.StatusUnauthorized, "Invalid user_id")
				return
			}

			userID := uint(userIDFloat)

			ctx := context.WithValue(r.Context(), userIDKey, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(r *http.Request) (uint, bool) {
	userID, ok := r.Context().Value(userIDKey).(uint)
	return userID, ok
}
