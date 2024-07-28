package handler

import (
	"encoding/json"
	"github.com/egasa21/hello-pet-api/helpers"
	"github.com/egasa21/hello-pet-api/models/product_model"
	repository "github.com/egasa21/hello-pet-api/repository/product"
	"github.com/egasa21/hello-pet-api/request/product_request"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ProductHandler struct {
	productRepository *repository.ProductRepository
}

func NewProductHandler(productRepository *repository.ProductRepository) *ProductHandler {
	return &ProductHandler{productRepository: productRepository}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req product_request.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	product := product_model.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}

	if err := h.productRepository.Create(&product); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, product, true, "Product created successfully", "", http.StatusCreated)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")

	var product product_model.Product
	if err := h.productRepository.FindById(&product, productID); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "NOT_FOUND", http.StatusNotFound)
		return
	}

	helpers.Respond(w, product, true, "Product found", "", http.StatusOK)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "productID")

	var req product_request.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
	}

	defer r.Body.Close()

	var product product_model.Product
	if err := h.productRepository.FindById(&product, productID); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "NOT_FOUND", http.StatusNotFound)
	}

	product.Name = req.Name
	product.Price = req.Price
	product.Stock = req.Stock

	if err := h.productRepository.Update(&product); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, product, true, "Product updated successfully", "", http.StatusOK)
}
