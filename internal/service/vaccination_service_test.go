package service

import (
	"errors"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockVaccinationRepository mocks repository.VaccinationRepositoryInterface
type MockVaccinationRepository struct {
	mock.Mock
}

func (m *MockVaccinationRepository) GetNationalVaccinations() ([]models.NationalVaccine, error) {
	args := m.Called()
	return args.Get(0).([]models.NationalVaccine), args.Error(1)
}

func (m *MockVaccinationRepository) GetNationalVaccinationsPaginated(limit, offset int) ([]models.NationalVaccine, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.NationalVaccine), args.Int(1), args.Error(2)
}

func (m *MockVaccinationRepository) GetProvinceVaccinationsPaginated(provinceID, limit, offset int) ([]models.ProvinceVaccine, int, error) {
	args := m.Called(provinceID, limit, offset)
	return args.Get(0).([]models.ProvinceVaccine), args.Int(1), args.Error(2)
}

func (m *MockVaccinationRepository) GetVaccineLocationsPaginated(provinceID, limit, offset int) ([]models.VaccineLocation, int, error) {
	args := m.Called(provinceID, limit, offset)
	return args.Get(0).([]models.VaccineLocation), args.Int(1), args.Error(2)
}

func (m *MockVaccinationRepository) GetProvinceVaccinations(provinceID int) ([]models.ProvinceVaccine, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.ProvinceVaccine), args.Error(1)
}

func (m *MockVaccinationRepository) GetVaccineLocations(provinceID int) ([]models.VaccineLocation, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.VaccineLocation), args.Error(1)
}

func setupVaccinationService() (*MockVaccinationRepository, *VaccinationService) {
	mockRepo := new(MockVaccinationRepository)
	svc := NewVaccinationService(mockRepo)
	return mockRepo, svc
}

func TestVaccinationService_GetNationalVaccinations(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	expected := []models.NationalVaccine{
		{ID: 1, Day: 1, Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC), TotalVaccinationTarget: 1000000},
		{ID: 2, Day: 2, Date: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC), TotalVaccinationTarget: 1000000},
	}
	mockRepo.On("GetNationalVaccinations").Return(expected, nil)

	result, err := svc.GetNationalVaccinations()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetNationalVaccinations_Error(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	mockRepo.On("GetNationalVaccinations").Return([]models.NationalVaccine{}, errors.New("db error"))

	result, err := svc.GetNationalVaccinations()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetProvinceVaccinations(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	expected := []models.ProvinceVaccine{
		{ProvinceID: 72, NationalVaccine: models.NationalVaccine{ID: 1, TotalVaccinationTarget: 500000}},
	}
	mockRepo.On("GetProvinceVaccinations", 72).Return(expected, nil)

	result, err := svc.GetProvinceVaccinations()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetProvinceVaccinations_Error(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	mockRepo.On("GetProvinceVaccinations", 72).Return([]models.ProvinceVaccine{}, errors.New("db error"))

	result, err := svc.GetProvinceVaccinations()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetVaccineLocations(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	expected := []models.VaccineLocation{
		{ID: 1, RegencyID: 7201, Name: "Puskesmas Banggai", Address: "Jl. Merdeka"},
		{ID: 2, RegencyID: 7202, Name: "Puskesmas Palu", Address: "Jl. Veteran"},
	}
	mockRepo.On("GetVaccineLocations", 72).Return(expected, nil)

	result, err := svc.GetVaccineLocations()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetVaccineLocations_Error(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	mockRepo.On("GetVaccineLocations", 72).Return([]models.VaccineLocation{}, errors.New("db error"))

	result, err := svc.GetVaccineLocations()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetNationalVaccinationsPaginated(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	expected := []models.NationalVaccine{
		{ID: 1, Day: 1, Date: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)},
	}
	mockRepo.On("GetNationalVaccinationsPaginated", 10, 0).Return(expected, 1, nil)

	result, total, err := svc.GetNationalVaccinationsPaginated(10, 0)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetNationalVaccinationsPaginated_Error(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	mockRepo.On("GetNationalVaccinationsPaginated", 10, 0).Return([]models.NationalVaccine{}, 0, errors.New("db error"))

	result, total, err := svc.GetNationalVaccinationsPaginated(10, 0)

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, 0, total)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetProvinceVaccinationsPaginated(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	expected := []models.ProvinceVaccine{
		{ProvinceID: 72, NationalVaccine: models.NationalVaccine{ID: 1}},
	}
	mockRepo.On("GetProvinceVaccinationsPaginated", 72, 10, 0).Return(expected, 1, nil)

	result, total, err := svc.GetProvinceVaccinationsPaginated(10, 0)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetProvinceVaccinationsPaginated_Error(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	mockRepo.On("GetProvinceVaccinationsPaginated", 72, 10, 0).Return([]models.ProvinceVaccine{}, 0, errors.New("db error"))

	result, total, err := svc.GetProvinceVaccinationsPaginated(10, 0)

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, 0, total)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetVaccineLocationsPaginated(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	expected := []models.VaccineLocation{
		{ID: 1, RegencyID: 7201, Name: "Puskesmas Banggai"},
	}
	mockRepo.On("GetVaccineLocationsPaginated", 72, 10, 0).Return(expected, 1, nil)

	result, total, err := svc.GetVaccineLocationsPaginated(10, 0)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockRepo.AssertExpectations(t)
}

func TestVaccinationService_GetVaccineLocationsPaginated_Error(t *testing.T) {
	mockRepo, svc := setupVaccinationService()

	mockRepo.On("GetVaccineLocationsPaginated", 72, 10, 0).Return([]models.VaccineLocation{}, 0, errors.New("db error"))

	result, total, err := svc.GetVaccineLocationsPaginated(10, 0)

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, 0, total)
	mockRepo.AssertExpectations(t)
}
