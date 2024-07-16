package middlewares

import (
	"context"
	"github.com/egasa21/hello-pet-api/helpers"
	"github.com/egasa21/hello-pet-api/models/customer_model"
	"github.com/egasa21/hello-pet-api/models/user_model"
	"github.com/egasa21/hello-pet-api/repository"
	authRepo "github.com/egasa21/hello-pet-api/repository/auth"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type contextKey string

const UserKey contextKey = "user"
const AdminKey contextKey = "admin"

func IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email, isAdmin, err := helpers.GetCurrentUser(r)
		if err != nil {
			helpers.Respond(w, nil, false, "Unauthorized", "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, map[string]interface{}{
			"email":   email,
			"isAdmin": isAdmin,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(UserKey).(map[string]interface{})
		if !ok || !user["isAdmin"].(bool) {
			helpers.Respond(w, nil, false, "Forbidden", "FORBIDDEN", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LoadUser(authRepo *authRepo.AuthRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			email, _, err := helpers.GetCurrentUser(r)
			if err != nil {
				helpers.Respond(w, nil, false, err.Error(), "UNAUTHORIZED", http.StatusUnauthorized)
				return
			}

			var user user_model.User
			if err := authRepo.FindUserByEmail(&user, email); err != nil {
				helpers.Respond(w, nil, false, err.Error(), "UNAUTHORIZED", http.StatusUnauthorized)
				return
			}

			// add user to request context
			ctx := context.WithValue(r.Context(), UserKey, &user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserFromContext(ctx context.Context) *user_model.User {
	user, ok := ctx.Value(UserKey).(*user_model.User)
	if !ok {
		return nil
	}

	return user
}

func IsOwner(repo repository.Repository, extractOwnerID func(interface{}) uint) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value(UserKey).(*user_model.User)
			if !ok {
				helpers.Respond(w, nil, false, "Failed to get user from context", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
				return
			}

			idStr := chi.URLParam(r, "customerID")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				helpers.Respond(w, nil, false, "Invalid ID format", "BAD_REQUEST", http.StatusBadRequest)
				return
			}

			var resource customer_model.Customer
			if err := repo.GetById(uint(id), &resource); err != nil {
				helpers.Respond(w, nil, false, "Resource not found", "NOT_FOUND", http.StatusNotFound)
				return
			}

			ownerID := extractOwnerID(&resource)
			if user.IsAdmin || ownerID == user.ID {
				next.ServeHTTP(w, r)
			} else {
				helpers.Respond(w, nil, false, "You do not have permission to access this resource", "FORBIDDEN", http.StatusForbidden)
			}
		})
	}
}
