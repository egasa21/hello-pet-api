package handler

import (
	"encoding/json"
	"github.com/egasa21/hello-pet-api/helpers"
	"github.com/egasa21/hello-pet-api/models/doctor_model"
	repository "github.com/egasa21/hello-pet-api/repository/doctor"
	"github.com/egasa21/hello-pet-api/request/doctor_request"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type DoctorHandler struct {
	doctorRepository *repository.DoctorRepository
	validate         *validator.Validate
}

func NewDoctorHandler(doctorRepository *repository.DoctorRepository) *DoctorHandler {
	return &DoctorHandler{doctorRepository: doctorRepository, validate: validator.New()}
}

func (h *DoctorHandler) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	var req doctor_request.DoctorRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if err := h.validate.Struct(&req); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	doctor := doctor_model.Doctor{
		Name:     req.Name,
		Age:      req.Age,
		Address:  req.Address,
		Phone:    req.Phone,
		Position: req.Position,
	}

	if err := h.doctorRepository.CreateDoctor(&doctor); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, doctor, true, "Doctor created successfully", "", http.StatusOK)
}

func (h *DoctorHandler) GetDoctor(w http.ResponseWriter, r *http.Request) {
	doctorID := chi.URLParam(r, "doctorID")

	var doctor doctor_model.Doctor
	if err := h.doctorRepository.GetDoctorById(&doctor, doctorID); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "NOT_FOUND", http.StatusNotFound)
		return
	}

	helpers.Respond(w, doctor, true, "Doctor found", "", http.StatusOK)

}

func (h *DoctorHandler) UpdateDoctor(w http.ResponseWriter, r *http.Request) {
	doctorID := chi.URLParam(r, "doctorID")

	var req doctor_request.DoctorRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "BAD_REQUEST", http.StatusBadRequest)
	}

	defer r.Body.Close()

	var doctor doctor_model.Doctor
	if err := h.doctorRepository.GetDoctorById(&doctor, doctorID); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "NOT_FOUND", http.StatusNotFound)
	}

	doctor.Name = req.Name
	doctor.Age = req.Age
	doctor.Phone = req.Phone
	doctor.Address = req.Address
	doctor.Position = req.Position

	if err := h.doctorRepository.UpdateDoctor(&doctor); err != nil {
		helpers.Respond(w, nil, false, err.Error(), "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
		return
	}

	helpers.Respond(w, doctor, true, "Doctor updated successfully", "", http.StatusOK)
}
