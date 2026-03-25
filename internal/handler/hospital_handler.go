package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/gorilla/mux"
)

// HospitalHandler handles HTTP requests for hospital endpoints
type HospitalHandler struct {
	service service.HospitalServiceInterface
}

// NewHospitalHandler creates a new HospitalHandler
func NewHospitalHandler(service service.HospitalServiceInterface) *HospitalHandler {
	return &HospitalHandler{service: service}
}

// GetHospitals godoc
// @Summary Get hospitals in Sulawesi Tengah (paginated)
// @Description Returns paginated list of hospitals. Use ?load_all=true to get all without pagination.
// @Tags hospitals
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all hospitals without pagination"
// @Success 200 {object} Response
// @Router /hospitals [get]
func (h *HospitalHandler) GetHospitals(w http.ResponseWriter, r *http.Request) {
	p := parsePaginationParams(r)

	if p.LoadAll {
		hospitals, err := h.service.GetHospitals()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeSuccessResponse(w, hospitals)
		return
	}

	hospitals, total, err := h.service.GetHospitalsPaginated(p.PerPage, p.Offset)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writePaginatedResponse(w, hospitals, buildPaginationMeta(p, total))
}

// GetHospitalByCode godoc
// @Summary Get a hospital by code
// @Description Returns a single hospital with bed availability and contacts
// @Tags hospitals
// @Produce json
// @Param code path string true "Hospital Code"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /hospitals/{code} [get]
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
