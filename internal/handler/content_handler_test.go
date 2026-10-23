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

type MockContentService struct{ mock.Mock }

func (m *MockContentService) GetBanners(inactive bool, limit, offset int) ([]models.Banner, int, error) {
	a := m.Called(inactive, limit, offset)
	return a.Get(0).([]models.Banner), a.Int(1), a.Error(2)
}
func (m *MockContentService) GetInfographics(limit, offset int) ([]models.Infographic, int, error) {
	a := m.Called(limit, offset)
	return a.Get(0).([]models.Infographic), a.Int(1), a.Error(2)
}
func (m *MockContentService) GetDonations(limit, offset int) ([]models.Donation, int, error) {
	a := m.Called(limit, offset)
	return a.Get(0).([]models.Donation), a.Int(1), a.Error(2)
}
func (m *MockContentService) GetTelemedicines(limit, offset int) ([]models.Telemedicine, int, error) {
	a := m.Called(limit, offset)
	return a.Get(0).([]models.Telemedicine), a.Int(1), a.Error(2)
}
func (m *MockContentService) GetPartners(f *bool, limit, offset int) ([]models.Partner, int, error) {
	a := m.Called(f, limit, offset)
	return a.Get(0).([]models.Partner), a.Int(1), a.Error(2)
}
func (m *MockContentService) GetContacts(t string, id int64, limit, offset int) ([]models.ContactEntry, int, error) {
	a := m.Called(t, id, limit, offset)
	return a.Get(0).([]models.ContactEntry), a.Int(1), a.Error(2)
}

func doContent(h http.HandlerFunc, url string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	h(w, httptest.NewRequest(http.MethodGet, url, nil))
	return w
}

func TestGetBanners_Paginated(t *testing.T) {
	svc := new(MockContentService)
	svc.On("GetBanners", false, 10, 0).Return([]models.Banner{{ID: 1, Title: "x", Link: "https://a/b", IsActive: true}}, 25, nil)
	w := doContent(NewContentHandler(svc).GetBanners, "/api/v1/banners")

	assert.Equal(t, http.StatusOK, w.Code)
	var resp struct {
		Status string `json:"status"`
		Data   struct {
			Data       []map[string]interface{} `json:"data"`
			Pagination PaginationMeta           `json:"pagination"`
		} `json:"data"`
	}
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.Equal(t, "success", resp.Status)
	assert.Equal(t, "https://a/b", resp.Data.Data[0]["link"])
	assert.Equal(t, 25, resp.Data.Pagination.Total)
	assert.Equal(t, 3, resp.Data.Pagination.TotalPages)
	assert.True(t, resp.Data.Pagination.HasNext)
	svc.AssertExpectations(t)
}

func TestGetBanners_IncludeInactiveAndPage(t *testing.T) {
	svc := new(MockContentService)
	svc.On("GetBanners", true, 5, 5).Return([]models.Banner{}, 7, nil)
	w := doContent(NewContentHandler(svc).GetBanners, "/api/v1/banners?include_inactive=true&page=2&per_page=5")
	assert.Equal(t, http.StatusOK, w.Code)
	svc.AssertExpectations(t)
}

func TestGetBanners_LoadAll(t *testing.T) {
	svc := new(MockContentService)
	svc.On("GetBanners", false, 0, 0).Return([]models.Banner{{ID: 1}}, 1, nil)
	w := doContent(NewContentHandler(svc).GetBanners, "/api/v1/banners?load_all=true")
	assert.Equal(t, http.StatusOK, w.Code)
	var resp Response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resp))
	assert.IsType(t, []interface{}{}, resp.Data) // plain list, no pagination envelope
	svc.AssertExpectations(t)
}

func TestGetBanners_Error(t *testing.T) {
	svc := new(MockContentService)
	svc.On("GetBanners", false, 10, 0).Return([]models.Banner(nil), 0, errors.New("db down"))
	w := doContent(NewContentHandler(svc).GetBanners, "/api/v1/banners")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "db down")
}

func TestGetInfographics(t *testing.T) {
	svc := new(MockContentService)
	svc.On("GetInfographics", 10, 0).Return([]models.Infographic{{ID: 1, Images: []models.InfographicImage{{ID: 1, Queue: 1, ImageURL: "u"}}}}, 1, nil)
	w := doContent(NewContentHandler(svc).GetInfographics, "/api/v1/infographics")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"image_url":"u"`)
	svc.AssertExpectations(t)
}

func TestGetDonations(t *testing.T) {
	svc := new(MockContentService)
	svc.On("GetDonations", 10, 0).Return([]models.Donation{{ID: 1, BankAccounts: []models.BankAccount{}, Contacts: []models.ContactEntry{}}}, 1, nil)
	w := doContent(NewContentHandler(svc).GetDonations, "/api/v1/donations")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"delivery_address":null`)
	assert.Contains(t, w.Body.String(), `"bank_accounts":[]`)
	svc.AssertExpectations(t)
}

func TestGetTelemedicines(t *testing.T) {
	svc := new(MockContentService)
	svc.On("GetTelemedicines", 10, 0).Return([]models.Telemedicine{{ID: 1, IsFree: true}}, 1, nil)
	w := doContent(NewContentHandler(svc).GetTelemedicines, "/api/v1/telemedicines")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"is_free":true`)
	svc.AssertExpectations(t)
}

func TestGetPartners_Filter(t *testing.T) {
	svc := new(MockContentService)
	yes, no := true, false
	svc.On("GetPartners", &yes, 10, 0).Return([]models.Partner{{ID: 1}}, 1, nil)
	svc.On("GetPartners", &no, 10, 0).Return([]models.Partner{}, 0, nil)
	svc.On("GetPartners", (*bool)(nil), 10, 0).Return([]models.Partner{}, 0, nil)
	h := NewContentHandler(svc)
	assert.Equal(t, http.StatusOK, doContent(h.GetPartners, "/api/v1/partners?in_footer=true").Code)
	assert.Equal(t, http.StatusOK, doContent(h.GetPartners, "/api/v1/partners?in_footer=false").Code)
	assert.Equal(t, http.StatusOK, doContent(h.GetPartners, "/api/v1/partners").Code)
	svc.AssertExpectations(t)
}

func TestGetContacts_Filters(t *testing.T) {
	svc := new(MockContentService)
	svc.On("GetContacts", "hospital", int64(4), 10, 0).Return([]models.ContactEntry{{ID: 1, OwnerType: "hospital", OwnerID: 4}}, 1, nil)
	w := doContent(NewContentHandler(svc).GetContacts, "/api/v1/contacts?owner_type=hospital&owner_id=4")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"owner_type":"hospital"`)
	svc.AssertExpectations(t)
}

func TestGetContacts_InvalidOwnerType(t *testing.T) {
	svc := new(MockContentService)
	w := doContent(NewContentHandler(svc).GetContacts, "/api/v1/contacts?owner_type=bogus")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	svc.AssertNotCalled(t, "GetContacts")
}

func TestContentRoutes_Registered(t *testing.T) {
	router := SetupRoutes(Services{ContentService: new(MockContentService)}, nil, false)
	for _, path := range []string{"banners", "infographics", "donations", "telemedicines", "partners", "contacts"} {
		var m mux.RouteMatch
		assert.True(t, router.Match(httptest.NewRequest(http.MethodGet, "/api/v1/"+path, nil), &m), path)
	}
}
