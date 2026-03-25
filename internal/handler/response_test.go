package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteJSONResponse(t *testing.T) {
	rr := httptest.NewRecorder()

	response := Response{
		Status:  "success",
		Message: "Operation successful",
		Data:    map[string]string{"key": "value"},
	}

	writeJSONResponse(rr, http.StatusOK, response)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var result Response
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, "success", result.Status)
	assert.Equal(t, "Operation successful", result.Message)
}

func TestWriteSuccessResponse(t *testing.T) {
	rr := httptest.NewRecorder()

	data := map[string]interface{}{
		"count": 5,
		"items": []string{"item1", "item2"},
	}

	writeSuccessResponse(rr, data)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response Response
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)
	assert.NotNil(t, response.Data)
}

func TestWriteErrorResponse(t *testing.T) {
	rr := httptest.NewRecorder()

	errorMessage := "Something went wrong"

	writeErrorResponse(rr, http.StatusBadRequest, errorMessage)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	var response Response
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.Status)
	assert.Equal(t, errorMessage, response.Error)
	assert.Empty(t, response.Message)
	assert.Nil(t, response.Data)
}

func TestWriteErrorResponse_InternalServerError(t *testing.T) {
	rr := httptest.NewRecorder()

	errorMessage := "Database connection failed"

	writeErrorResponse(rr, http.StatusInternalServerError, errorMessage)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var response Response
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.Status)
	assert.Equal(t, errorMessage, response.Error)
}

func TestWritePaginatedResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	data := []string{"item1", "item2"}
	meta := PaginationMeta{Page: 1, PerPage: 10, Total: 2, TotalPages: 1, HasNext: false, HasPrev: false}

	writePaginatedResponse(rr, data, meta)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response Response
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)
	assert.NotNil(t, response.Data)
}

func TestParsePaginationParams_Defaults(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	p := parsePaginationParams(req)
	assert.Equal(t, 1, p.Page)
	assert.Equal(t, 10, p.PerPage)
	assert.Equal(t, 0, p.Offset)
	assert.False(t, p.LoadAll)
}

func TestParsePaginationParams_Custom(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/?page=3&per_page=20", nil)
	p := parsePaginationParams(req)
	assert.Equal(t, 3, p.Page)
	assert.Equal(t, 20, p.PerPage)
	assert.Equal(t, 40, p.Offset)
}

func TestParsePaginationParams_LoadAll(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/?load_all=true", nil)
	p := parsePaginationParams(req)
	assert.True(t, p.LoadAll)
}

func TestParsePaginationParams_LoadAllNumeric(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/?load_all=1", nil)
	p := parsePaginationParams(req)
	assert.True(t, p.LoadAll)
}

func TestParsePaginationParams_PerPageCapped(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/?per_page=999", nil)
	p := parsePaginationParams(req)
	assert.Equal(t, 100, p.PerPage)
}

func TestParsePaginationParams_InvalidValues(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/?page=abc&per_page=xyz", nil)
	p := parsePaginationParams(req)
	assert.Equal(t, 1, p.Page)
	assert.Equal(t, 10, p.PerPage)
}

func TestBuildPaginationMeta(t *testing.T) {
	p := PaginationParams{Page: 2, PerPage: 10}
	meta := buildPaginationMeta(p, 25)
	assert.Equal(t, 2, meta.Page)
	assert.Equal(t, 10, meta.PerPage)
	assert.Equal(t, 25, meta.Total)
	assert.Equal(t, 3, meta.TotalPages)
	assert.True(t, meta.HasNext)
	assert.True(t, meta.HasPrev)
}

func TestBuildPaginationMeta_LastPage(t *testing.T) {
	p := PaginationParams{Page: 3, PerPage: 10}
	meta := buildPaginationMeta(p, 25)
	assert.False(t, meta.HasNext)
	assert.True(t, meta.HasPrev)
}

func TestBuildPaginationMeta_SinglePage(t *testing.T) {
	p := PaginationParams{Page: 1, PerPage: 10}
	meta := buildPaginationMeta(p, 5)
	assert.Equal(t, 1, meta.TotalPages)
	assert.False(t, meta.HasNext)
	assert.False(t, meta.HasPrev)
}

func TestBuildPaginationMeta_Empty(t *testing.T) {
	p := PaginationParams{Page: 1, PerPage: 10}
	meta := buildPaginationMeta(p, 0)
	assert.Equal(t, 1, meta.TotalPages)
	assert.False(t, meta.HasNext)
}
