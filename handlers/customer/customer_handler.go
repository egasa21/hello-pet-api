package customer

import (
	"encoding/json"
	"github.com/egasa21/hello-pet-api/helpers"
	"github.com/egasa21/hello-pet-api/models/customer_model"
	"github.com/egasa21/hello-pet-api/models/user_model"
	authRepo "github.com/egasa21/hello-pet-api/repository/auth"
	custRepo "github.com/egasa21/hello-pet-api/repository/customer"
	"github.com/egasa21/hello-pet-api/request/customer_request"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type CustomerHandler struct {
	customerRepository *custRepo.CustomerRepository
	authRepository     *authRepo.AuthRepository
}

func NewCustomerHandler(customerRepository *custRepo.CustomerRepository, authRepo *authRepo.AuthRepository) *CustomerHandler {
	return &CustomerHandler{customerRepository: customerRepository, authRepository: authRepo}
}

func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req customer_request.CustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	email, _, err := helpers.GetCurrentUser(r)
	if err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	var user user_model.User
	if err := h.authRepository.FindUserByEmail(&user, email); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	customer := customer_model.Customer{
		UserID:     user.ID,
		Name:       req.Name,
		Address:    req.Address,
		Phone:      req.Phone,
		AnimalType: req.AnimalType,
	}

	if err := h.customerRepository.CreateCustomer(&customer); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	if err := h.customerRepository.LoadUser(&customer); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, customer, true, "User registered successfully", "", http.StatusCreated)
}

func (h *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "customerID")

	var customer customer_model.Customer
	if err := h.customerRepository.GetCustomerById(&customer, customerID); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "NOT_FOUND", http.StatusNotFound)
		return
	}

	helpers.Respond(w, customer, true, "Customer found", "", http.StatusOK)

}

func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	customerID := chi.URLParam(r, "customerID")

	var req customer_request.CustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
	}

	defer r.Body.Close()

	var customer customer_model.Customer
	if err := h.customerRepository.GetCustomerById(&customer, customerID); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "NOT_FOUND", http.StatusNotFound)
	}

	customer.Name = req.Name
	customer.Address = req.Address
	customer.Phone = req.Phone
	customer.AnimalType = req.AnimalType
	if err := h.customerRepository.UpdateCustomer(&customer); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, customer, true, "Customer updated successfully", "", http.StatusOK)
}
