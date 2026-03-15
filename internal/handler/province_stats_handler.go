package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/service"
)

type ProvinceStatsHandler struct {
	service service.ProvinceStatsServiceInterface
}

func NewProvinceStatsHandler(service service.ProvinceStatsServiceInterface) *ProvinceStatsHandler {
	return &ProvinceStatsHandler{service: service}
}

// GetGenderCases godoc
// @Summary Get COVID-19 cases by gender and age group in Sulawesi Tengah
// @Tags province-stats
// @Produce json
// @Success 200 {object} Response
// @Router /api/v1/stats/gender [get]
func (h *ProvinceStatsHandler) GetGenderCases(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetGenderCases()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, data)
}

// GetLatestGenderCase godoc
// @Summary Get latest gender/age case data for Sulawesi Tengah
// @Tags province-stats
// @Produce json
// @Success 200 {object} Response
// @Router /api/v1/stats/gender/latest [get]
func (h *ProvinceStatsHandler) GetLatestGenderCase(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetLatestGenderCase()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, data)
}

// GetTests godoc
// @Summary Get COVID-19 test data for Sulawesi Tengah
// @Tags province-stats
// @Produce json
// @Success 200 {object} Response
// @Router /api/v1/stats/tests [get]
func (h *ProvinceStatsHandler) GetTests(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetTests()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, data)
}

// GetTestTypes godoc
// @Summary Get available COVID-19 test types
// @Tags province-stats
// @Produce json
// @Success 200 {object} Response
// @Router /api/v1/stats/test-types [get]
func (h *ProvinceStatsHandler) GetTestTypes(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetTestTypes()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, data)
}
