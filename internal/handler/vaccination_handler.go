package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/service"
)

type VaccinationHandler struct {
	service service.VaccinationServiceInterface
}

func NewVaccinationHandler(service service.VaccinationServiceInterface) *VaccinationHandler {
	return &VaccinationHandler{service: service}
}

// GetNationalVaccinations godoc
// @Summary Get national vaccination data
// @Tags vaccination
// @Produce json
// @Success 200 {object} Response
// @Router /api/v1/vaccination/national [get]
func (h *VaccinationHandler) GetNationalVaccinations(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetNationalVaccinations()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, data)
}

// GetProvinceVaccinations godoc
// @Summary Get Sulawesi Tengah vaccination data
// @Tags vaccination
// @Produce json
// @Success 200 {object} Response
// @Router /api/v1/vaccination/province [get]
func (h *VaccinationHandler) GetProvinceVaccinations(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetProvinceVaccinations()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, data)
}

// GetVaccineLocations godoc
// @Summary Get vaccination centers in Sulawesi Tengah
// @Tags vaccination
// @Produce json
// @Success 200 {object} Response
// @Router /api/v1/vaccination/locations [get]
func (h *VaccinationHandler) GetVaccineLocations(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetVaccineLocations()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, data)
}
