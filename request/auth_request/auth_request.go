package auth_request

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
