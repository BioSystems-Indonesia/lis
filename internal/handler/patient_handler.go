package handler

import (
	"encoding/json"
	"net/http"

	"github.com/BioSystems-Indonesia/lis/internal/domain/dto"
	"github.com/BioSystems-Indonesia/lis/internal/usecase"
)

type PatientHandler struct {
	patientUC usecase.PatientUsecase
}

func NewPatientHandler(patientUC usecase.PatientUsecase) *PatientHandler {
	return &PatientHandler{
		patientUC: patientUC,
	}
}

func (h *PatientHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.PatientRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	patient, err := h.patientUC.Create(r.Context(), &req)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusCreated, patient)
}

func (h *PatientHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.respondError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	patient, err := h.patientUC.GetByID(r.Context(), id)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, patient)
}

func (h *PatientHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.respondError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	var req dto.PatientRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	patient, err := h.patientUC.Update(r.Context(), id, &req)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, patient)
}

func (h *PatientHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.respondError(w, http.StatusBadRequest, "id parameter is required")
		return
	}

	err := h.patientUC.Delete(r.Context(), id)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, map[string]string{
		"message": "Patient deleted successfully",
	})
}

func (h *PatientHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	patients, err := h.patientUC.GetAll(r.Context())
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, patients)
}

func (h *PatientHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		h.respondError(w, http.StatusBadRequest, "q parameter is required")
		return
	}

	patients, err := h.patientUC.Search(r.Context(), query)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, patients)
}

func (h *PatientHandler) respondSuccess(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := dto.Response{
		Code:   code,
		Status: "success",
		Data:   data,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *PatientHandler) respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := dto.ResponseError{
		Code:    code,
		Status:  "error",
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
