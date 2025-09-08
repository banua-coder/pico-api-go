package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/utils"
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

func (m *MockCovidService) GetProvincesWithLatestCase() ([]models.ProvinceWithLatestCase, error) {
	args := m.Called()
	return args.Get(0).([]models.ProvinceWithLatestCase), args.Error(1)
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

// Paginated methods
func (m *MockCovidService) GetProvinceCasesPaginated(provinceID string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(provinceID, limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockCovidService) GetProvinceCasesByDateRangePaginated(provinceID, startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(provinceID, startDate, endDate, limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockCovidService) GetAllProvinceCasesPaginated(limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockCovidService) GetAllProvinceCasesByDateRangePaginated(startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(startDate, endDate, limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

// Sorted methods
func (m *MockCovidService) GetNationalCasesSorted(sortParams utils.SortParams) ([]models.NationalCase, error) {
	args := m.Called(sortParams)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

func (m *MockCovidService) GetNationalCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.NationalCase, error) {
	args := m.Called(startDate, endDate, sortParams)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

func (m *MockCovidService) GetProvinceCasesSorted(provinceID string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockCovidService) GetProvinceCasesPaginatedSorted(provinceID string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(provinceID, limit, offset, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockCovidService) GetProvinceCasesByDateRangeSorted(provinceID, startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID, startDate, endDate, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockCovidService) GetProvinceCasesByDateRangePaginatedSorted(provinceID, startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(provinceID, startDate, endDate, limit, offset, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockCovidService) GetAllProvinceCasesSorted(sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockCovidService) GetAllProvinceCasesPaginatedSorted(limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(limit, offset, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockCovidService) GetAllProvinceCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(startDate, endDate, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockCovidService) GetAllProvinceCasesByDateRangePaginatedSorted(startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(startDate, endDate, limit, offset, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func TestCovidHandler_GetNationalCases(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.NationalCase{
		{ID: 1, Positive: 100, Recovered: 80, Deceased: 5},
	}

	mockService.On("GetNationalCasesSorted", utils.SortParams{Field: "date", Order: "asc"}).Return(expectedCases, nil)

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

	mockService.On("GetNationalCasesByDateRangeSorted", "2020-03-01", "2020-03-31", utils.SortParams{Field: "date", Order: "asc"}).Return(expectedCases, nil)

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

	mockService.On("GetNationalCasesSorted", utils.SortParams{Field: "date", Order: "asc"}).Return([]models.NationalCase{}, errors.New("database error"))

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

	expectedProvinces := []models.ProvinceWithLatestCase{
		{
			Province: models.Province{ID: "11", Name: "Aceh"},
			LatestCase: &models.ProvinceCaseResponse{
				Day: 100,
				Daily: models.ProvinceDailyCases{
					Positive: 10,
				},
			},
		},
		{
			Province: models.Province{ID: "31", Name: "DKI Jakarta"},
			LatestCase: &models.ProvinceCaseResponse{
				Day: 101,
				Daily: models.ProvinceDailyCases{
					Positive: 25,
				},
			},
		},
	}

	mockService.On("GetProvincesWithLatestCase").Return(expectedProvinces, nil)

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

func TestCovidHandler_GetProvinceCases_AllProvinces_Paginated(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}},
	}
	expectedTotal := 100

	mockService.On("GetAllProvinceCasesPaginatedSorted", 50, 0, utils.SortParams{Field: "date", Order: "asc"}).Return(expectedCases, expectedTotal, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces/cases", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetProvinceCases(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	// Verify paginated response structure
	paginatedData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, paginatedData, "data")
	assert.Contains(t, paginatedData, "pagination")

	pagination, ok := paginatedData["pagination"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(50), pagination["limit"])
	assert.Equal(t, float64(0), pagination["offset"])
	assert.Equal(t, float64(100), pagination["total"])
	assert.Equal(t, float64(1), pagination["page"])
	assert.Equal(t, true, pagination["has_next"])
	assert.Equal(t, false, pagination["has_prev"])

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinceCases_SpecificProvince_Paginated(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}},
	}
	expectedTotal := 50

	mockService.On("GetProvinceCasesPaginatedSorted", "11", 50, 0, utils.SortParams{Field: "date", Order: "asc"}).Return(expectedCases, expectedTotal, nil)

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

	// Verify paginated response structure
	paginatedData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, paginatedData, "data")
	assert.Contains(t, paginatedData, "pagination")

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinceCases_AllData(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}},
		{ProvinceCase: models.ProvinceCase{ID: 2, ProvinceID: "31", Positive: 100}},
	}

	mockService.On("GetAllProvinceCasesSorted", utils.SortParams{Field: "date", Order: "asc"}).Return(expectedCases, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces/cases?all=true", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetProvinceCases(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	// Verify non-paginated response structure (direct array)
	responseArray, ok := response.Data.([]interface{})
	assert.True(t, ok)
	assert.Len(t, responseArray, 2)

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinceCases_CustomPagination(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 3, ProvinceID: "12", Positive: 25}},
	}
	expectedTotal := 200

	mockService.On("GetAllProvinceCasesPaginatedSorted", 100, 50, utils.SortParams{Field: "date", Order: "asc"}).Return(expectedCases, expectedTotal, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces/cases?limit=100&offset=50", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetProvinceCases(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	// Verify custom pagination metadata
	paginatedData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)

	pagination, ok := paginatedData["pagination"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(100), pagination["limit"])
	assert.Equal(t, float64(50), pagination["offset"])
	assert.Equal(t, float64(200), pagination["total"])
	assert.Equal(t, float64(1), pagination["page"])
	assert.Equal(t, true, pagination["has_next"])
	assert.Equal(t, true, pagination["has_prev"])

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinceCases_DateRange_Paginated(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}},
	}
	expectedTotal := 30

	mockService.On("GetAllProvinceCasesByDateRangePaginatedSorted", "2024-01-01", "2024-01-31", 50, 0, utils.SortParams{Field: "date", Order: "asc"}).Return(expectedCases, expectedTotal, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces/cases?start_date=2024-01-01&end_date=2024-01-31", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetProvinceCases(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	// Verify paginated response
	paginatedData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, paginatedData, "pagination")

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinceCases_DateRange_AllData(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}},
	}

	mockService.On("GetAllProvinceCasesByDateRangeSorted", "2024-01-01", "2024-01-31", utils.SortParams{Field: "date", Order: "asc"}).Return(expectedCases, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces/cases?start_date=2024-01-01&end_date=2024-01-31&all=true", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetProvinceCases(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	// Verify non-paginated response structure
	responseArray, ok := response.Data.([]interface{})
	assert.True(t, ok)
	assert.Len(t, responseArray, 1)

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinceCases_SpecificProvince_AllData(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "31", Positive: 200}},
	}

	mockService.On("GetProvinceCasesSorted", "31", utils.SortParams{Field: "date", Order: "asc"}).Return(expectedCases, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces/31/cases?all=true", nil)
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

	// Verify non-paginated response structure
	responseArray, ok := response.Data.([]interface{})
	assert.True(t, ok)
	assert.Len(t, responseArray, 1)

	mockService.AssertExpectations(t)
}

