package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/gorilla/mux"
)

type CovidHandler struct {
	covidService service.CovidService
}

func NewCovidHandler(covidService service.CovidService) *CovidHandler {
	return &CovidHandler{
		covidService: covidService,
	}
}

func (h *CovidHandler) GetNationalCases(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate != "" && endDate != "" {
		cases, err := h.covidService.GetNationalCasesByDateRange(startDate, endDate)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		// Transform to new response structure
		responseData := models.TransformSliceToResponse(cases)
		writeSuccessResponse(w, responseData)
		return
	}

	cases, err := h.covidService.GetNationalCases()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Transform to new response structure
	responseData := models.TransformSliceToResponse(cases)
	writeSuccessResponse(w, responseData)
}

func (h *CovidHandler) GetLatestNationalCase(w http.ResponseWriter, r *http.Request) {
	nationalCase, err := h.covidService.GetLatestNationalCase()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if nationalCase == nil {
		writeErrorResponse(w, http.StatusNotFound, "No national case data found")
		return
	}

	// Transform to new response structure
	responseData := nationalCase.TransformToResponse()
	writeSuccessResponse(w, responseData)
}

func (h *CovidHandler) GetProvinces(w http.ResponseWriter, r *http.Request) {
	provinces, err := h.covidService.GetProvinces()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeSuccessResponse(w, provinces)
}

func (h *CovidHandler) GetProvinceCases(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provinceID := vars["provinceId"]

	if provinceID == "" {
		startDate := r.URL.Query().Get("start_date")
		endDate := r.URL.Query().Get("end_date")

		if startDate != "" && endDate != "" {
			cases, err := h.covidService.GetAllProvinceCasesByDateRange(startDate, endDate)
			if err != nil {
				writeErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			writeSuccessResponse(w, cases)
			return
		}

		cases, err := h.covidService.GetAllProvinceCases()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeSuccessResponse(w, cases)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate != "" && endDate != "" {
		cases, err := h.covidService.GetProvinceCasesByDateRange(provinceID, startDate, endDate)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeSuccessResponse(w, cases)
		return
	}

	cases, err := h.covidService.GetProvinceCases(provinceID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeSuccessResponse(w, cases)
}

func (h *CovidHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeSuccessResponse(w, map[string]string{
		"status":  "healthy",
		"service": "COVID-19 API",
		"version": "1.0.0",
	})
}