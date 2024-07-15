package routers

import (
	"github.com/egasa21/hello-pet-api/infra/database"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterRoutes(router *chi.Mux, db *database.DB) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	// Add router here
	AuthRoutes(router, db)
	CustomerRoutes(router, db)
	DoctorRoutes(router, db)
}
