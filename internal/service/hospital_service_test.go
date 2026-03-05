package service

import (
	"errors"
	"testing"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockHospitalRepository mocks repository.HospitalRepositoryInterface
type MockHospitalRepository struct {
	mock.Mock
}

func (m *MockHospitalRepository) GetAll(provinceID int) ([]models.Hospital, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.Hospital), args.Error(1)
}

func (m *MockHospitalRepository) GetByCode(code string) (*models.Hospital, error) {
	args := m.Called(code)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.Hospital), args.Error(1)
}

func setupHospitalService() (*MockHospitalRepository, *HospitalService) {
	mockRepo := new(MockHospitalRepository)
	svc := NewHospitalService(mockRepo)
	return mockRepo, svc
}

func TestHospitalService_GetHospitals(t *testing.T) {
	mockRepo, svc := setupHospitalService()

	code := "7201001"
	expected := []models.Hospital{
		{ID: 1, RegencyID: 7201, Name: "RSUD Banggai", HospitalCode: &code},
		{ID: 2, RegencyID: 7202, Name: "RSUD Palu"},
	}
	mockRepo.On("GetAll", 72).Return(expected, nil)

	result, err := svc.GetHospitals()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestHospitalService_GetHospitals_Error(t *testing.T) {
	mockRepo, svc := setupHospitalService()

	mockRepo.On("GetAll", 72).Return([]models.Hospital{}, errors.New("db error"))

	result, err := svc.GetHospitals()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestHospitalService_GetHospitalByCode(t *testing.T) {
	mockRepo, svc := setupHospitalService()

	code := "7201001"
	expected := &models.Hospital{ID: 1, RegencyID: 7201, Name: "RSUD Banggai", HospitalCode: &code}
	mockRepo.On("GetByCode", code).Return(expected, nil)

	result, err := svc.GetHospitalByCode(code)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestHospitalService_GetHospitalByCode_NotFound(t *testing.T) {
	mockRepo, svc := setupHospitalService()

	mockRepo.On("GetByCode", "NOTFOUND").Return(nil, nil)

	result, err := svc.GetHospitalByCode("NOTFOUND")

	assert.NoError(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestHospitalService_GetHospitalByCode_Error(t *testing.T) {
	mockRepo, svc := setupHospitalService()

	mockRepo.On("GetByCode", "7201001").Return(nil, errors.New("db error"))

	result, err := svc.GetHospitalByCode("7201001")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
