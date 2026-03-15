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
	svc.On("GetNationalVaccinations").Return([]models.NationalVaccine{sampleNationalVaccine()}, nil)

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/national", nil)
	w := httptest.NewRecorder()
	h.GetNationalVaccinations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetNationalVaccinations_Error(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetNationalVaccinations").Return([]models.NationalVaccine{}, errors.New("db error"))

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/national", nil)
	w := httptest.NewRecorder()
	h.GetNationalVaccinations(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetProvinceVaccinations_Success(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetProvinceVaccinations").Return([]models.ProvinceVaccine{{NationalVaccine: sampleNationalVaccine(), ProvinceID: 72}}, nil)

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/province", nil)
	w := httptest.NewRecorder()
	h.GetProvinceVaccinations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetVaccineLocations_Success(t *testing.T) {
	svc := new(MockVaccinationService)
	svc.On("GetVaccineLocations").Return([]models.VaccineLocation{{ID: 1, Name: "Puskesmas A"}}, nil)

	h := NewVaccinationHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/vaccination/locations", nil)
	w := httptest.NewRecorder()
	h.GetVaccineLocations(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}
