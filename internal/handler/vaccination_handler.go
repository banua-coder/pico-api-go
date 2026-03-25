package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/dto"
	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/service"
)

type VaccinationHandler struct {
	service service.VaccinationServiceInterface
}

func NewVaccinationHandler(service service.VaccinationServiceInterface) *VaccinationHandler {
	return &VaccinationHandler{service: service}
}

// GetNationalVaccinations godoc
// @Summary Get national vaccination data (paginated)
// @Description Returns paginated national vaccination records. Use ?load_all=true to get all.
// @Tags vaccination
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all records without pagination"
// @Success 200 {object} Response{data=[]dto.VaccinationResponse}
// @Router /vaccination/national [get]
func (h *VaccinationHandler) GetNationalVaccinations(w http.ResponseWriter, r *http.Request) {
	p := parsePaginationParams(r)

	if p.LoadAll {
		data, err := h.service.GetNationalVaccinations()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeSuccessResponse(w, transformNationalSlice(data))
		return
	}

	data, total, err := h.service.GetNationalVaccinationsPaginated(p.PerPage, p.Offset)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writePaginatedResponse(w, transformNationalSlice(data), buildPaginationMeta(p, total))
}

// GetProvinceVaccinations godoc
// @Summary Get Sulawesi Tengah vaccination data (paginated)
// @Description Returns paginated provincial vaccination records. Use ?load_all=true to get all.
// @Tags vaccination
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all records without pagination"
// @Success 200 {object} Response{data=[]dto.ProvinceVaccinationResponse}
// @Router /vaccination/province [get]
func (h *VaccinationHandler) GetProvinceVaccinations(w http.ResponseWriter, r *http.Request) {
	p := parsePaginationParams(r)

	if p.LoadAll {
		data, err := h.service.GetProvinceVaccinations()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeSuccessResponse(w, transformProvinceSlice(data))
		return
	}

	data, total, err := h.service.GetProvinceVaccinationsPaginated(p.PerPage, p.Offset)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writePaginatedResponse(w, transformProvinceSlice(data), buildPaginationMeta(p, total))
}

// GetVaccineLocations godoc
// @Summary Get vaccination centers in Sulawesi Tengah (paginated)
// @Description Returns paginated list of vaccination locations. Use ?load_all=true to get all.
// @Tags vaccination
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all locations without pagination"
// @Success 200 {object} Response
// @Router /vaccination/locations [get]
func (h *VaccinationHandler) GetVaccineLocations(w http.ResponseWriter, r *http.Request) {
	p := parsePaginationParams(r)

	if p.LoadAll {
		data, err := h.service.GetVaccineLocations()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeSuccessResponse(w, data)
		return
	}

	data, total, err := h.service.GetVaccineLocationsPaginated(p.PerPage, p.Offset)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writePaginatedResponse(w, data, buildPaginationMeta(p, total))
}

func transformNationalSlice(data []models.NationalVaccine) []dto.VaccinationResponse {
	result := make([]dto.VaccinationResponse, len(data))
	for i, v := range data {
		result[i] = dto.TransformNationalVaccine(v)
	}
	return result
}

func transformProvinceSlice(data []models.ProvinceVaccine) []dto.ProvinceVaccinationResponse {
	result := make([]dto.ProvinceVaccinationResponse, len(data))
	for i, v := range data {
		result[i] = dto.TransformProvinceVaccine(v)
	}
	return result
}
