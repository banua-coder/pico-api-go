package service

import (
	"errors"
	"testing"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRegencyRepository mocks repository.RegencyRepositoryInterface
type MockRegencyRepository struct {
	mock.Mock
}

func (m *MockRegencyRepository) GetAll(provinceID int) ([]models.Regency, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.Regency), args.Error(1)
}

func (m *MockRegencyRepository) GetByID(id int) (*models.Regency, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.Regency), args.Error(1)
}

// MockRegencyCaseRepository mocks repository.RegencyCaseRepositoryInterface
type MockRegencyCaseRepository struct {
	mock.Mock
}

func (m *MockRegencyCaseRepository) GetByRegencyID(regencyID int) ([]models.RegencyCase, error) {
	args := m.Called(regencyID)
	return args.Get(0).([]models.RegencyCase), args.Error(1)
}

func (m *MockRegencyCaseRepository) GetLatestByProvinceID(provinceID int) ([]models.RegencyCase, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.RegencyCase), args.Error(1)
}

func setupRegencyService() (*MockRegencyRepository, *MockRegencyCaseRepository, *RegencyService) {
	mockRepo := new(MockRegencyRepository)
	mockCaseRepo := new(MockRegencyCaseRepository)
	svc := NewRegencyService(mockRepo, mockCaseRepo)
	return mockRepo, mockCaseRepo, svc
}

func TestRegencyService_GetRegencies(t *testing.T) {
	mockRepo, _, svc := setupRegencyService()

	expected := []models.Regency{
		{ID: 7201, ProvinceID: 72, Name: "Kabupaten Banggai"},
		{ID: 7202, ProvinceID: 72, Name: "Kabupaten Banggai Kepulauan"},
	}
	mockRepo.On("GetAll", 72).Return(expected, nil)

	result, err := svc.GetRegencies()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestRegencyService_GetRegencies_Error(t *testing.T) {
	mockRepo, _, svc := setupRegencyService()

	mockRepo.On("GetAll", 72).Return([]models.Regency{}, errors.New("db error"))

	result, err := svc.GetRegencies()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestRegencyService_GetRegencyByID(t *testing.T) {
	mockRepo, _, svc := setupRegencyService()

	expected := &models.Regency{ID: 7201, ProvinceID: 72, Name: "Kabupaten Banggai"}
	mockRepo.On("GetByID", 7201).Return(expected, nil)

	result, err := svc.GetRegencyByID(7201)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestRegencyService_GetRegencyByID_NotFound(t *testing.T) {
	mockRepo, _, svc := setupRegencyService()

	mockRepo.On("GetByID", 9999).Return(nil, nil)

	result, err := svc.GetRegencyByID(9999)

	assert.NoError(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestRegencyService_GetRegencyByID_Error(t *testing.T) {
	mockRepo, _, svc := setupRegencyService()

	mockRepo.On("GetByID", 7201).Return(nil, errors.New("db error"))

	result, err := svc.GetRegencyByID(7201)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestRegencyService_GetRegencyCases(t *testing.T) {
	_, mockCaseRepo, svc := setupRegencyService()

	expected := []models.RegencyCase{
		{ID: 1, RegencyID: 7201, Positive: 10, Recovered: 8, Deceased: 1},
	}
	mockCaseRepo.On("GetByRegencyID", 7201).Return(expected, nil)

	result, err := svc.GetRegencyCases(7201)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockCaseRepo.AssertExpectations(t)
}

func TestRegencyService_GetRegencyCases_Error(t *testing.T) {
	_, mockCaseRepo, svc := setupRegencyService()

	mockCaseRepo.On("GetByRegencyID", 7201).Return([]models.RegencyCase{}, errors.New("db error"))

	result, err := svc.GetRegencyCases(7201)

	assert.Error(t, err)
	assert.Empty(t, result)
	mockCaseRepo.AssertExpectations(t)
}

func TestRegencyService_GetLatestRegencyCases(t *testing.T) {
	_, mockCaseRepo, svc := setupRegencyService()

	expected := []models.RegencyCase{
		{ID: 1, RegencyID: 7201, Positive: 10},
		{ID: 2, RegencyID: 7202, Positive: 5},
	}
	mockCaseRepo.On("GetLatestByProvinceID", 72).Return(expected, nil)

	result, err := svc.GetLatestRegencyCases()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockCaseRepo.AssertExpectations(t)
}

func TestRegencyService_GetLatestRegencyCases_Error(t *testing.T) {
	_, mockCaseRepo, svc := setupRegencyService()

	mockCaseRepo.On("GetLatestByProvinceID", 72).Return([]models.RegencyCase{}, errors.New("db error"))

	result, err := svc.GetLatestRegencyCases()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockCaseRepo.AssertExpectations(t)
}
