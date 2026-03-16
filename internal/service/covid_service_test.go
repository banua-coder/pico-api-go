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

func (m *MockNationalCaseRepository) GetAllPaginated(limit, offset int) ([]models.NationalCase, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.NationalCase), args.Int(1), args.Error(2)
}

func (m *MockNationalCaseRepository) GetAllPaginatedSorted(limit, offset int, sortParams utils.SortParams) ([]models.NationalCase, int, error) {
	args := m.Called(limit, offset, sortParams)
	return args.Get(0).([]models.NationalCase), args.Int(1), args.Error(2)
}

func (m *MockNationalCaseRepository) GetByDateRangePaginated(startDate, endDate time.Time, limit, offset int) ([]models.NationalCase, int, error) {
	args := m.Called(startDate, endDate, limit, offset)
	return args.Get(0).([]models.NationalCase), args.Int(1), args.Error(2)
}

func (m *MockNationalCaseRepository) GetByDateRangePaginatedSorted(startDate, endDate time.Time, limit, offset int, sortParams utils.SortParams) ([]models.NationalCase, int, error) {
	args := m.Called(startDate, endDate, limit, offset, sortParams)
	return args.Get(0).([]models.NationalCase), args.Int(1), args.Error(2)
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

func TestCovidService_GetNationalCasesSorted(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	expected := []models.NationalCase{{ID: 1, Positive: 100}}
	mockNationalRepo.On("GetAllSorted", sort).Return(expected, nil)
	result, err := service.GetNationalCasesSorted(sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetNationalCasesByDateRangeSorted(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expected := []models.NationalCase{{ID: 1, Positive: 100}}
	mockNationalRepo.On("GetByDateRangeSorted", start, end, sort).Return(expected, nil)
	result, err := service.GetNationalCasesByDateRangeSorted("2020-03-01", "2020-03-31", sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetNationalCaseByDay(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()
	expected := &models.NationalCase{ID: 1, Positive: 100}
	mockNationalRepo.On("GetByDay", int64(1)).Return(expected, nil)
	result, err := service.GetNationalCaseByDay(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinceByID(t *testing.T) {
	_, mockProvinceRepo, _, service := setupMockService()
	expected := &models.Province{ID: "11", Name: "Aceh"}
	mockProvinceRepo.On("GetByID", "11").Return(expected, nil)
	result, err := service.GetProvinceByID("11")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockProvinceRepo.AssertExpectations(t)
}

func TestCovidService_GetNationalCasesPaginated(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()
	expected := []models.NationalCase{{ID: 1, Positive: 100}}
	mockNationalRepo.On("GetAllPaginated", 10, 0).Return(expected, 1, nil)
	result, total, err := service.GetNationalCasesPaginated(10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetNationalCasesPaginatedSorted(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	expected := []models.NationalCase{{ID: 1, Positive: 100}}
	mockNationalRepo.On("GetAllPaginatedSorted", 10, 0, sort).Return(expected, 1, nil)
	result, total, err := service.GetNationalCasesPaginatedSorted(10, 0, sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetNationalCasesByDateRangePaginated(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expected := []models.NationalCase{{ID: 1, Positive: 100}}
	mockNationalRepo.On("GetByDateRangePaginated", start, end, 10, 0).Return(expected, 1, nil)
	result, total, err := service.GetNationalCasesByDateRangePaginated("2020-03-01", "2020-03-31", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetNationalCasesByDateRangePaginatedSorted(t *testing.T) {
	mockNationalRepo, _, _, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expected := []models.NationalCase{{ID: 1, Positive: 100}}
	mockNationalRepo.On("GetByDateRangePaginatedSorted", start, end, 10, 0, sort).Return(expected, 1, nil)
	result, total, err := service.GetNationalCasesByDateRangePaginatedSorted("2020-03-01", "2020-03-31", 10, 0, sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockNationalRepo.AssertExpectations(t)
}

func TestCovidService_GetProvincesWithLatestCase(t *testing.T) {
	_, mockProvinceRepo, mockProvinceCaseRepo, service := setupMockService()
	provinces := []models.Province{{ID: "11", Name: "Aceh"}}
	latestCase := &models.ProvinceCaseWithDate{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11", Positive: 50}}
	mockProvinceRepo.On("GetAll").Return(provinces, nil)
	mockProvinceCaseRepo.On("GetLatestByProvinceID", "11").Return(latestCase, nil)
	result, err := service.GetProvincesWithLatestCase()
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	mockProvinceRepo.AssertExpectations(t)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetAllProvinceCasesSorted(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1}}}
	mockProvinceCaseRepo.On("GetAllSorted", sort).Return(expected, nil)
	result, err := service.GetAllProvinceCasesSorted(sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinceCasesPaginated(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11"}}}
	mockProvinceCaseRepo.On("GetByProvinceIDPaginated", "11", 10, 0).Return(expected, 1, nil)
	result, total, err := service.GetProvinceCasesPaginated("11", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinceCasesByDateRangePaginated(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1}}}
	mockProvinceCaseRepo.On("GetByProvinceIDAndDateRangePaginated", "11", start, end, 10, 0).Return(expected, 1, nil)
	result, total, err := service.GetProvinceCasesByDateRangePaginated("11", "2020-03-01", "2020-03-31", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetAllProvinceCasesPaginated(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1}}}
	mockProvinceCaseRepo.On("GetAllPaginated", 10, 0).Return(expected, 1, nil)
	result, total, err := service.GetAllProvinceCasesPaginated(10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetAllProvinceCasesByDateRangePaginated(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1}}}
	mockProvinceCaseRepo.On("GetByDateRangePaginated", start, end, 10, 0).Return(expected, 1, nil)
	result, total, err := service.GetAllProvinceCasesByDateRangePaginated("2020-03-01", "2020-03-31", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetAllProvinceCasesPaginatedSorted(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1}}}
	mockProvinceCaseRepo.On("GetAllPaginatedSorted", 10, 0, sort).Return(expected, 1, nil)
	result, total, err := service.GetAllProvinceCasesPaginatedSorted(10, 0, sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetAllProvinceCasesByDateRangeSorted(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1}}}
	mockProvinceCaseRepo.On("GetByDateRangeSorted", start, end, sort).Return(expected, nil)
	result, err := service.GetAllProvinceCasesByDateRangeSorted("2020-03-01", "2020-03-31", sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetAllProvinceCasesByDateRangePaginatedSorted(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1}}}
	mockProvinceCaseRepo.On("GetByDateRangePaginatedSorted", start, end, 10, 0, sort).Return(expected, 1, nil)
	result, total, err := service.GetAllProvinceCasesByDateRangePaginatedSorted("2020-03-01", "2020-03-31", 10, 0, sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinceCasesSorted(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11"}}}
	mockProvinceCaseRepo.On("GetByProvinceIDSorted", "11", sort).Return(expected, nil)
	result, err := service.GetProvinceCasesSorted("11", sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinceCasesPaginatedSorted(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1, ProvinceID: "11"}}}
	mockProvinceCaseRepo.On("GetByProvinceIDPaginatedSorted", "11", 10, 0, sort).Return(expected, 1, nil)
	result, total, err := service.GetProvinceCasesPaginatedSorted("11", 10, 0, sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinceCasesByDateRangeSorted(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1}}}
	mockProvinceCaseRepo.On("GetByProvinceIDAndDateRangeSorted", "11", start, end, sort).Return(expected, nil)
	result, err := service.GetProvinceCasesByDateRangeSorted("11", "2020-03-01", "2020-03-31", sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockProvinceCaseRepo.AssertExpectations(t)
}

func TestCovidService_GetProvinceCasesByDateRangePaginatedSorted(t *testing.T) {
	_, _, mockProvinceCaseRepo, service := setupMockService()
	sort := utils.SortParams{Field: "day", Order: "asc"}
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	expected := []models.ProvinceCaseWithDate{{ProvinceCase: models.ProvinceCase{ID: 1}}}
	mockProvinceCaseRepo.On("GetByProvinceIDAndDateRangePaginatedSorted", "11", start, end, 10, 0, sort).Return(expected, 1, nil)
	result, total, err := service.GetProvinceCasesByDateRangePaginatedSorted("11", "2020-03-01", "2020-03-31", 10, 0, sort)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockProvinceCaseRepo.AssertExpectations(t)
}
