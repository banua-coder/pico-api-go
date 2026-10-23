package service

import (
	"errors"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockContentRepository mocks repository.ContentRepositoryInterface and ContentServiceInterface.
type MockContentRepository struct{ mock.Mock }

func (m *MockContentRepository) GetBanners(inactive bool, limit, offset int) ([]models.Banner, int, error) {
	a := m.Called(inactive, limit, offset)
	return a.Get(0).([]models.Banner), a.Int(1), a.Error(2)
}
func (m *MockContentRepository) GetInfographics(limit, offset int) ([]models.Infographic, int, error) {
	a := m.Called(limit, offset)
	return a.Get(0).([]models.Infographic), a.Int(1), a.Error(2)
}
func (m *MockContentRepository) GetDonations(limit, offset int) ([]models.Donation, int, error) {
	a := m.Called(limit, offset)
	return a.Get(0).([]models.Donation), a.Int(1), a.Error(2)
}
func (m *MockContentRepository) GetTelemedicines(limit, offset int) ([]models.Telemedicine, int, error) {
	a := m.Called(limit, offset)
	return a.Get(0).([]models.Telemedicine), a.Int(1), a.Error(2)
}
func (m *MockContentRepository) GetPartners(f *bool, limit, offset int) ([]models.Partner, int, error) {
	a := m.Called(f, limit, offset)
	return a.Get(0).([]models.Partner), a.Int(1), a.Error(2)
}
func (m *MockContentRepository) GetContacts(t string, id int64, limit, offset int) ([]models.ContactEntry, int, error) {
	a := m.Called(t, id, limit, offset)
	return a.Get(0).([]models.ContactEntry), a.Int(1), a.Error(2)
}

func TestContentService_PassThrough(t *testing.T) {
	repo := new(MockContentRepository)
	svc := NewContentService(repo)
	repo.On("GetBanners", false, 10, 0).Return([]models.Banner{{ID: 1}}, 1, nil)
	repo.On("GetInfographics", 10, 0).Return([]models.Infographic{{ID: 1}}, 1, nil)
	repo.On("GetDonations", 10, 0).Return([]models.Donation{{ID: 1}}, 1, nil)
	repo.On("GetTelemedicines", 10, 0).Return([]models.Telemedicine{{ID: 1}}, 1, nil)
	repo.On("GetPartners", (*bool)(nil), 10, 0).Return([]models.Partner{{ID: 1}}, 1, nil)
	repo.On("GetContacts", "hospital", int64(2), 10, 0).Return([]models.ContactEntry{{ID: 1}}, 1, nil)

	b, n, err := svc.GetBanners(false, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, 1, n)
	assert.Equal(t, int64(1), b[0].ID)
	_, _, err = svc.GetInfographics(10, 0)
	assert.NoError(t, err)
	_, _, err = svc.GetDonations(10, 0)
	assert.NoError(t, err)
	_, _, err = svc.GetTelemedicines(10, 0)
	assert.NoError(t, err)
	_, _, err = svc.GetPartners(nil, 10, 0)
	assert.NoError(t, err)
	_, _, err = svc.GetContacts("hospital", 2, 10, 0)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestContentService_Error(t *testing.T) {
	repo := new(MockContentRepository)
	repo.On("GetBanners", false, 10, 0).Return([]models.Banner(nil), 0, errors.New("db"))
	_, _, err := NewContentService(repo).GetBanners(false, 10, 0)
	assert.Error(t, err)
}

func TestCachedContentService_CachesPerKey(t *testing.T) {
	inner := new(MockContentRepository)
	c := cache.New(time.Minute)
	svc := NewCachedContentService(inner, c)
	yes := true
	inner.On("GetBanners", false, 10, 0).Return([]models.Banner{{ID: 1}}, 1, nil).Once()
	inner.On("GetBanners", true, 10, 0).Return([]models.Banner{{ID: 1}, {ID: 2}}, 2, nil).Once()
	inner.On("GetInfographics", 10, 0).Return([]models.Infographic{{ID: 1}}, 1, nil).Once()
	inner.On("GetDonations", 10, 0).Return([]models.Donation{{ID: 1}}, 1, nil).Once()
	inner.On("GetTelemedicines", 10, 0).Return([]models.Telemedicine{{ID: 1}}, 1, nil).Once()
	inner.On("GetPartners", &yes, 10, 0).Return([]models.Partner{{ID: 1}}, 1, nil).Once()
	inner.On("GetContacts", "", int64(0), 10, 0).Return([]models.ContactEntry{{ID: 1}}, 1, nil).Once()

	for i := 0; i < 2; i++ { // second round must be served from cache (.Once would panic otherwise)
		b, n, err := svc.GetBanners(false, 10, 0)
		assert.NoError(t, err)
		assert.Equal(t, 1, n)
		assert.Len(t, b, 1)
		b, n, err = svc.GetBanners(true, 10, 0)
		assert.NoError(t, err)
		assert.Equal(t, 2, n)
		assert.Len(t, b, 2)
		_, _, err = svc.GetInfographics(10, 0)
		assert.NoError(t, err)
		_, _, err = svc.GetDonations(10, 0)
		assert.NoError(t, err)
		_, _, err = svc.GetTelemedicines(10, 0)
		assert.NoError(t, err)
		_, _, err = svc.GetPartners(&yes, 10, 0)
		assert.NoError(t, err)
		_, _, err = svc.GetContacts("", 0, 10, 0)
		assert.NoError(t, err)
	}
	inner.AssertExpectations(t)
}

func TestCachedContentService_ErrorNotCached(t *testing.T) {
	inner := new(MockContentRepository)
	svc := NewCachedContentService(inner, cache.New(time.Minute))
	inner.On("GetInfographics", 10, 0).Return([]models.Infographic(nil), 0, errors.New("db")).Once()
	inner.On("GetInfographics", 10, 0).Return([]models.Infographic{{ID: 1}}, 1, nil).Once()

	_, _, err := svc.GetInfographics(10, 0)
	assert.Error(t, err)
	got, _, err := svc.GetInfographics(10, 0)
	assert.NoError(t, err)
	assert.Len(t, got, 1)
}

func TestCachedContentService_EmptyPageNotCached(t *testing.T) {
	inner := new(MockContentRepository)
	svc := NewCachedContentService(inner, cache.New(time.Minute))
	inner.On("GetInfographics", 10, 9990).Return([]models.Infographic{}, 3, nil).Twice()
	for i := 0; i < 2; i++ {
		_, _, err := svc.GetInfographics(10, 9990)
		assert.NoError(t, err)
	}
	inner.AssertExpectations(t)
}
