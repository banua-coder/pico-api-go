package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRegencyService struct{ mock.Mock }

func (m *MockRegencyService) GetRegenciesPaginated(limit, offset int) ([]models.Regency, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.Regency), args.Int(1), args.Error(2)
}

func (m *MockRegencyService) GetRegencies() ([]models.Regency, error) {
	args := m.Called()
	return args.Get(0).([]models.Regency), args.Error(1)
}
func (m *MockRegencyService) GetRegencyByID(id int) (*models.Regency, error) {
	args := m.Called(id)
	if r := args.Get(0); r != nil {
		return r.(*models.Regency), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *MockRegencyService) GetRegencyCases(id int) ([]models.RegencyCase, error) {
	args := m.Called(id)
	return args.Get(0).([]models.RegencyCase), args.Error(1)
}
func (m *MockRegencyService) GetLatestRegencyCases() ([]models.RegencyCase, error) {
	args := m.Called()
	return args.Get(0).([]models.RegencyCase), args.Error(1)
}

func TestGetRegencies_Success(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegenciesPaginated", 10, 0).Return([]models.Regency{{ID: 7201, Name: "Kab. Banggai"}}, 1, nil)

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies", nil)
	w := httptest.NewRecorder()
	h.GetRegencies(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetRegencies_Error(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegenciesPaginated", 10, 0).Return([]models.Regency{}, 0, errors.New("db error"))

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies", nil)
	w := httptest.NewRecorder()
	h.GetRegencies(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetRegencies_LoadAll(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegencies").Return([]models.Regency{{ID: 7201, Name: "Kab. Banggai"}}, nil)

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies?load_all=true", nil)
	w := httptest.NewRecorder()
	h.GetRegencies(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetRegencies_LoadAll_Error(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegencies").Return([]models.Regency{}, errors.New("db error"))

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies?load_all=true", nil)
	w := httptest.NewRecorder()
	h.GetRegencies(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestGetRegencyByID_Success(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegencyByID", 7201).Return(&models.Regency{ID: 7201, Name: "Kab. Banggai"}, nil)

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies/7201", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/regencies/{code}", h.GetRegencyByID)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetRegencyByID_NotFound(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegencyByID", 9999).Return(nil, nil)

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies/9999", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/regencies/{code}", h.GetRegencyByID)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	svc.AssertExpectations(t)
}

func TestGetRegencyByID_InvalidCode(t *testing.T) {
	svc := new(MockRegencyService)
	h := NewRegencyHandler(svc)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies/abc", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/regencies/{code}", h.GetRegencyByID)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetRegencyCases_Success(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegencyCases", 7201).Return([]models.RegencyCase{{ID: 1, RegencyID: 7201}}, nil)

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies/7201/cases", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/regencies/{code}/cases", h.GetRegencyCases)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetRegencyCases_NotFound(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegencyCases", 9999).Return([]models.RegencyCase(nil), nil)

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies/9999/cases", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/regencies/{code}/cases", h.GetRegencyCases)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	svc.AssertExpectations(t)
}

func TestGetRegencies_IncludeLatestCase(t *testing.T) {
	svc := new(MockRegencyService)
	rt := 0.8
	svc.On("GetRegencies").Return([]models.Regency{{ID: 7201, Name: "Banggai"}, {ID: 7202, Name: "Poso"}}, nil)
	svc.On("GetLatestRegencyCases").Return([]models.RegencyCase{
		{RegencyID: 7201, Positive: 3, CumulativePositive: 90, Rt: &rt, Regency: &models.Regency{ID: 7201}},
	}, nil)

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies?load_all=true&include_latest_case=true", nil)
	w := httptest.NewRecorder()
	h.GetRegencies(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var body struct {
		Data []map[string]interface{} `json:"data"`
	}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &body))
	assert.Len(t, body.Data, 2)
	assert.Equal(t, "Banggai", body.Data[0]["name"])
	latest, ok := body.Data[0]["latest_case"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(90), latest["cumulative_positive"])
	assert.Equal(t, 0.8, latest["rt"])
	assert.NotContains(t, latest, "regency")
	assert.Nil(t, body.Data[1]["latest_case"])
	svc.AssertExpectations(t)
}

func TestGetRegencies_IncludeLatestCase_Paginated(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegenciesPaginated", 10, 0).Return([]models.Regency{{ID: 7201, Name: "Banggai"}}, 1, nil)
	svc.On("GetLatestRegencyCases").Return([]models.RegencyCase{{RegencyID: 7201}}, nil)

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies?include_latest_case=true", nil)
	w := httptest.NewRecorder()
	h.GetRegencies(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"latest_case"`)
	assert.Contains(t, w.Body.String(), `"pagination"`)
}

func TestGetRegencies_IncludeLatestCase_Error(t *testing.T) {
	svc := new(MockRegencyService)
	svc.On("GetRegencies").Return([]models.Regency{{ID: 7201}}, nil)
	svc.On("GetLatestRegencyCases").Return([]models.RegencyCase{}, errors.New("db error"))

	h := NewRegencyHandler(svc)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/regencies?load_all=true&include_latest_case=true", nil)
	w := httptest.NewRecorder()
	h.GetRegencies(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
