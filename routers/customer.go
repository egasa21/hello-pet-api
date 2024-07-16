package routers

import (
	handlers "github.com/egasa21/hello-pet-api/handlers/customer"
	"github.com/egasa21/hello-pet-api/infra/database"
	"github.com/egasa21/hello-pet-api/models/customer_model"
	authRepo "github.com/egasa21/hello-pet-api/repository/auth"
	custRepo "github.com/egasa21/hello-pet-api/repository/customer"
	"github.com/egasa21/hello-pet-api/routers/middlewares"
	"github.com/go-chi/chi/v5"
)

func CustomerRoutes(router *chi.Mux, db *database.DB) {
	customerRepository := custRepo.NewCustomerRepository(db)
	authRepository := authRepo.NewAuthRepository(db)
	customerHandler := handlers.NewCustomerHandler(customerRepository, authRepository)

	router.Route("/api/customers", func(r chi.Router) {
		r.Use(middlewares.IsAuthorized)
		r.Post("/create", customerHandler.CreateCustomer)

		r.Route("/{customerID}", func(r chi.Router) {
			r.Use(middlewares.LoadUser(authRepository))
			r.Use(middlewares.IsOwner(customerRepository, func(resource interface{}) uint {
				customer, ok := resource.(*customer_model.Customer)
				if !ok {
					return 0
				}
				return customer.UserID
			}))

			r.Get("/", customerHandler.GetCustomer)
			r.Put("/", customerHandler.UpdateCustomer)
		})
	})
}
