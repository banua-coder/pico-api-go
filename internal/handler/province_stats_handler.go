package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/dto"
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
// @Description Returns all COVID-19 case records grouped by gender (male/female) and age groups (0-14, 15-19, 20-24, 25-49, 50-54, 55+) for positive and PDP categories.
// @Tags province-stats
// @Produce json
// @Success 200 {object} Response{data=[]dto.GenderStatsResponse}
// @Failure 500 {object} Response
// @Router /stats/gender [get]
func (h *ProvinceStatsHandler) GetGenderCases(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetGenderCases()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, dto.ToGenderStatsResponseList(data))
}

// GetLatestGenderCase godoc
// @Summary Get latest gender/age case data for Sulawesi Tengah
// @Description Returns the most recent COVID-19 case record grouped by gender and age groups for positive and PDP categories.
// @Tags province-stats
// @Produce json
// @Success 200 {object} Response{data=dto.GenderStatsResponse}
// @Failure 500 {object} Response
// @Router /stats/gender/latest [get]
func (h *ProvinceStatsHandler) GetLatestGenderCase(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetLatestGenderCase()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if data == nil {
		writeSuccessResponse(w, nil)
		return
	}
	writeSuccessResponse(w, dto.ToGenderStatsResponse(*data))
}

// GetTests godoc
// @Summary Get COVID-19 test data for Sulawesi Tengah
// @Tags province-stats
// @Produce json
// @Success 200 {object} Response
// @Router /stats/tests [get]
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
// @Router /stats/test-types [get]
func (h *ProvinceStatsHandler) GetTestTypes(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetTestTypes()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, data)
}
