package service

import (
	"errors"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockNationalCaseRepository struct {
	mock.Mock
}

func (m *MockNationalCaseRepository) GetAll() ([]models.NationalCase, error) {
	args := m.Called()
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

func (m *MockNationalCaseRepository) GetByDateRange(startDate, endDate time.Time) ([]models.NationalCase, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

func (m *MockNationalCaseRepository) GetLatest() (*models.NationalCase, error) {
	args := m.Called()
	return args.Get(0).(*models.NationalCase), args.Error(1)
}

func (m *MockNationalCaseRepository) GetByDay(day int64) (*models.NationalCase, error) {
	args := m.Called(day)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.NationalCase), args.Error(1)
}

func (m *MockNationalCaseRepository) GetAllSorted(sortParams utils.SortParams) ([]models.NationalCase, error) {
	args := m.Called(sortParams)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

func (m *MockNationalCaseRepository) GetByDateRangeSorted(startDate, endDate time.Time, sortParams utils.SortParams) ([]models.NationalCase, error) {
	args := m.Called(startDate, endDate, sortParams)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}

type MockProvinceRepository struct {
	mock.Mock
}

func (m *MockProvinceRepository) GetAll() ([]models.Province, error) {
	args := m.Called()
	return args.Get(0).([]models.Province), args.Error(1)
}

func (m *MockProvinceRepository) GetByID(id string) (*models.Province, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.Province), args.Error(1)
}

type MockProvinceCaseRepository struct {
	mock.Mock
}

func (m *MockProvinceCaseRepository) GetAll() ([]models.ProvinceCaseWithDate, error) {
	args := m.Called()
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepository) GetByProvinceID(provinceID string) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepository) GetByProvinceIDAndDateRange(provinceID string, startDate, endDate time.Time) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID, startDate, endDate)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepository) GetByDateRange(startDate, endDate time.Time) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(startDate, endDate)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepository) GetLatestByProvinceID(provinceID string) (*models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.ProvinceCaseWithDate), args.Error(1)
}

