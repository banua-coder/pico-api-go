package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/gorilla/mux"
)

// HospitalHandler handles HTTP requests for hospital endpoints
type HospitalHandler struct {
	service *service.HospitalService
}

// NewHospitalHandler creates a new HospitalHandler
func NewHospitalHandler(service *service.HospitalService) *HospitalHandler {
	return &HospitalHandler{service: service}
}

// GetHospitals godoc
// @Summary Get all hospitals in Sulawesi Tengah
// @Description Returns all hospitals with bed availability and contacts
// @Tags hospitals
// @Produce json
// @Success 200 {object} Response
// @Router /api/v1/hospitals [get]
func (h *HospitalHandler) GetHospitals(w http.ResponseWriter, r *http.Request) {
	hospitals, err := h.service.GetHospitals()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, hospitals)
}

// GetHospitalByCode godoc
// @Summary Get a hospital by code
// @Description Returns a single hospital with bed availability and contacts
// @Tags hospitals
// @Produce json
// @Param code path string true "Hospital Code"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /api/v1/hospitals/{code} [get]
func (h *HospitalHandler) GetHospitalByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	hospital, err := h.service.GetHospitalByCode(code)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if hospital == nil {
		writeErrorResponse(w, http.StatusNotFound, "Rumah sakit dengan kode "+code+" tidak ditemukan")
		return
	}
	writeSuccessResponse(w, hospital)
}
