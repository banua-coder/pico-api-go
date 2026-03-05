package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHospitalService struct{ mock.Mock }

func (m *MockHospitalService) GetHospitals() ([]models.Hospital, error) {
	args := m.Called()
	return args.Get(0).([]models.Hospital), args.Error(1)
}
func (m *MockHospitalService) GetHospitalByCode(code string) (*models.Hospital, error) {
	args := m.Called(code)
	if r := args.Get(0); r != nil {
		return r.(*models.Hospital), args.Error(1)
	}
	return nil, args.Error(1)
}

func strPtr(s string) *string { return &s }

func TestGetHospitals_Success(t *testing.T) {
	svc := new(MockHospitalService)
	hospitals := []models.Hospital{{ID: 1, Name: "RSUD Undata", HospitalCode: strPtr("7200001")}}
	svc.On("GetHospitals").Return(hospitals, nil)

	h := NewHospitalHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hospitals", nil)
	w := httptest.NewRecorder()
	h.GetHospitals(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetHospitals_Error(t *testing.T) {
	svc := new(MockHospitalService)
	svc.On("GetHospitals").Return([]models.Hospital{}, errors.New("db error"))

	h := NewHospitalHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hospitals", nil)
	w := httptest.NewRecorder()
	h.GetHospitals(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetHospitalByCode_Success(t *testing.T) {
	svc := new(MockHospitalService)
	hospital := &models.Hospital{ID: 1, Name: "RSUD Undata", HospitalCode: strPtr("7200001")}
	svc.On("GetHospitalByCode", "7200001").Return(hospital, nil)

	h := NewHospitalHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hospitals/7200001", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/hospitals/{code}", h.GetHospitalByCode)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetHospitalByCode_NotFound(t *testing.T) {
	svc := new(MockHospitalService)
	svc.On("GetHospitalByCode", "UNKNOWN").Return(nil, nil)

	h := NewHospitalHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hospitals/UNKNOWN", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/hospitals/{code}", h.GetHospitalByCode)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	svc.AssertExpectations(t)
}

func TestGetHospitalByCode_Error(t *testing.T) {
	svc := new(MockHospitalService)
	svc.On("GetHospitalByCode", "7200001").Return(nil, errors.New("db error"))

	h := NewHospitalHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hospitals/7200001", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/hospitals/{code}", h.GetHospitalByCode)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}
