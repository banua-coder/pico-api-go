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
// @Summary Get regencies in Sulawesi Tengah (paginated)
// @Description Returns paginated kabupaten/kota list. Use ?load_all=true to get all.
// @Tags regencies
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all regencies without pagination"
// @Success 200 {object} Response
// @Failure 500 {object} Response
// @Router /regencies [get]
func (h *RegencyHandler) GetRegencies(w http.ResponseWriter, r *http.Request) {
	p := parsePaginationParams(r)

	if p.LoadAll {
		regencies, err := h.service.GetRegencies()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeSuccessResponse(w, regencies)
		return
	}

	regencies, total, err := h.service.GetRegenciesPaginated(p.PerPage, p.Offset)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writePaginatedResponse(w, regencies, buildPaginationMeta(p, total))
}

// GetRegencyByID godoc
// @Summary Get a regency by ID
// @Description Returns a single regency by its ID (kode kabupaten)
// @Tags regencies
// @Produce json
// @Param code path int true "Regency ID/Code"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /regencies/{code} [get]
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
// @Router /regencies/{code}/cases [get]
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