func TestCovidHandler_GetProvinces_ExcludeLatestCase(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	expectedProvinces := []models.Province{
		{ID: "11", Name: "Aceh"},
		{ID: "31", Name: "DKI Jakarta"},
	}

	mockService.On("GetProvinces").Return(expectedProvinces, nil)

	req, err := http.NewRequest("GET", "/api/v1/provinces?exclude_latest_case=true", nil)
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

func TestCovidHandler_GetAPIIndex(t *testing.T) {
	mockService := new(MockCovidService)
	handler := NewCovidHandler(mockService, nil)

	req, err := http.NewRequest("GET", "/api/v1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.GetAPIIndex(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	// Verify structure contains expected keys
	data, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, data, "api")
	assert.Contains(t, data, "documentation")
	assert.Contains(t, data, "endpoints")
	assert.Contains(t, data, "features")
	assert.Contains(t, data, "examples")

	// Verify API info
	apiInfo, ok := data["api"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "Sulawesi Tengah COVID-19 Data API", apiInfo["title"])
	assert.Equal(t, "2.3.0", apiInfo["version"])

	// Verify endpoints structure
	endpoints, ok := data["endpoints"].(map[string]interface{})
	assert.True(t, ok)
	assert.Contains(t, endpoints, "health")
	assert.Contains(t, endpoints, "national")
	assert.Contains(t, endpoints, "provinces")
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
	assert.Equal(t, "2.3.0", data["version"])
	assert.Contains(t, data, "database")

	dbData, ok := data["database"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "unavailable", dbData["status"])
}
