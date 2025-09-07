package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCovidService struct {
	mock.Mock
}

func (m *MockCovidService) GetNationalCases() ([]models.NationalCase, error) {
	args := m.Called()
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

func (m *MockCovidService) GetNationalCasesByDateRange(startDate, endDate string) ([]models.NationalCase, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

func (m *MockCovidService) GetLatestNationalCase() (*models.NationalCase, error) {
	args := m.Called()
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.NationalCase), args.Error(1)
}

func (m *MockCovidService) GetProvinces() ([]models.Province, error) {
	args := m.Called()
	return args.Get(0).([]models.Province), args.Error(1)
}

func (m *MockCovidService) GetProvinceCases(provinceID string) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockCovidService) GetProvinceCasesByDateRange(provinceID, startDate, endDate string) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID, startDate, endDate)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockCovidService) GetAllProvinceCases() ([]models.ProvinceCaseWithDate, error) {
	args := m.Called()
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockCovidService) GetAllProvinceCasesByDateRange(startDate, endDate string) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func TestCovidHandler_GetNationalCases(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.NationalCase{
		{ID: 1, Positive: 100, Recovered: 80, Deceased: 5},
	}

	mockService.On("GetNationalCases").Return(expectedCases, nil)

	req, err := http.NewRequest("GET", "/api/v1/national", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetNationalCases(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)
	assert.NotNil(t, response.Data)

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetNationalCases_WithDateRange(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.NationalCase{
		{ID: 1, Positive: 100, Date: time.Date(2020, 3, 15, 0, 0, 0, 0, time.UTC)},
	}

	mockService.On("GetNationalCasesByDateRange", "2020-03-01", "2020-03-31").Return(expectedCases, nil)

	req, err := http.NewRequest("GET", "/api/v1/national?start_date=2020-03-01&end_date=2020-03-31", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetNationalCases(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetNationalCases_ServiceError(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	mockService.On("GetNationalCases").Return([]models.NationalCase{}, errors.New("database error"))

	req, err := http.NewRequest("GET", "/api/v1/national", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetNationalCases(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.Status)
	assert.Contains(t, response.Error, "database error")

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetLatestNationalCase(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCase := &models.NationalCase{ID: 1, Positive: 100}
	mockService.On("GetLatestNationalCase").Return(expectedCase, nil)

	req, err := http.NewRequest("GET", "/api/v1/national/latest", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetLatestNationalCase(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetLatestNationalCase_NotFound(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	mockService.On("GetLatestNationalCase").Return((*models.NationalCase)(nil), nil)

	req, err := http.NewRequest("GET", "/api/v1/national/latest", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetLatestNationalCase(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.Status)
	assert.Contains(t, response.Error, "No national case data found")

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinces(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedProvinces := []models.Province{
		{ID: "11", Name: "Aceh"},
		{ID: "31", Name: "DKI Jakarta"},
	}

	mockService.On("GetProvinces").Return(expectedProvinces, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetProvinces(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinceCases_AllProvinces(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}},
	}

	mockService.On("GetAllProvinceCases").Return(expectedCases, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces/cases", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetProvinceCases(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinceCases_SpecificProvince(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}},
	}

	mockService.On("GetProvinceCases", "11").Return(expectedCases, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces/11/cases", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/provinces/{provinceId}/cases", handler.GetProvinceCases)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	mockService.AssertExpectations(t)
}

func TestCovidHandler_HealthCheck(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	req, err := http.NewRequest("GET", "/api/v1/health", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.HealthCheck(rr, req)

	assert.Equal(t, http.StatusServiceUnavailable, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	data, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "degraded", data["status"])
	assert.Equal(t, "COVID-19 API", data["service"])
	assert.Equal(t, "2.0.0", data["version"])
	assert.Contains(t, data, "database")
	
	dbData, ok := data["database"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "unavailable", dbData["status"])
}