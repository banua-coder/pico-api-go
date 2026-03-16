package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVaccinationService struct{ mock.Mock }

func (m *MockVaccinationService) GetNationalVaccinationsPaginated(limit, offset int) ([]models.NationalVaccine, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.NationalVaccine), args.Int(1), args.Error(2)
}

func (m *MockVaccinationService) GetProvinceVaccinationsPaginated(limit, offset int) ([]models.ProvinceVaccine, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.ProvinceVaccine), args.Int(1), args.Error(2)
}

func (m *MockVaccinationService) GetVaccineLocationsPaginated(limit, offset int) ([]models.VaccineLocation, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.VaccineLocation), args.Int(1), args.Error(2)
}

func (m *MockVaccinationService) GetNationalVaccinations() ([]models.NationalVaccine, error) {
	args := m.Called()
	return args.Get(0).([]models.NationalVaccine), args.Error(1)
}
func (m *MockVaccinationService) GetProvinceVaccinations() ([]models.ProvinceVaccine, error) {
	args := m.Called()
	return args.Get(0).([]models.ProvinceVaccine), args.Error(1)
}
func (m *MockVaccinationService) GetVaccineLocations() ([]models.VaccineLocation, error) {
	args := m.Called()
	return args.Get(0).([]models.VaccineLocation), args.Error(1)
}

func sampleNationalVaccine() models.NationalVaccine {
	return models.NationalVaccine{ID: 1, Day: 1, Date: time.Now()}
}

func TestGetNationalVaccinations_Success(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetNationalVaccinationsPaginated", 10, 0).Return([]models.NationalVaccine{sampleNationalVaccine()}, 1, nil)

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/national", nil)
	w := httptest.NewRecorder()
	h.GetNationalVaccinations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetNationalVaccinations_Error(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetNationalVaccinationsPaginated", 10, 0).Return([]models.NationalVaccine{}, 0, errors.New("db error"))

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/national", nil)
	w := httptest.NewRecorder()
	h.GetNationalVaccinations(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetProvinceVaccinations_Success(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetProvinceVaccinationsPaginated", 10, 0).Return([]models.ProvinceVaccine{{NationalVaccine: sampleNationalVaccine(), ProvinceID: 72}}, 1, nil)

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/province", nil)
	w := httptest.NewRecorder()
	h.GetProvinceVaccinations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetVaccineLocations_Success(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetVaccineLocationsPaginated", 10, 0).Return([]models.VaccineLocation{{ID: 1, Name: "Puskesmas A"}}, 1, nil)

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/locations", nil)
	w := httptest.NewRecorder()
	h.GetVaccineLocations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetNationalVaccinations_LoadAll(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetNationalVaccinations").Return([]models.NationalVaccine{sampleNationalVaccine()}, nil)

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/national?load_all=true", nil)
	w := httptest.NewRecorder()
	h.GetNationalVaccinations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetNationalVaccinations_LoadAll_Error(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetNationalVaccinations").Return([]models.NationalVaccine{}, errors.New("db error"))

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/national?load_all=true", nil)
	w := httptest.NewRecorder()
	h.GetNationalVaccinations(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetProvinceVaccinations_LoadAll(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetProvinceVaccinations").Return([]models.ProvinceVaccine{{NationalVaccine: sampleNationalVaccine(), ProvinceID: 72}}, nil)

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/province?load_all=true", nil)
	w := httptest.NewRecorder()
	h.GetProvinceVaccinations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetProvinceVaccinations_LoadAll_Error(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetProvinceVaccinations").Return([]models.ProvinceVaccine{}, errors.New("db error"))

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/province?load_all=true", nil)
	w := httptest.NewRecorder()
	h.GetProvinceVaccinations(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetProvinceVaccinations_Error(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetProvinceVaccinationsPaginated", 10, 0).Return([]models.ProvinceVaccine{}, 0, errors.New("db error"))

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/province", nil)
	w := httptest.NewRecorder()
	h.GetProvinceVaccinations(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetVaccineLocations_LoadAll(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetVaccineLocations").Return([]models.VaccineLocation{{ID: 1, Name: "Puskesmas A"}}, nil)

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/locations?load_all=true", nil)
	w := httptest.NewRecorder()
	h.GetVaccineLocations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetVaccineLocations_LoadAll_Error(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetVaccineLocations").Return([]models.VaccineLocation{}, errors.New("db error"))

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/locations?load_all=true", nil)
	w := httptest.NewRecorder()
	h.GetVaccineLocations(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetVaccineLocations_Error(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetVaccineLocationsPaginated", 10, 0).Return([]models.VaccineLocation{}, 0, errors.New("db error"))

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/locations", nil)
	w := httptest.NewRecorder()
	h.GetVaccineLocations(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}
