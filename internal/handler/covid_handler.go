package handler

import (
	"net/http"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/database"
	"github.com/banua-coder/pico-api-go/pkg/utils"
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

// GetNationalCases godoc
//
//	@Summary		Get national COVID-19 cases
//	@Description	Retrieve national COVID-19 cases data with optional date range filtering and sorting
//	@Tags			national
//	@Accept			json
//	@Produce		json
//	@Param			start_date	query		string	false	"Start date (YYYY-MM-DD)"
//	@Param			end_date	query		string	false	"End date (YYYY-MM-DD)"
//	@Param			sort		query		string	false	"Sort by field:order (e.g., date:desc, positive:asc). Default: date:asc"
//	@Success		200			{object}	Response{data=[]models.NationalCaseResponse}
//	@Failure		400			{object}	Response
//	@Failure		429			{object}	Response	"Rate limit exceeded"
//	@Failure		500			{object}	Response
//	@Header			200			{string}	X-RateLimit-Limit		"Request limit per window"
//	@Header			200			{string}	X-RateLimit-Remaining	"Requests remaining in current window"
//	@Header			429			{string}	X-RateLimit-Reset		"Unix timestamp when rate limit resets"
//	@Header			429			{string}	Retry-After				"Seconds to wait before retrying"
//	@Router			/national [get]
func (h *CovidHandler) GetNationalCases(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	// Parse sort parameters (default: date ascending)
	sortParams := utils.ParseSortParam(r, "date")

	if startDate != "" && endDate != "" {
		cases, err := h.covidService.GetNationalCasesByDateRangeSorted(startDate, endDate, sortParams)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		// Transform to new response structure
		responseData := models.TransformSliceToResponse(cases)
		writeSuccessResponse(w, responseData)
		return
	}

	cases, err := h.covidService.GetNationalCasesSorted(sortParams)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Transform to new response structure
	responseData := models.TransformSliceToResponse(cases)
	writeSuccessResponse(w, responseData)
}

// GetLatestNationalCase godoc
//
//	@Summary		Get latest national COVID-19 case
//	@Description	Retrieve the most recent national COVID-19 case data
//	@Tags			national
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Response{data=models.NationalCaseResponse}
//	@Failure		404	{object}	Response
//	@Failure		500	{object}	Response
//	@Router			/national/latest [get]
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

// GetProvinces godoc
//
//	@Summary		Get provinces with COVID-19 data
//	@Description	Retrieve all provinces with their latest COVID-19 case data by default. Use exclude_latest_case=true for basic province list only.
//	@Tags			provinces
//	@Accept			json
//	@Produce		json
//	@Param			exclude_latest_case	query		boolean	false	"Exclude latest case data (default: false)"
//	@Success		200					{object}	Response{data=[]models.ProvinceWithLatestCase}	"Provinces with latest case data"
//	@Success		200					{object}	Response{data=[]models.Province}				"Basic province list when exclude_latest_case=true"
//	@Failure		500					{object}	Response
//	@Router			/provinces [get]
func (h *CovidHandler) GetProvinces(w http.ResponseWriter, r *http.Request) {
	// Check if exclude_latest_case query parameter is set to get basic province list only
	excludeLatestCase := r.URL.Query().Get("exclude_latest_case") == "true"

	if excludeLatestCase {
		provinces, err := h.covidService.GetProvinces()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeSuccessResponse(w, provinces)
		return
	}

	// Default behavior: include latest case data for COVID-19 context
	provincesWithCases, err := h.covidService.GetProvincesWithLatestCase()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, provincesWithCases)
}

