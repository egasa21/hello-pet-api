package routers

import (
	handlers "github.com/egasa21/hello-pet-api/handlers/customer"
	"github.com/egasa21/hello-pet-api/infra/database"
	authRepo "github.com/egasa21/hello-pet-api/repository/auth"
	custRepo "github.com/egasa21/hello-pet-api/repository/customer"
	"github.com/go-chi/chi/v5"
)

func CustomerRoutes(router *chi.Mux, db *database.DB) {
	customerRepository := custRepo.NewCustomerRepository(db)
	authRepository := authRepo.NewAuthRepository(db)
	customerHandler := handlers.NewCustomerHandler(customerRepository, authRepository)

	router.Route("/api/customers", func(r chi.Router) {
		r.Post("/create", customerHandler.CreateCustomer)
		r.Get("/{customerID}", customerHandler.GetCustomer)
		r.Put("/{customerID}", customerHandler.UpdateCustomer)
	})
}
