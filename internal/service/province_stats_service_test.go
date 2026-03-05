package service

import (
	"errors"
	"testing"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProvinceStatsRepository mocks repository.ProvinceStatsRepositoryInterface
type MockProvinceStatsRepository struct {
	mock.Mock
}

func (m *MockProvinceStatsRepository) GetGenderCases(provinceID int) ([]models.ProvinceGenderCase, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.ProvinceGenderCase), args.Error(1)
}

func (m *MockProvinceStatsRepository) GetLatestGenderCase(provinceID int) (*models.ProvinceGenderCase, error) {
	args := m.Called(provinceID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.ProvinceGenderCase), args.Error(1)
}

func (m *MockProvinceStatsRepository) GetTests(provinceID int) ([]models.ProvinceTest, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.ProvinceTest), args.Error(1)
}

func (m *MockProvinceStatsRepository) GetTestTypes() ([]models.TestType, error) {
	args := m.Called()
	return args.Get(0).([]models.TestType), args.Error(1)
}

func setupProvinceStatsService() (*MockProvinceStatsRepository, *ProvinceStatsService) {
	mockRepo := new(MockProvinceStatsRepository)
	svc := NewProvinceStatsService(mockRepo)
	return mockRepo, svc
}

func TestProvinceStatsService_GetGenderCases(t *testing.T) {
	mockRepo, svc := setupProvinceStatsService()

	expected := []models.ProvinceGenderCase{
		{ID: 1, Day: 1, ProvinceID: 72, PositiveMale: 50, PositiveFemale: 40},
		{ID: 2, Day: 2, ProvinceID: 72, PositiveMale: 60, PositiveFemale: 55},
	}
	mockRepo.On("GetGenderCases", 72).Return(expected, nil)

	result, err := svc.GetGenderCases()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestProvinceStatsService_GetGenderCases_Error(t *testing.T) {
	mockRepo, svc := setupProvinceStatsService()

	mockRepo.On("GetGenderCases", 72).Return([]models.ProvinceGenderCase{}, errors.New("db error"))

	result, err := svc.GetGenderCases()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestProvinceStatsService_GetLatestGenderCase(t *testing.T) {
	mockRepo, svc := setupProvinceStatsService()

	expected := &models.ProvinceGenderCase{ID: 2, Day: 2, ProvinceID: 72, PositiveMale: 60, PositiveFemale: 55}
	mockRepo.On("GetLatestGenderCase", 72).Return(expected, nil)

	result, err := svc.GetLatestGenderCase()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestProvinceStatsService_GetLatestGenderCase_NotFound(t *testing.T) {
	mockRepo, svc := setupProvinceStatsService()

	mockRepo.On("GetLatestGenderCase", 72).Return(nil, nil)

	result, err := svc.GetLatestGenderCase()

	assert.NoError(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestProvinceStatsService_GetLatestGenderCase_Error(t *testing.T) {
	mockRepo, svc := setupProvinceStatsService()

	mockRepo.On("GetLatestGenderCase", 72).Return(nil, errors.New("db error"))

	result, err := svc.GetLatestGenderCase()

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestProvinceStatsService_GetTests(t *testing.T) {
	mockRepo, svc := setupProvinceStatsService()

	expected := []models.ProvinceTest{
		{ID: 1, ProvinceID: 72, TestTypeID: 1, Positive: 100},
		{ID: 2, ProvinceID: 72, TestTypeID: 2, Positive: 50},
	}
	mockRepo.On("GetTests", 72).Return(expected, nil)

	result, err := svc.GetTests()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestProvinceStatsService_GetTests_Error(t *testing.T) {
	mockRepo, svc := setupProvinceStatsService()

	mockRepo.On("GetTests", 72).Return([]models.ProvinceTest{}, errors.New("db error"))

	result, err := svc.GetTests()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestProvinceStatsService_GetTestTypes(t *testing.T) {
	mockRepo, svc := setupProvinceStatsService()

	expected := []models.TestType{
		{ID: 1, Name: "PCR"},
		{ID: 2, Name: "Antigen"},
	}
	mockRepo.On("GetTestTypes").Return(expected, nil)

	result, err := svc.GetTestTypes()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestProvinceStatsService_GetTestTypes_Error(t *testing.T) {
	mockRepo, svc := setupProvinceStatsService()

	mockRepo.On("GetTestTypes").Return([]models.TestType{}, errors.New("db error"))

	result, err := svc.GetTestTypes()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}
