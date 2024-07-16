package handlers

import (
	"encoding/json"
	"github.com/egasa21/hello-pet-api/helpers"
	"github.com/egasa21/hello-pet-api/models/user_model"
	repository "github.com/egasa21/hello-pet-api/repository/auth"
	"github.com/egasa21/hello-pet-api/request/auth_request"
	"net/http"
)

type AuthHandler struct {
	authRepo *repository.AuthRepository
}

func NewAuthHandler(authRepo *repository.AuthRepository) *AuthHandler {
	return &AuthHandler{authRepo: authRepo}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var regisReq auth_request.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&regisReq); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//	validate input
	if regisReq.Email == "" || regisReq.Password == "" {
		helpers.Respond(w, nil, false, "Email and password are required", "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	hashedPassword, err := helpers.HashPassword(regisReq.Password)
	if err != nil {
		helpers.Respond(w, nil, false, "Error hashing password", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	user := user_model.User{
		Username: regisReq.Username,
		Email:    regisReq.Email,
		Password: hashedPassword,
	}

	if err := h.authRepo.Register(&user); err != nil {
		helpers.Respond(w, nil, false, "Error registering user", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, user, true, "User registered successfully", "", http.StatusCreated)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq auth_request.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var user user_model.User
	if err := h.authRepo.FindUserByEmail(&user, loginReq.Email); err != nil {
		helpers.Respond(w, nil, false, "user not found", "USER_NOT_FOUND", http.StatusBadRequest)
		return
	}

	if !helpers.CheckPasswordHash(loginReq.Password, user.Password) {
		helpers.Respond(w, nil, false, "Invalid passworx", "UNAUTHORIZED", http.StatusUnauthorized)
		return
	}

	token, err := helpers.CreateAccessToken(user.Email, user.IsAdmin)
	if err != nil {
		helpers.Respond(w, nil, false, "Error creating access token", "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, map[string]string{"access_token": token}, true, "Login Successfully", "", http.StatusOK)
}
