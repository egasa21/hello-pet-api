package routers

import (
	handlers "github.com/egasa21/hello-pet-api/handlers/auth"
	"github.com/egasa21/hello-pet-api/infra/database"
	repository "github.com/egasa21/hello-pet-api/repository/auth"
	"github.com/go-chi/chi/v5"
)

func AuthRoutes(router *chi.Mux, db *database.DB) {
	repo := repository.NewAuthRepository(db)
	authHandler := handlers.NewAuthHandler(repo)

	router.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	})
}
