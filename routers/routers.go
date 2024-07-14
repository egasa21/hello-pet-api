package routers

import (
	"github.com/egasa21/hello-pet-api/infra/database"
	"github.com/egasa21/hello-pet-api/routers/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRoute(db *database.DB) *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.Cors())

	RegisterRoutes(router, db)
	return router
}
