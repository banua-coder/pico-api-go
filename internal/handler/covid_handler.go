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
//	@Description	Retrieve national COVID-19 cases data with optional date range filtering
//	@Tags			national
//	@Accept			json
//	@Produce		json
//	@Param			start_date	query		string	false	"Start date (YYYY-MM-DD)"
//	@Param			end_date	query		string	false	"End date (YYYY-MM-DD)"
//	@Success		200			{object}	Response{data=[]models.NationalCaseResponse}
//	@Failure		400			{object}	Response
//	@Failure		500			{object}	Response
//	@Router			/national [get]
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
	
	// Validate pagination params
	limit, offset = utils.ValidatePaginationParams(limit, offset)

	if provinceID == "" {
		// Handle all provinces cases
		if all {
			// Return all data without pagination
			if startDate != "" && endDate != "" {
				cases, err := h.covidService.GetAllProvinceCasesByDateRange(startDate, endDate)
				if err != nil {
					writeErrorResponse(w, http.StatusInternalServerError, err.Error())
					return
				}
				responseData := models.TransformProvinceCaseSliceToResponse(cases)
				writeSuccessResponse(w, responseData)
				return
			}
			
			cases, err := h.covidService.GetAllProvinceCases()
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
			cases, total, err := h.covidService.GetAllProvinceCasesByDateRangePaginated(startDate, endDate, limit, offset)
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
		
		cases, total, err := h.covidService.GetAllProvinceCasesPaginated(limit, offset)
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
			cases, err := h.covidService.GetProvinceCasesByDateRange(provinceID, startDate, endDate)
			if err != nil {
				writeErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}
			responseData := models.TransformProvinceCaseSliceToResponse(cases)
			writeSuccessResponse(w, responseData)
			return
		}
		
		cases, err := h.covidService.GetProvinceCases(provinceID)
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
		cases, total, err := h.covidService.GetProvinceCasesByDateRangePaginated(provinceID, startDate, endDate, limit, offset)
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
	
	cases, total, err := h.covidService.GetProvinceCasesPaginated(provinceID, limit, offset)
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
