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
