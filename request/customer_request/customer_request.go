package customer_request

import "github.com/egasa21/hello-pet-api/request/auth_request"

type CustomerRequest struct {
	UserID     uint                      `json:"user_id"`
	User       auth_request.LoginRequest `json:"user,omitempty"`
	Name       string                    `json:"name"`
	Address    string                    `json:"address"`
	Phone      string                    `json:"phone"`
	AnimalType string                    `json:"animal_type"`
}
