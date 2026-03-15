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

type MockTaskForceService struct{ mock.Mock }

func (m *MockTaskForceService) GetTaskForces() ([]models.TaskForceByRegency, error) {
	args := m.Called()
	return args.Get(0).([]models.TaskForceByRegency), args.Error(1)
}

func TestGetTaskForces_Success(t *testing.T) {
	svc := new(MockTaskForceService)
	data := []models.TaskForceByRegency{{RegencyID: 7201, RegencyName: "Kab. Banggai"}}
	svc.On("GetTaskForces").Return(data, nil)

	h := NewTaskForceHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/task-forces", nil)
	w := httptest.NewRecorder()
	h.GetTaskForces(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetTaskForces_Error(t *testing.T) {
	svc := new(MockTaskForceService)
	svc.On("GetTaskForces").Return([]models.TaskForceByRegency{}, errors.New("db error"))

	h := NewTaskForceHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/task-forces", nil)
	w := httptest.NewRecorder()
	h.GetTaskForces(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}
