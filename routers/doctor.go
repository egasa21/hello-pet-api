package routers

import (
	handler "github.com/egasa21/hello-pet-api/handlers/doctor"
	"github.com/egasa21/hello-pet-api/infra/database"
	repository "github.com/egasa21/hello-pet-api/repository/doctor"
	"github.com/egasa21/hello-pet-api/routers/middlewares"
	"github.com/go-chi/chi/v5"
)

func DoctorRoutes(router *chi.Mux, db *database.DB) {
	doctorRepository := repository.NewDoctorRepository(db)
	doctorHandler := handler.NewDoctorHandler(doctorRepository)

	router.Route("/api/doctors", func(r chi.Router) {
		r.Use(middlewares.IsAuthorized)
		r.Use(middlewares.IsAdmin)
		r.Post("/create", doctorHandler.CreateDoctor)
		r.Get("/{doctorID}", doctorHandler.GetDoctor)
		r.Put("/{doctorID}", doctorHandler.UpdateDoctor)
		r.Delete("/{doctorID}", doctorHandler.DeleteDoctor)
	})

}
