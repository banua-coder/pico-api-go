package handler

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginationMeta holds pagination metadata
type PaginationMeta struct {
	Page       int  `json:"page"`
	PerPage    int  `json:"per_page"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

// PaginatedResponse wraps paginated list data with metadata
type PaginatedResponse struct {
	Data       interface{}    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// PaginationParams holds parsed pagination parameters
type PaginationParams struct {
	Page    int
	PerPage int
	LoadAll bool
	Offset  int
}

// parsePaginationParams reads page/per_page/load_all query params
func parsePaginationParams(r *http.Request) PaginationParams {
	p := PaginationParams{
		Page:    1,
		PerPage: 10,
	}

	if v := r.URL.Query().Get("load_all"); v == "true" || v == "1" {
		p.LoadAll = true
		return p
	}

	if v := r.URL.Query().Get("page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			p.Page = n
		}
	}
	if v := r.URL.Query().Get("per_page"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			if n > 100 {
				n = 100
			}
			p.PerPage = n
		}
	}

	p.Offset = (p.Page - 1) * p.PerPage
	return p
}

// buildPaginationMeta computes pagination metadata
func buildPaginationMeta(p PaginationParams, total int) PaginationMeta {
	totalPages := int(math.Ceil(float64(total) / float64(p.PerPage)))
	if totalPages < 1 {
		totalPages = 1
	}
	return PaginationMeta{
		Page:       p.Page,
		PerPage:    p.PerPage,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    p.Page < totalPages,
		HasPrev:    p.Page > 1,
	}
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	writeJSONResponse(w, http.StatusOK, Response{
		Status: "success",
		Data:   data,
	})
}

func writePaginatedResponse(w http.ResponseWriter, data interface{}, meta PaginationMeta) {
	writeJSONResponse(w, http.StatusOK, Response{
		Status: "success",
		Data: PaginatedResponse{
			Data:       data,
			Pagination: meta,
		},
	})
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	writeJSONResponse(w, statusCode, Response{
		Status: "error",
		Error:  message,
	})
}
