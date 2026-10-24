package handler

import (
	"net/http"
	"strconv"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/utils"
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
// @Param include_latest_case query bool false "Set true to add a latest_case object (most recent non-zero case row, with date and rt) to each regency. Covers every regency that has regency_cases rows, in any province; regencies without data get null"
// @Success 200 {object} Response{data=[]models.RegencyWithLatestCase} "Regencies with latest_case when include_latest_case=true"
// @Failure 500 {object} Response
// @Router /regencies [get]
func (h *RegencyHandler) GetRegencies(w http.ResponseWriter, r *http.Request) {
	p := parsePaginationParams(r)
	withLatest := utils.ParseBoolQueryParam(r, "include_latest_case")

	if p.LoadAll {
		regencies, err := h.service.GetRegencies()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		if withLatest {
			h.writeWithLatest(w, regencies, nil)
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
	if withLatest {
		meta := buildPaginationMeta(p, total)
		h.writeWithLatest(w, regencies, &meta)
		return
	}
	writePaginatedResponse(w, regencies, buildPaginationMeta(p, total))
}

// writeWithLatest attaches each regency's latest case (cached in the service) and writes the response.
// A nil meta means an unpaginated list.
func (h *RegencyHandler) writeWithLatest(w http.ResponseWriter, regencies []models.Regency, meta *PaginationMeta) {
	latest, err := h.service.GetLatestRegencyCases()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	byRegency := make(map[int]*models.RegencyCase, len(latest))
	for i := range latest {
		c := latest[i]
		c.Regency = nil // already represented by the parent object
		byRegency[c.RegencyID] = &c
	}
	out := make([]models.RegencyWithLatestCase, len(regencies))
	for i, reg := range regencies {
		out[i] = models.RegencyWithLatestCase{Regency: reg, LatestCase: byRegency[reg.ID]}
	}
	if meta == nil {
		writeSuccessResponse(w, out)
		return
	}
	writePaginatedResponse(w, out, *meta)
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