// GetProvinceCases godoc
//
//	@Summary		Get province COVID-19 cases
//	@Description	Retrieve COVID-19 cases for all provinces or a specific province with hybrid pagination support
//	@Tags			province-cases
//	@Accept			json
//	@Produce		json
//	@Param			provinceId	path		string	false	"Province ID (e.g., '31' for Jakarta)"
//	@Param			limit		query		integer	false	"Records per page (default: 50, max: 1000)"
//	@Param			offset		query		integer	false	"Records to skip (default: 0)"
//	@Param			all			query		boolean	false	"Return all data without pagination"
//	@Param			start_date	query		string	false	"Start date (YYYY-MM-DD)"
//	@Param			end_date	query		string	false	"End date (YYYY-MM-DD)"
//	@Param			sort		query		string	false	"Sort by field:order (e.g., date:desc, positive:asc). Default: date:asc"
//	@Success		200			{object}	Response{data=models.PaginatedResponse{data=[]models.ProvinceCaseResponse}}	"Paginated response"
//	@Success		200			{object}	Response{data=[]models.ProvinceCaseResponse}							"All data response when all=true"
//	@Failure		400			{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/provinces/cases [get]
//	@Router			/provinces/{provinceId}/cases [get]
func (h *CovidHandler) GetProvinceCases(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provinceID := vars["provinceId"]

	// Parse query parameters
	limit := utils.ParseIntQueryParam(r, "limit", 50)
	offset := utils.ParseIntQueryParam(r, "offset", 0)
	all := utils.ParseBoolQueryParam(r, "all")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	// Parse sort parameters (default: date ascending)
	sortParams := utils.ParseSortParam(r, "date")

	// Validate pagination params
	limit, offset = utils.ValidatePaginationParams(limit, offset)

	if provinceID == "" {
		// Handle all provinces cases
		if all {
			// Return all data without pagination
			if startDate != "" && endDate != "" {
				cases, err := h.covidService.GetAllProvinceCasesByDateRangeSorted(startDate, endDate, sortParams)
				if err != nil {
					writeErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}
				responseData := models.TransformProvinceCaseSliceToResponse(cases)
				writeSuccessResponse(w, responseData)
				return
			}

			cases, err := h.covidService.GetAllProvinceCasesSorted(sortParams)
			if err != nil {
				writeErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			responseData := models.TransformProvinceCaseSliceToResponse(cases)
			writeSuccessResponse(w, responseData)
			return
		}

		// Return paginated data
		if startDate != "" && endDate != "" {
			cases, total, err := h.covidService.GetAllProvinceCasesByDateRangePaginatedSorted(startDate, endDate, limit, offset, sortParams)
			if err != nil {
				writeErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			responseData := models.TransformProvinceCaseSliceToResponse(cases)
			pagination := models.CalculatePaginationMeta(limit, offset, total)
			paginatedResponse := models.PaginatedResponse{
				Data:       responseData,
				Pagination: pagination,
			}
			writeSuccessResponse(w, paginatedResponse)
			return
		}

		cases, total, err := h.covidService.GetAllProvinceCasesPaginatedSorted(limit, offset, sortParams)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		responseData := models.TransformProvinceCaseSliceToResponse(cases)
		pagination := models.CalculatePaginationMeta(limit, offset, total)
		paginatedResponse := models.PaginatedResponse{
			Data:       responseData,
			Pagination: pagination,
		}
		writeSuccessResponse(w, paginatedResponse)
		return
	}

	// Handle specific province cases
	if all {
		// Return all data without pagination
		if startDate != "" && endDate != "" {
			cases, err := h.covidService.GetProvinceCasesByDateRangeSorted(provinceID, startDate, endDate, sortParams)
			if err != nil {
				writeErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			responseData := models.TransformProvinceCaseSliceToResponse(cases)
			writeSuccessResponse(w, responseData)
			return
		}

		cases, err := h.covidService.GetProvinceCasesSorted(provinceID, sortParams)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		responseData := models.TransformProvinceCaseSliceToResponse(cases)
		writeSuccessResponse(w, responseData)
		return
	}

	// Return paginated data
	if startDate != "" && endDate != "" {
		cases, total, err := h.covidService.GetProvinceCasesByDateRangePaginatedSorted(provinceID, startDate, endDate, limit, offset, sortParams)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		responseData := models.TransformProvinceCaseSliceToResponse(cases)
		pagination := models.CalculatePaginationMeta(limit, offset, total)
		paginatedResponse := models.PaginatedResponse{
			Data:       responseData,
			Pagination: pagination,
		}
		writeSuccessResponse(w, paginatedResponse)
		return
	}

	cases, total, err := h.covidService.GetProvinceCasesPaginatedSorted(provinceID, limit, offset, sortParams)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseData := models.TransformProvinceCaseSliceToResponse(cases)
	pagination := models.CalculatePaginationMeta(limit, offset, total)
	paginatedResponse := models.PaginatedResponse{
		Data:       responseData,
		Pagination: pagination,
	}
	writeSuccessResponse(w, paginatedResponse)
}

// HealthCheck godoc
//
//	@Summary		Health check
//	@Description	Check API health status and database connectivity
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Response{data=map[string]interface{}}	"API is healthy"
//	@Success		503	{object}	Response{data=map[string]interface{}}	"API is degraded (database issues)"
//	@Router			/health [get]
func (h *CovidHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"status":    "healthy",
		"service":   "COVID-19 API",
		"version":   "2.4.0",
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

// GetAPIIndex godoc
//
//	@Summary		API endpoint index
//	@Description	Get a list of all available API endpoints with descriptions
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Response{data=map[string]interface{}}
//	@Router			/ [get]
func (h *CovidHandler) GetAPIIndex(w http.ResponseWriter, r *http.Request) {
	endpoints := map[string]interface{}{
		"api": map[string]interface{}{
			"title":       "Sulawesi Tengah COVID-19 Data API",
			"version":     "2.4.0",
			"description": "A comprehensive REST API for COVID-19 data in Sulawesi Tengah (Central Sulawesi)",
		},
		"documentation": map[string]interface{}{
			"swagger_ui": "/swagger/index.html",
			"openapi": map[string]string{
				"yaml": "/docs/swagger.yaml",
				"json": "/docs/swagger.json",
			},
		},
		"endpoints": map[string]interface{}{
			"health": map[string]interface{}{
				"url":         "/api/v1/health",
				"method":      "GET",
				"description": "Check API health status and database connectivity",
			},
			"national": map[string]interface{}{
				"list": map[string]string{
					"url":         "/api/v1/national",
					"method":      "GET",
					"description": "Get national COVID-19 cases (with optional date range)",
				},
				"latest": map[string]string{
					"url":         "/api/v1/national/latest",
					"method":      "GET",
					"description": "Get latest national COVID-19 case data",
				},
			},
			"provinces": map[string]interface{}{
				"list": map[string]string{
					"url":         "/api/v1/provinces",
					"method":      "GET",
					"description": "Get provinces with latest case data (default)",
				},
				"cases": map[string]interface{}{
					"all": map[string]string{
						"url":         "/api/v1/provinces/cases",
						"method":      "GET",
						"description": "Get province cases (paginated by default, ?all=true for complete data)",
					},
					"specific": map[string]string{
						"url":         "/api/v1/provinces/{provinceId}/cases",
						"method":      "GET",
						"description": "Get cases for specific province (e.g., /api/v1/provinces/72/cases for Sulawesi Tengah)",
					},
				},
			},
		},
		"features": []string{
			"Hybrid pagination system (paginated by default, ?all=true for complete data)",
			"Date range filtering (?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD)",
			"Enhanced ODP/PDP data grouping",
			"Provinces with latest case data by default",
			"Sulawesi Tengah focused with national context data",
		},
		"examples": map[string]interface{}{
			"sulawesi_tengah_cases": "/api/v1/provinces/72/cases",
			"paginated_data":        "/api/v1/provinces/cases?limit=100&offset=50",
			"date_range":            "/api/v1/national?start_date=2024-01-01&end_date=2024-12-31",
			"complete_dataset":      "/api/v1/provinces/cases?all=true",
		},
	}

	writeSuccessResponse(w, endpoints)
}
