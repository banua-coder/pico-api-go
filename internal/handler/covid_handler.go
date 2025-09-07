package handler

import (
	"net/http"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/database"
	"github.com/gorilla/mux"
)

type CovidHandler struct {
	covidService service.CovidService
	db           *database.DB
}

func NewCovidHandler(covidService service.CovidService, db *database.DB) *CovidHandler {
	return &CovidHandler{
		covidService: covidService,
		db:           db,
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
			// Transform to new response structure
			responseData := models.TransformProvinceCaseSliceToResponse(cases)
			writeSuccessResponse(w, responseData)
			return
		}

		cases, err := h.covidService.GetAllProvinceCases()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		// Transform to new response structure
		responseData := models.TransformProvinceCaseSliceToResponse(cases)
		writeSuccessResponse(w, responseData)
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
		// Transform to new response structure
		responseData := models.TransformProvinceCaseSliceToResponse(cases)
		writeSuccessResponse(w, responseData)
		return
	}

	cases, err := h.covidService.GetProvinceCases(provinceID)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Transform to new response structure
	responseData := models.TransformProvinceCaseSliceToResponse(cases)
	writeSuccessResponse(w, responseData)
}

func (h *CovidHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"status":    "healthy",
		"service":   "COVID-19 API",
		"version":   "2.0.2",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	// Database health check
	dbHealth := map[string]interface{}{
		"status": "healthy",
	}

	if h.db != nil {
		if err := h.db.HealthCheck(); err != nil {
			dbHealth["status"] = "unhealthy"
			dbHealth["error"] = err.Error()
			health["status"] = "degraded"
		} else {
			stats := h.db.GetConnectionStats()
			dbHealth["connections"] = map[string]int{
				"open":       stats.OpenConnections,
				"idle":       stats.Idle,
				"in_use":     stats.InUse,
				"max_open":   stats.MaxOpenConnections,
				"wait_count": int(stats.WaitCount),
			}
		}
	} else {
		dbHealth["status"] = "unavailable"
		dbHealth["error"] = "database connection not initialized"
		health["status"] = "degraded"
	}

	health["database"] = dbHealth

	// Set appropriate HTTP status code based on health status
	statusCode := http.StatusOK
	if health["status"] == "degraded" || health["status"] == "unhealthy" {
		statusCode = http.StatusServiceUnavailable
	}

	writeJSONResponse(w, statusCode, Response{
		Status: "success",
		Data:   health,
	})
}
