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

// MockRegencyService mocks RegencyServiceInterface
type MockRegencyService struct {
	mock.Mock
}

func (m *MockRegencyService) GetRegencies() ([]models.Regency, error) {
	args := m.Called()
	return args.Get(0).([]models.Regency), args.Error(1)
}

func (m *MockRegencyService) GetRegenciesPaginated(limit, offset int) ([]models.Regency, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.Regency), args.Int(1), args.Error(2)
}

func (m *MockRegencyService) GetRegencyByID(id int) (*models.Regency, error) {
	args := m.Called(id)
	res := args.Get(0)
	if res == nil {
		return nil, args.Error(1)
	}
	return res.(*models.Regency), args.Error(1)
}

func (m *MockRegencyService) GetRegencyCases(regencyID int) ([]models.RegencyCase, error) {
	args := m.Called(regencyID)
	return args.Get(0).([]models.RegencyCase), args.Error(1)
}

func (m *MockRegencyService) GetLatestRegencyCases() ([]models.RegencyCase, error) {
	args := m.Called()
	return args.Get(0).([]models.RegencyCase), args.Error(1)
}

func TestCachedRegencyService_GetRegencies(t *testing.T) {
	t.Run("cache miss - calls underlying service", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := []models.Regency{{}}
		mockSvc.On("GetRegencies").Return(expected, nil).Once()

		result, err := svc.GetRegencies()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockSvc.AssertExpectations(t)
	})

	t.Run("cache hit - does not call underlying service again", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := []models.Regency{{}}
		mockSvc.On("GetRegencies").Return(expected, nil).Once()

		svc.GetRegencies()
		result, err := svc.GetRegencies()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockSvc.AssertNumberOfCalls(t, "GetRegencies", 1)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		mockSvc.On("GetRegencies").Return([]models.Regency{}, errors.New("db error"))
		_, err := svc.GetRegencies()
		assert.Error(t, err)
	})
}

func TestCachedRegencyService_GetRegenciesPaginated(t *testing.T) {
	t.Run("cache miss", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := []models.Regency{{}}
		mockSvc.On("GetRegenciesPaginated", 10, 0).Return(expected, 1, nil).Once()

		items, total, err := svc.GetRegenciesPaginated(10, 0)
		assert.NoError(t, err)
		assert.Equal(t, expected, items)
		assert.Equal(t, 1, total)
	})

	t.Run("cache hit", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := []models.Regency{{}}
		mockSvc.On("GetRegenciesPaginated", 10, 0).Return(expected, 1, nil).Once()

		svc.GetRegenciesPaginated(10, 0)
		items, total, err := svc.GetRegenciesPaginated(10, 0)
		assert.NoError(t, err)
		assert.Equal(t, expected, items)
		assert.Equal(t, 1, total)
		mockSvc.AssertNumberOfCalls(t, "GetRegenciesPaginated", 1)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		mockSvc.On("GetRegenciesPaginated", 10, 0).Return([]models.Regency{}, 0, errors.New("err"))
		_, _, err := svc.GetRegenciesPaginated(10, 0)
		assert.Error(t, err)
	})
}

func TestCachedRegencyService_GetRegencyByID(t *testing.T) {
	t.Run("cache miss", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := &models.Regency{}
		mockSvc.On("GetRegencyByID", 1).Return(expected, nil).Once()

		result, err := svc.GetRegencyByID(1)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("cache hit", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := &models.Regency{}
		mockSvc.On("GetRegencyByID", 1).Return(expected, nil).Once()

		svc.GetRegencyByID(1)
		result, err := svc.GetRegencyByID(1)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockSvc.AssertNumberOfCalls(t, "GetRegencyByID", 1)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		mockSvc.On("GetRegencyByID", 1).Return(nil, errors.New("not found"))
		_, err := svc.GetRegencyByID(1)
		assert.Error(t, err)
	})
}

func TestCachedRegencyService_GetRegencyCases(t *testing.T) {
	t.Run("cache miss", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := []models.RegencyCase{{}}
		mockSvc.On("GetRegencyCases", 1).Return(expected, nil).Once()

		result, err := svc.GetRegencyCases(1)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("cache hit", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := []models.RegencyCase{{}}
		mockSvc.On("GetRegencyCases", 1).Return(expected, nil).Once()

		svc.GetRegencyCases(1)
		result, err := svc.GetRegencyCases(1)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockSvc.AssertNumberOfCalls(t, "GetRegencyCases", 1)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		mockSvc.On("GetRegencyCases", 1).Return([]models.RegencyCase{}, errors.New("err"))
		_, err := svc.GetRegencyCases(1)
		assert.Error(t, err)
	})
}

func TestCachedRegencyService_GetLatestRegencyCases(t *testing.T) {
	t.Run("cache miss", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := []models.RegencyCase{{}}
		mockSvc.On("GetLatestRegencyCases").Return(expected, nil).Once()

		result, err := svc.GetLatestRegencyCases()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("cache hit", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		expected := []models.RegencyCase{{}}
		mockSvc.On("GetLatestRegencyCases").Return(expected, nil).Once()

		svc.GetLatestRegencyCases()
		result, err := svc.GetLatestRegencyCases()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockSvc.AssertNumberOfCalls(t, "GetLatestRegencyCases", 1)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc := new(MockRegencyService)
		c := cache.New(time.Hour)
		svc := NewCachedRegencyService(mockSvc, c)

		mockSvc.On("GetLatestRegencyCases").Return([]models.RegencyCase{}, errors.New("err"))
		_, err := svc.GetLatestRegencyCases()
		assert.Error(t, err)
	})
}
