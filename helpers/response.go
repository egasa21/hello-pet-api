package helpers

import (
	"encoding/json"
	"net/http"
)

type Meta struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func Respond(w http.ResponseWriter, data interface{}, success bool, message string, errorCode string, statusCode int) {
	response := Response{
		Meta: Meta{
			Success:   success,
			Message:   message,
			ErrorCode: errorCode,
		},
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
