package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProvinceStatsService struct{ mock.Mock }

func (m *MockProvinceStatsService) GetGenderCases() ([]models.ProvinceGenderCase, error) {
	args := m.Called()
	return args.Get(0).([]models.ProvinceGenderCase), args.Error(1)
}
func (m *MockProvinceStatsService) GetLatestGenderCase() (*models.ProvinceGenderCase, error) {
	args := m.Called()
	if r := args.Get(0); r != nil {
		return r.(*models.ProvinceGenderCase), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockProvinceStatsService) GetTests() ([]models.ProvinceTest, error) {
	args := m.Called()
	return args.Get(0).([]models.ProvinceTest), args.Error(1)
}
func (m *MockProvinceStatsService) GetTestTypes() ([]models.TestType, error) {
	args := m.Called()
	return args.Get(0).([]models.TestType), args.Error(1)
}

func TestGetGenderCases_Success(t *testing.T) {
	svc := new(MockProvinceStatsService)
	data := []models.ProvinceGenderCase{{ID: 1, ProvinceID: 72}}
	svc.On("GetGenderCases").Return(data, nil)

	h := NewProvinceStatsHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/gender", nil)
	w := httptest.NewRecorder()
	h.GetGenderCases(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetGenderCases_Error(t *testing.T) {
	svc := new(MockProvinceStatsService)
	svc.On("GetGenderCases").Return([]models.ProvinceGenderCase{}, errors.New("db error"))

	h := NewProvinceStatsHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/gender", nil)
	w := httptest.NewRecorder()
	h.GetGenderCases(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetLatestGenderCase_Success(t *testing.T) {
	svc := new(MockProvinceStatsService)
	data := &models.ProvinceGenderCase{ID: 1, ProvinceID: 72}
	svc.On("GetLatestGenderCase").Return(data, nil)

	h := NewProvinceStatsHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/gender/latest", nil)
	w := httptest.NewRecorder()
	h.GetLatestGenderCase(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetTests_Success(t *testing.T) {
	svc := new(MockProvinceStatsService)
	svc.On("GetTests").Return([]models.ProvinceTest{{ID: 1}}, nil)

	h := NewProvinceStatsHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/tests", nil)
	w := httptest.NewRecorder()
	h.GetTests(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetTestTypes_Success(t *testing.T) {
	svc := new(MockProvinceStatsService)
	svc.On("GetTestTypes").Return([]models.TestType{{ID: 1, Key: "pcr", Name: "PCR"}}, nil)

	h := NewProvinceStatsHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats/test-types", nil)
	w := httptest.NewRecorder()
	h.GetTestTypes(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}
