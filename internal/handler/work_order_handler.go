package handler

import (
	"encoding/json"
	"net/http"

	"github.com/BioSystems-Indonesia/lis/internal/domain/dto"
	"github.com/BioSystems-Indonesia/lis/internal/usecase"
)

type WorkOrderHandler struct {
	workOrderUC usecase.WorkOrderUsecase
}

func NewWorkOrderHandler(workOrderUC usecase.WorkOrderUsecase) *WorkOrderHandler {
	return &WorkOrderHandler{
		workOrderUC: workOrderUC,
	}
}

func (h *WorkOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req dto.WorkOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	workOrder, err := h.workOrderUC.Create(r.Context(), &req)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusCreated, workOrder)
}

func (h *WorkOrderHandler) GetByNoOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	noOrder := r.URL.Query().Get("no_order")
	if noOrder == "" {
		h.respondError(w, http.StatusBadRequest, "no_order parameter is required")
		return
	}

	workOrder, err := h.workOrderUC.GetByNoOrder(r.Context(), noOrder)
	if err != nil {
		h.respondError(w, http.StatusNotFound, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, workOrder)
}

func (h *WorkOrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	noOrder := r.URL.Query().Get("no_order")
	if noOrder == "" {
		h.respondError(w, http.StatusBadRequest, "no_order parameter is required")
		return
	}

	var req dto.WorkOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	workOrder, err := h.workOrderUC.Update(r.Context(), noOrder, &req)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, workOrder)
}

func (h *WorkOrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	noOrder := r.URL.Query().Get("no_order")
	if noOrder == "" {
		h.respondError(w, http.StatusBadRequest, "no_order parameter is required")
		return
	}

	if err := h.workOrderUC.Delete(r.Context(), noOrder); err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, map[string]string{
		"message": "Work order deleted successfully",
	})
}

func (h *WorkOrderHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	workOrders, err := h.workOrderUC.GetAll(r.Context())
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, workOrders)
}

func (h *WorkOrderHandler) GetByDoctor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	doctor := r.URL.Query().Get("doctor")
	if doctor == "" {
		h.respondError(w, http.StatusBadRequest, "doctor parameter is required")
		return
	}

	workOrders, err := h.workOrderUC.GetByDoctor(r.Context(), doctor)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, workOrders)
}

func (h *WorkOrderHandler) GetByAnalyst(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	analyst := r.URL.Query().Get("analyst")
	if analyst == "" {
		h.respondError(w, http.StatusBadRequest, "analyst parameter is required")
		return
	}

	workOrders, err := h.workOrderUC.GetByAnalyst(r.Context(), analyst)
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondSuccess(w, http.StatusOK, workOrders)
}

func (h *WorkOrderHandler) respondSuccess(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := dto.Response{
		Code:   code,
		Status: "success",
		Data:   data,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *WorkOrderHandler) respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := dto.ResponseError{
		Code:    code,
		Status:  "error",
		Message: message,
	}

	json.NewEncoder(w).Encode(response)
}