// Paginated methods
func (m *MockProvinceCaseRepository) GetAllPaginated(limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockProvinceCaseRepository) GetByProvinceIDPaginated(provinceID string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(provinceID, limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockProvinceCaseRepository) GetByProvinceIDAndDateRangePaginated(provinceID string, startDate, endDate time.Time, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(provinceID, startDate, endDate, limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockProvinceCaseRepository) GetByDateRangePaginated(startDate, endDate time.Time, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(startDate, endDate, limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

// Sorted methods
func (m *MockProvinceCaseRepository) GetAllSorted(sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepository) GetByProvinceIDSorted(provinceID string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepository) GetByProvinceIDAndDateRangeSorted(provinceID string, startDate, endDate time.Time, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(provinceID, startDate, endDate, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

func (m *MockProvinceCaseRepository) GetByDateRangeSorted(startDate, endDate time.Time, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(startDate, endDate, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}

// Paginated sorted methods
func (m *MockProvinceCaseRepository) GetAllPaginatedSorted(limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(limit, offset, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockProvinceCaseRepository) GetByProvinceIDPaginatedSorted(provinceID string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(provinceID, limit, offset, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockProvinceCaseRepository) GetByProvinceIDAndDateRangePaginatedSorted(provinceID string, startDate, endDate time.Time, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(provinceID, startDate, endDate, limit, offset, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func (m *MockProvinceCaseRepository) GetByDateRangePaginatedSorted(startDate, endDate time.Time, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(startDate, endDate, limit, offset, sortParams)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func setupMockService() (*MockNationalCaseRepository, *MockProvinceRepository, *MockProvinceCaseRepository, CovidService) {
	mockNationalRepo := new(MockNationalCaseRepository)
	mockProvinceRepo := new(MockProvinceRepository)
	mockProvinceCaseRepo := new(MockProvinceCaseRepository)

	service := NewCovidService(mockNationalRepo, mockProvinceRepo, mockProvinceCaseRepo)

	return mockNationalRepo, mockProvinceRepo, mockProvinceCaseRepo, service
}

func TestCovidService_GetNationalCases(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()

	expectedCases := []models.NationalCase{
		{ID: 1, Positive: 100, Recovered: 80, Deceased: 5},
		{ID: 2, Positive: 150, Recovered: 120, Deceased: 8},
	}

	mockNationalRepo.On("GetAll").Return(expectedCases, nil)

	cases, err := service.GetNationalCases()

	assert.NoError(t, err)
	assert.Equal(t, expectedCases, cases)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetNationalCases_Error(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()

	mockNationalRepo.On("GetAll").Return([]models.NationalCase{}, errors.New("database error"))

	cases, err := service.GetNationalCases()

	assert.Error(t, err)
	assert.Nil(t, cases)
	assert.Contains(t, err.Error(), "failed to get national cases")
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetNationalCasesByDateRange(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()

	startDate := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expectedCases := []models.NationalCase{
		{ID: 1, Positive: 100, Date: startDate},
	}

	mockNationalRepo.On("GetByDateRange", startDate, endDate).Return(expectedCases, nil)

	cases, err := service.GetNationalCasesByDateRange("2020-03-01", "2020-03-31")

	assert.NoError(t, err)
	assert.Equal(t, expectedCases, cases)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetNationalCasesByDateRange_InvalidStartDate(t *testing.T) {
	_, _, _, service := setupMockService()

	cases, err := service.GetNationalCasesByDateRange("invalid-date", "2020-03-31")

	assert.Error(t, err)
	assert.Nil(t, cases)
	assert.Contains(t, err.Error(), "invalid start date format")
}

func TestCovidService_GetNationalCasesByDateRange_InvalidEndDate(t *testing.T) {
	_, _, _, service := setupMockService()

	cases, err := service.GetNationalCasesByDateRange("2020-03-01", "invalid-date")

	assert.Error(t, err)
	assert.Nil(t, cases)
	assert.Contains(t, err.Error(), "invalid end date format")
}

func TestCovidService_GetLatestNationalCase(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()

	expectedCase := &models.NationalCase{ID: 1, Positive: 100}
	mockNationalRepo.On("GetLatest").Return(expectedCase, nil)

	nationalCase, err := service.GetLatestNationalCase()

	assert.NoError(t, err)
	assert.Equal(t, expectedCase, nationalCase)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinces(t *testing.T) {
	_, mockProvinceRepo, _, service := setupMockService()

	expectedProvinces := []models.Province{
		{ID: "11", Name: "Aceh"},
		{ID: "31", Name: "DKI Jakarta"},
	}

	mockProvinceRepo.On("GetAll").Return(expectedProvinces, nil)

	provinces, err := service.GetProvinces()

	assert.NoError(t, err)
	assert.Equal(t, expectedProvinces, provinces)
	mockProvinceRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinceCases(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()

	provinceID := "11"
	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: provinceID, Positive: 50}},
	}

	mockProvinceCaseRepo.On("GetByProvinceID", provinceID).Return(expectedCases, nil)

	cases, err := service.GetProvinceCases(provinceID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCases, cases)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinceCasesByDateRange(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()

	provinceID := "11"
	startDate := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: provinceID, Positive: 50}},
	}

	mockProvinceCaseRepo.On("GetByProvinceIDAndDateRange", provinceID, startDate, endDate).Return(expectedCases, nil)

	cases, err := service.GetProvinceCasesByDateRange(provinceID, "2020-03-01", "2020-03-31")

	assert.NoError(t, err)
	assert.Equal(t, expectedCases, cases)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetAllProvinceCases(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()

	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}},
		{ProvinceCase: models.ProvinceCase{ID: 2, ProvinceID: "31", Positive: 100}},
	}

	mockProvinceCaseRepo.On("GetAll").Return(expectedCases, nil)

	cases, err := service.GetAllProvinceCases()

	assert.NoError(t, err)
	assert.Equal(t, expectedCases, cases)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetAllProvinceCasesByDateRange(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()

	startDate := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expectedCases := []models.ProvinceCaseWithDate{
		{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}},
	}

	mockProvinceCaseRepo.On("GetByDateRange", startDate, endDate).Return(expectedCases, nil)

	cases, err := service.GetAllProvinceCasesByDateRange("2020-03-01", "2020-03-31")

	assert.NoError(t, err)
	assert.Equal(t, expectedCases, cases)
	mockProvinceCaseRepo.AssertExpectations(t)
}
