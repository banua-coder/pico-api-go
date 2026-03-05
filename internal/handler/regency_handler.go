package handler

import (
	"net/http"
	"strconv"

	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/gorilla/mux"
)

// RegencyHandler handles HTTP requests for regency endpoints
type RegencyHandler struct {
	service service.RegencyServiceInterface
}

// NewRegencyHandler creates a new RegencyHandler
func NewRegencyHandler(service service.RegencyServiceInterface) *RegencyHandler {
	return &RegencyHandler{service: service}
}

// GetRegencies godoc
// @Summary Get all regencies in Sulawesi Tengah
// @Description Returns all kabupaten/kota in Sulawesi Tengah
// @Tags regencies
// @Produce json
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /api/v1/regencies [get]
func (h *RegencyHandler) GetRegencies(w http.ResponseWriter, r *http.Request) {
	regencies, err := h.service.GetRegencies()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, regencies)
}

// GetRegencyByID godoc
// @Summary Get a regency by ID
// @Description Returns a single regency by its ID (kode kabupaten)
// @Tags regencies
// @Produce json
// @Param code path int true "Regency ID/Code"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /api/v1/regencies/{code} [get]
func (h *RegencyHandler) GetRegencyByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["code"])
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid regency code")
		return
	}

	regency, err := h.service.GetRegencyByID(id)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if regency == nil {
		writeErrorResponse(w, http.StatusNotFound, "Kabupaten dengan kode "+vars["code"]+" tidak ditemukan")
		return
	}
	writeSuccessResponse(w, regency)
}

// GetRegencyCases godoc
// @Summary Get daily cases for a regency
// @Description Returns all daily COVID-19 case data for a specific regency
// @Tags regencies
// @Produce json
// @Param code path int true "Regency ID/Code"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /api/v1/regencies/{code}/cases [get]
func (h *RegencyHandler) GetRegencyCases(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["code"])
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Invalid regency code")
		return
	}

	cases, err := h.service.GetRegencyCases(id)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if cases == nil {
		writeErrorResponse(w, http.StatusNotFound, "Tidak ditemukan data untuk kabupaten/kota dengan kode "+vars["code"])
		return
	}
	writeSuccessResponse(w, cases)
}
