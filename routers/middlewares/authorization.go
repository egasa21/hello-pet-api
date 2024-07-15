package middlewares

import (
	"context"
	"github.com/egasa21/hello-pet-api/helpers"
	"net/http"
	"strings"
)

type contextKey string

const UserKey contextKey = "user"

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			helpers.Respond(w, nil, false, "No Authorization", "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			helpers.Respond(w, nil, false, "Invalid Authorization header format", "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		_, err := helpers.ParseToken(tokenStr)
		if err != nil {
			helpers.Respond(w, nil, false, "Invalid token", "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, authHeader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//todo add admin role
