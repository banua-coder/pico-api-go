package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/handler"
	"github.com/banua-coder/pico-api-go/internal/middleware"
	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type MockNationalCaseRepo struct {
	mock.Mock
}

func (m *MockNationalCaseRepo) GetAll() ([]models.NationalCase, error) {
	args := m.Called()
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

func (m *MockNationalCaseRepo) GetByDateRange(startDate, endDate time.Time) ([]models.NationalCase, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

func (m *MockNationalCaseRepo) GetLatest() (*models.NationalCase, error) {
	args := m.Called()
	return args.Get(0).(*models.NationalCase), args.Error(1)
}

func (m *MockNationalCaseRepo) GetByDay(day int64) (*models.NationalCase, error) {
	args := m.Called(day)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.NationalCase), args.Error(1)
}

type MockProvinceRepo struct {
	mock.Mock
}

func (m *MockProvinceRepo) GetAll() ([]models.Province, error) {
	args := m.Called()
	return args.Get(0).([]models.Province), args.Error(1)
}

func (m *MockProvinceRepo) GetByID(id string) (*models.Province, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.Province), args.Error(1)
}

type MockProvinceCaseRepo struct {
	mock.Mock
}

func (m *MockProvinceCaseRepo) GetAll() ([]models.ProvinceCaseWithDate, error) {
	args := m.Called()
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepo) GetByProvinceID(provinceID string) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepo) GetByProvinceIDAndDateRange(provinceID string, startDate, endDate time.Time) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID, startDate, endDate)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepo) GetByDateRange(startDate, endDate time.Time) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepo) GetLatestByProvinceID(provinceID string) (*models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.ProvinceCaseWithDate), args.Error(1)
}

func setupTestServer() (*httptest.Server, *MockNationalCaseRepo, *MockProvinceRepo, *MockProvinceCaseRepo) {
	mockNationalRepo := new(MockNationalCaseRepo)
	mockProvinceRepo := new(MockProvinceRepo)
	mockProvinceCaseRepo := new(MockProvinceCaseRepo)

	covidService := service.NewCovidService(mockNationalRepo, mockProvinceRepo, mockProvinceCaseRepo)
	router := handler.SetupRoutes(covidService)

	router.Use(middleware.Recovery)
	router.Use(middleware.CORS)

	server := httptest.NewServer(router)
	return server, mockNationalRepo, mockProvinceRepo, mockProvinceCaseRepo
}

func TestAPI_HealthCheck(t *testing.T) {
	server, _, _, _ := setupTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/api/v1/health")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	var response handler.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	data, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "healthy", data["status"])
	assert.Equal(t, "COVID-19 API", data["service"])
}

func TestAPI_GetNationalCases(t *testing.T) {
	server, mockNationalRepo, _, _ := setupTestServer()
	defer server.Close()

	now := time.Now()
	rt := 1.2
	expectedCases := []models.NationalCase{
		{
			ID:       1,
			Day:      1,
			Date:     now,
			Positive: 100,
			Recovered: 80,
			Deceased: 5,
			Rt:       &rt,
		},
	}

	mockNationalRepo.On("GetAll").Return(expectedCases, nil)

	resp, err := http.Get(server.URL + "/api/v1/national")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response handler.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)
	assert.NotNil(t, response.Data)

	mockNationalRepo.AssertExpectations(t)
}

func TestAPI_GetNationalCasesWithDateRange(t *testing.T) {
	server, mockNationalRepo, _, _ := setupTestServer()
	defer server.Close()

	startDate := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expectedCases := []models.NationalCase{
		{ID: 1, Date: startDate, Positive: 100},
	}

	mockNationalRepo.On("GetByDateRange", startDate, endDate).Return(expectedCases, nil)

	resp, err := http.Get(server.URL + "/api/v1/national?start_date=2020-03-01&end_date=2020-03-31")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response handler.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	mockNationalRepo.AssertExpectations(t)
}

func TestAPI_GetLatestNationalCase(t *testing.T) {
	server, mockNationalRepo, _, _ := setupTestServer()
	defer server.Close()

	expectedCase := &models.NationalCase{
		ID:       1,
		Positive: 100,
		Date:     time.Now(),
	}

	mockNationalRepo.On("GetLatest").Return(expectedCase, nil)

	resp, err := http.Get(server.URL + "/api/v1/national/latest")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response handler.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	mockNationalRepo.AssertExpectations(t)
}

func TestAPI_GetProvinces(t *testing.T) {
	server, _, mockProvinceRepo, _ := setupTestServer()
	defer server.Close()

	expectedProvinces := []models.Province{
		{ID: "11", Name: "Aceh"},
		{ID: "31", Name: "DKI Jakarta"},
	}

	mockProvinceRepo.On("GetAll").Return(expectedProvinces, nil)

	resp, err := http.Get(server.URL + "/api/v1/provinces")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response handler.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	mockProvinceRepo.AssertExpectations(t)
}

func TestAPI_GetProvinceCases(t *testing.T) {
	server, _, _, mockProvinceCaseRepo := setupTestServer()
	defer server.Close()

	expectedCases := []models.ProvinceCaseWithDate{
		{
			ProvinceCase: models.ProvinceCase{
				ID:         1,
				ProvinceID: "11",
				Positive:   50,
			},
			Date: time.Now(),
		},
	}

	mockProvinceCaseRepo.On("GetAll").Return(expectedCases, nil)

	resp, err := http.Get(server.URL + "/api/v1/provinces/cases")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response handler.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)

	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestAPI_CORS(t *testing.T) {
	server, _, _, _ := setupTestServer()
	defer server.Close()

	req, err := http.NewRequest("OPTIONS", server.URL+"/api/v1/health", nil)
	assert.NoError(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "*", resp.Header.Get("Access-Control-Allow-Origin"))
	assert.Contains(t, resp.Header.Get("Access-Control-Allow-Methods"), "GET")
}

func TestAPI_InvalidEndpoint(t *testing.T) {
	server, _, _, _ := setupTestServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/api/v1/invalid")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}