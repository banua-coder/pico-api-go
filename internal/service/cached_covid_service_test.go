package service

import (
	"errors"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/cache"
	"github.com/banua-coder/pico-api-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCovidService mocks CovidService interface
type MockCovidService struct {
	mock.Mock
}

func (m *MockCovidService) GetNationalCases() ([]models.NationalCase, error) {
	args := m.Called()
	return args.Get(0).([]models.NationalCase), args.Error(1)
}
func (m *MockCovidService) GetNationalCasesSorted(s utils.SortParams) ([]models.NationalCase, error) {
	args := m.Called(s)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}
func (m *MockCovidService) GetNationalCasesPaginated(limit, offset int) ([]models.NationalCase, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.NationalCase), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetNationalCasesPaginatedSorted(limit, offset int, s utils.SortParams) ([]models.NationalCase, int, error) {
	args := m.Called(limit, offset, s)
	return args.Get(0).([]models.NationalCase), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetNationalCasesByDateRange(start, end string) ([]models.NationalCase, error) {
	args := m.Called(start, end)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}
func (m *MockCovidService) GetNationalCasesByDateRangeSorted(start, end string, s utils.SortParams) ([]models.NationalCase, error) {
	args := m.Called(start, end, s)
	return args.Get(0).([]models.NationalCase), args.Error(1)
}
func (m *MockCovidService) GetNationalCasesByDateRangePaginated(start, end string, limit, offset int) ([]models.NationalCase, int, error) {
	args := m.Called(start, end, limit, offset)
	return args.Get(0).([]models.NationalCase), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetNationalCasesByDateRangePaginatedSorted(start, end string, limit, offset int, s utils.SortParams) ([]models.NationalCase, int, error) {
	args := m.Called(start, end, limit, offset, s)
	return args.Get(0).([]models.NationalCase), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetLatestNationalCase() (*models.NationalCase, error) {
	args := m.Called()
	res := args.Get(0)
	if res == nil {
		return nil, args.Error(1)
	}
	return res.(*models.NationalCase), args.Error(1)
}
func (m *MockCovidService) GetNationalCaseByDay(day int64) (*models.NationalCase, error) {
	args := m.Called(day)
	res := args.Get(0)
	if res == nil {
		return nil, args.Error(1)
	}
	return res.(*models.NationalCase), args.Error(1)
}
func (m *MockCovidService) GetProvinces() ([]models.Province, error) {
	args := m.Called()
	return args.Get(0).([]models.Province), args.Error(1)
}
func (m *MockCovidService) GetProvinceByID(id string) (*models.Province, error) {
	args := m.Called(id)
	res := args.Get(0)
	if res == nil {
		return nil, args.Error(1)
	}
	return res.(*models.Province), args.Error(1)
}
func (m *MockCovidService) GetProvincesWithLatestCase() ([]models.ProvinceWithLatestCase, error) {
	args := m.Called()
	return args.Get(0).([]models.ProvinceWithLatestCase), args.Error(1)
}
func (m *MockCovidService) GetProvinceCases(pid string) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(pid)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}
func (m *MockCovidService) GetProvinceCasesSorted(pid string, s utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(pid, s)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}
func (m *MockCovidService) GetProvinceCasesPaginated(pid string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(pid, limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetProvinceCasesPaginatedSorted(pid string, limit, offset int, s utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(pid, limit, offset, s)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetProvinceCasesByDateRange(pid, start, end string) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(pid, start, end)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}
func (m *MockCovidService) GetProvinceCasesByDateRangeSorted(pid, start, end string, s utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(pid, start, end, s)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}
func (m *MockCovidService) GetProvinceCasesByDateRangePaginated(pid, start, end string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(pid, start, end, limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetProvinceCasesByDateRangePaginatedSorted(pid, start, end string, limit, offset int, s utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(pid, start, end, limit, offset, s)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetAllProvinceCases() ([]models.ProvinceCaseWithDate, error) {
	args := m.Called()
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}
func (m *MockCovidService) GetAllProvinceCasesSorted(s utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(s)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}
func (m *MockCovidService) GetAllProvinceCasesPaginated(limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetAllProvinceCasesPaginatedSorted(limit, offset int, s utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(limit, offset, s)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetAllProvinceCasesByDateRange(start, end string) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(start, end)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}
func (m *MockCovidService) GetAllProvinceCasesByDateRangeSorted(start, end string, s utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	args := m.Called(start, end, s)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Error(1)
}
func (m *MockCovidService) GetAllProvinceCasesByDateRangePaginated(start, end string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(start, end, limit, offset)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}
func (m *MockCovidService) GetAllProvinceCasesByDateRangePaginatedSorted(start, end string, limit, offset int, s utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	args := m.Called(start, end, limit, offset, s)
	return args.Get(0).([]models.ProvinceCaseWithDate), args.Int(1), args.Error(2)
}

func newTestCache() *cache.Cache {
	return cache.New(time.Hour)
}

func TestCachedCovidService_GetNationalCases(t *testing.T) {
	t.Run("cache miss - calls underlying service", func(t *testing.T) {
		mockSvc := new(MockCovidService)
		c := newTestCache()
		svc := NewCachedCovidService(mockSvc, c)

		expected := []models.NationalCase{{}}
		mockSvc.On("GetNationalCases").Return(expected, nil).Once()

		result, err := svc.GetNationalCases()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockSvc.AssertExpectations(t)
	})

	t.Run("cache hit - does not call underlying service again", func(t *testing.T) {
		mockSvc := new(MockCovidService)
		c := newTestCache()
		svc := NewCachedCovidService(mockSvc, c)

		expected := []models.NationalCase{{}}
		mockSvc.On("GetNationalCases").Return(expected, nil).Once()

		svc.GetNationalCases() // prime cache
		result, err := svc.GetNationalCases() // should hit cache
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockSvc.AssertNumberOfCalls(t, "GetNationalCases", 1)
	})

	t.Run("error - returns error from underlying service", func(t *testing.T) {
		mockSvc := new(MockCovidService)
		c := newTestCache()
		svc := NewCachedCovidService(mockSvc, c)

		mockSvc.On("GetNationalCases").Return([]models.NationalCase{}, errors.New("db error"))

		_, err := svc.GetNationalCases()
		assert.Error(t, err)
	})
}

func TestCachedCovidService_GetNationalCasesSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.NationalCase{{}}
	mockSvc.On("GetNationalCasesSorted", sp).Return(expected, nil).Once()

	result, err := svc.GetNationalCasesSorted(sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	// cache hit
	result2, err := svc.GetNationalCasesSorted(sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, result2)
	mockSvc.AssertNumberOfCalls(t, "GetNationalCasesSorted", 1)
}

func TestCachedCovidService_GetNationalCasesSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "desc"}
	mockSvc.On("GetNationalCasesSorted", sp).Return([]models.NationalCase{}, errors.New("err"))
	_, err := svc.GetNationalCasesSorted(sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetNationalCasesPaginated(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.NationalCase{{}}
	mockSvc.On("GetNationalCasesPaginated", 10, 0).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetNationalCasesPaginated(10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	// cache hit
	svc.GetNationalCasesPaginated(10, 0)
	mockSvc.AssertNumberOfCalls(t, "GetNationalCasesPaginated", 1)
}

func TestCachedCovidService_GetNationalCasesPaginated_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetNationalCasesPaginated", 10, 0).Return([]models.NationalCase{}, 0, errors.New("err"))
	_, _, err := svc.GetNationalCasesPaginated(10, 0)
	assert.Error(t, err)
}

func TestCachedCovidService_GetNationalCasesPaginatedSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.NationalCase{{}}
	mockSvc.On("GetNationalCasesPaginatedSorted", 10, 0, sp).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetNationalCasesPaginatedSorted(10, 0, sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetNationalCasesPaginatedSorted(10, 0, sp)
	mockSvc.AssertNumberOfCalls(t, "GetNationalCasesPaginatedSorted", 1)
}

func TestCachedCovidService_GetNationalCasesPaginatedSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetNationalCasesPaginatedSorted", 10, 0, sp).Return([]models.NationalCase{}, 0, errors.New("err"))
	_, _, err := svc.GetNationalCasesPaginatedSorted(10, 0, sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetNationalCasesByDateRange(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.NationalCase{{}}
	mockSvc.On("GetNationalCasesByDateRange", "2021-01-01", "2021-12-31").Return(expected, nil).Once()

	result, err := svc.GetNationalCasesByDateRange("2021-01-01", "2021-12-31")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetNationalCasesByDateRange("2021-01-01", "2021-12-31")
	mockSvc.AssertNumberOfCalls(t, "GetNationalCasesByDateRange", 1)
}

func TestCachedCovidService_GetNationalCasesByDateRange_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetNationalCasesByDateRange", "2021-01-01", "2021-12-31").Return([]models.NationalCase{}, errors.New("err"))
	_, err := svc.GetNationalCasesByDateRange("2021-01-01", "2021-12-31")
	assert.Error(t, err)
}

func TestCachedCovidService_GetNationalCasesByDateRangeSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.NationalCase{{}}
	mockSvc.On("GetNationalCasesByDateRangeSorted", "2021-01-01", "2021-12-31", sp).Return(expected, nil).Once()

	result, err := svc.GetNationalCasesByDateRangeSorted("2021-01-01", "2021-12-31", sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetNationalCasesByDateRangeSorted("2021-01-01", "2021-12-31", sp)
	mockSvc.AssertNumberOfCalls(t, "GetNationalCasesByDateRangeSorted", 1)
}

func TestCachedCovidService_GetNationalCasesByDateRangeSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetNationalCasesByDateRangeSorted", "2021-01-01", "2021-12-31", sp).Return([]models.NationalCase{}, errors.New("err"))
	_, err := svc.GetNationalCasesByDateRangeSorted("2021-01-01", "2021-12-31", sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetNationalCasesByDateRangePaginated(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.NationalCase{{}}
	mockSvc.On("GetNationalCasesByDateRangePaginated", "2021-01-01", "2021-12-31", 10, 0).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetNationalCasesByDateRangePaginated("2021-01-01", "2021-12-31", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetNationalCasesByDateRangePaginated("2021-01-01", "2021-12-31", 10, 0)
	mockSvc.AssertNumberOfCalls(t, "GetNationalCasesByDateRangePaginated", 1)
}

func TestCachedCovidService_GetNationalCasesByDateRangePaginated_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetNationalCasesByDateRangePaginated", "2021-01-01", "2021-12-31", 10, 0).Return([]models.NationalCase{}, 0, errors.New("err"))
	_, _, err := svc.GetNationalCasesByDateRangePaginated("2021-01-01", "2021-12-31", 10, 0)
	assert.Error(t, err)
}

func TestCachedCovidService_GetNationalCasesByDateRangePaginatedSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.NationalCase{{}}
	mockSvc.On("GetNationalCasesByDateRangePaginatedSorted", "2021-01-01", "2021-12-31", 10, 0, sp).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetNationalCasesByDateRangePaginatedSorted("2021-01-01", "2021-12-31", 10, 0, sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetNationalCasesByDateRangePaginatedSorted("2021-01-01", "2021-12-31", 10, 0, sp)
	mockSvc.AssertNumberOfCalls(t, "GetNationalCasesByDateRangePaginatedSorted", 1)
}

func TestCachedCovidService_GetNationalCasesByDateRangePaginatedSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetNationalCasesByDateRangePaginatedSorted", "2021-01-01", "2021-12-31", 10, 0, sp).Return([]models.NationalCase{}, 0, errors.New("err"))
	_, _, err := svc.GetNationalCasesByDateRangePaginatedSorted("2021-01-01", "2021-12-31", 10, 0, sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetLatestNationalCase(t *testing.T) {
	t.Run("success with cache", func(t *testing.T) {
		mockSvc := new(MockCovidService)
		c := newTestCache()
		svc := NewCachedCovidService(mockSvc, c)

		expected := &models.NationalCase{}
		mockSvc.On("GetLatestNationalCase").Return(expected, nil).Once()

		result, err := svc.GetLatestNationalCase()
		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		svc.GetLatestNationalCase()
		mockSvc.AssertNumberOfCalls(t, "GetLatestNationalCase", 1)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc := new(MockCovidService)
		c := newTestCache()
		svc := NewCachedCovidService(mockSvc, c)

		mockSvc.On("GetLatestNationalCase").Return(nil, errors.New("err"))
		_, err := svc.GetLatestNationalCase()
		assert.Error(t, err)
	})
}

func TestCachedCovidService_GetNationalCaseByDay(t *testing.T) {
	t.Run("success with cache", func(t *testing.T) {
		mockSvc := new(MockCovidService)
		c := newTestCache()
		svc := NewCachedCovidService(mockSvc, c)

		expected := &models.NationalCase{}
		mockSvc.On("GetNationalCaseByDay", int64(1)).Return(expected, nil).Once()

		result, err := svc.GetNationalCaseByDay(1)
		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		svc.GetNationalCaseByDay(1)
		mockSvc.AssertNumberOfCalls(t, "GetNationalCaseByDay", 1)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc := new(MockCovidService)
		c := newTestCache()
		svc := NewCachedCovidService(mockSvc, c)

		mockSvc.On("GetNationalCaseByDay", int64(1)).Return(nil, errors.New("err"))
		_, err := svc.GetNationalCaseByDay(1)
		assert.Error(t, err)
	})
}

func TestCachedCovidService_GetProvinces(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.Province{{}}
	mockSvc.On("GetProvinces").Return(expected, nil).Once()

	result, err := svc.GetProvinces()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetProvinces()
	mockSvc.AssertNumberOfCalls(t, "GetProvinces", 1)
}

func TestCachedCovidService_GetProvinces_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetProvinces").Return([]models.Province{}, errors.New("err"))
	_, err := svc.GetProvinces()
	assert.Error(t, err)
}

func TestCachedCovidService_GetProvinceByID(t *testing.T) {
	t.Run("success with cache", func(t *testing.T) {
		mockSvc := new(MockCovidService)
		c := newTestCache()
		svc := NewCachedCovidService(mockSvc, c)

		expected := &models.Province{}
		mockSvc.On("GetProvinceByID", "1").Return(expected, nil).Once()

		result, err := svc.GetProvinceByID("1")
		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		svc.GetProvinceByID("1")
		mockSvc.AssertNumberOfCalls(t, "GetProvinceByID", 1)
	})

	t.Run("error", func(t *testing.T) {
		mockSvc := new(MockCovidService)
		c := newTestCache()
		svc := NewCachedCovidService(mockSvc, c)

		mockSvc.On("GetProvinceByID", "1").Return(nil, errors.New("err"))
		_, err := svc.GetProvinceByID("1")
		assert.Error(t, err)
	})
}

func TestCachedCovidService_GetProvincesWithLatestCase(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.ProvinceWithLatestCase{{}}
	mockSvc.On("GetProvincesWithLatestCase").Return(expected, nil).Once()

	result, err := svc.GetProvincesWithLatestCase()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetProvincesWithLatestCase()
	mockSvc.AssertNumberOfCalls(t, "GetProvincesWithLatestCase", 1)
}

func TestCachedCovidService_GetProvincesWithLatestCase_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetProvincesWithLatestCase").Return([]models.ProvinceWithLatestCase{}, errors.New("err"))
	_, err := svc.GetProvincesWithLatestCase()
	assert.Error(t, err)
}

func TestCachedCovidService_GetProvinceCases(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetProvinceCases", "p1").Return(expected, nil).Once()

	result, err := svc.GetProvinceCases("p1")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetProvinceCases("p1")
	mockSvc.AssertNumberOfCalls(t, "GetProvinceCases", 1)
}

func TestCachedCovidService_GetProvinceCases_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetProvinceCases", "p1").Return([]models.ProvinceCaseWithDate{}, errors.New("err"))
	_, err := svc.GetProvinceCases("p1")
	assert.Error(t, err)
}

func TestCachedCovidService_GetProvinceCasesSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetProvinceCasesSorted", "p1", sp).Return(expected, nil).Once()

	result, err := svc.GetProvinceCasesSorted("p1", sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetProvinceCasesSorted("p1", sp)
	mockSvc.AssertNumberOfCalls(t, "GetProvinceCasesSorted", 1)
}

func TestCachedCovidService_GetProvinceCasesSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetProvinceCasesSorted", "p1", sp).Return([]models.ProvinceCaseWithDate{}, errors.New("err"))
	_, err := svc.GetProvinceCasesSorted("p1", sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetProvinceCasesPaginated(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetProvinceCasesPaginated", "p1", 10, 0).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetProvinceCasesPaginated("p1", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetProvinceCasesPaginated("p1", 10, 0)
	mockSvc.AssertNumberOfCalls(t, "GetProvinceCasesPaginated", 1)
}

func TestCachedCovidService_GetProvinceCasesPaginated_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetProvinceCasesPaginated", "p1", 10, 0).Return([]models.ProvinceCaseWithDate{}, 0, errors.New("err"))
	_, _, err := svc.GetProvinceCasesPaginated("p1", 10, 0)
	assert.Error(t, err)
}

func TestCachedCovidService_GetProvinceCasesPaginatedSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetProvinceCasesPaginatedSorted", "p1", 10, 0, sp).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetProvinceCasesPaginatedSorted("p1", 10, 0, sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetProvinceCasesPaginatedSorted("p1", 10, 0, sp)
	mockSvc.AssertNumberOfCalls(t, "GetProvinceCasesPaginatedSorted", 1)
}

func TestCachedCovidService_GetProvinceCasesPaginatedSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetProvinceCasesPaginatedSorted", "p1", 10, 0, sp).Return([]models.ProvinceCaseWithDate{}, 0, errors.New("err"))
	_, _, err := svc.GetProvinceCasesPaginatedSorted("p1", 10, 0, sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetProvinceCasesByDateRange(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetProvinceCasesByDateRange", "p1", "2021-01-01", "2021-12-31").Return(expected, nil).Once()

	result, err := svc.GetProvinceCasesByDateRange("p1", "2021-01-01", "2021-12-31")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetProvinceCasesByDateRange("p1", "2021-01-01", "2021-12-31")
	mockSvc.AssertNumberOfCalls(t, "GetProvinceCasesByDateRange", 1)
}

func TestCachedCovidService_GetProvinceCasesByDateRange_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetProvinceCasesByDateRange", "p1", "2021-01-01", "2021-12-31").Return([]models.ProvinceCaseWithDate{}, errors.New("err"))
	_, err := svc.GetProvinceCasesByDateRange("p1", "2021-01-01", "2021-12-31")
	assert.Error(t, err)
}

func TestCachedCovidService_GetProvinceCasesByDateRangeSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetProvinceCasesByDateRangeSorted", "p1", "2021-01-01", "2021-12-31", sp).Return(expected, nil).Once()

	result, err := svc.GetProvinceCasesByDateRangeSorted("p1", "2021-01-01", "2021-12-31", sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetProvinceCasesByDateRangeSorted("p1", "2021-01-01", "2021-12-31", sp)
	mockSvc.AssertNumberOfCalls(t, "GetProvinceCasesByDateRangeSorted", 1)
}

func TestCachedCovidService_GetProvinceCasesByDateRangeSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetProvinceCasesByDateRangeSorted", "p1", "2021-01-01", "2021-12-31", sp).Return([]models.ProvinceCaseWithDate{}, errors.New("err"))
	_, err := svc.GetProvinceCasesByDateRangeSorted("p1", "2021-01-01", "2021-12-31", sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetProvinceCasesByDateRangePaginated(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetProvinceCasesByDateRangePaginated", "p1", "2021-01-01", "2021-12-31", 10, 0).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetProvinceCasesByDateRangePaginated("p1", "2021-01-01", "2021-12-31", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetProvinceCasesByDateRangePaginated("p1", "2021-01-01", "2021-12-31", 10, 0)
	mockSvc.AssertNumberOfCalls(t, "GetProvinceCasesByDateRangePaginated", 1)
}

func TestCachedCovidService_GetProvinceCasesByDateRangePaginated_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetProvinceCasesByDateRangePaginated", "p1", "2021-01-01", "2021-12-31", 10, 0).Return([]models.ProvinceCaseWithDate{}, 0, errors.New("err"))
	_, _, err := svc.GetProvinceCasesByDateRangePaginated("p1", "2021-01-01", "2021-12-31", 10, 0)
	assert.Error(t, err)
}

func TestCachedCovidService_GetProvinceCasesByDateRangePaginatedSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetProvinceCasesByDateRangePaginatedSorted", "p1", "2021-01-01", "2021-12-31", 10, 0, sp).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetProvinceCasesByDateRangePaginatedSorted("p1", "2021-01-01", "2021-12-31", 10, 0, sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetProvinceCasesByDateRangePaginatedSorted("p1", "2021-01-01", "2021-12-31", 10, 0, sp)
	mockSvc.AssertNumberOfCalls(t, "GetProvinceCasesByDateRangePaginatedSorted", 1)
}

func TestCachedCovidService_GetProvinceCasesByDateRangePaginatedSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetProvinceCasesByDateRangePaginatedSorted", "p1", "2021-01-01", "2021-12-31", 10, 0, sp).Return([]models.ProvinceCaseWithDate{}, 0, errors.New("err"))
	_, _, err := svc.GetProvinceCasesByDateRangePaginatedSorted("p1", "2021-01-01", "2021-12-31", 10, 0, sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetAllProvinceCases(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetAllProvinceCases").Return(expected, nil).Once()

	result, err := svc.GetAllProvinceCases()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetAllProvinceCases()
	mockSvc.AssertNumberOfCalls(t, "GetAllProvinceCases", 1)
}

func TestCachedCovidService_GetAllProvinceCases_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetAllProvinceCases").Return([]models.ProvinceCaseWithDate{}, errors.New("err"))
	_, err := svc.GetAllProvinceCases()
	assert.Error(t, err)
}

func TestCachedCovidService_GetAllProvinceCasesSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetAllProvinceCasesSorted", sp).Return(expected, nil).Once()

	result, err := svc.GetAllProvinceCasesSorted(sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetAllProvinceCasesSorted(sp)
	mockSvc.AssertNumberOfCalls(t, "GetAllProvinceCasesSorted", 1)
}

func TestCachedCovidService_GetAllProvinceCasesSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetAllProvinceCasesSorted", sp).Return([]models.ProvinceCaseWithDate{}, errors.New("err"))
	_, err := svc.GetAllProvinceCasesSorted(sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetAllProvinceCasesPaginated(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetAllProvinceCasesPaginated", 10, 0).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetAllProvinceCasesPaginated(10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetAllProvinceCasesPaginated(10, 0)
	mockSvc.AssertNumberOfCalls(t, "GetAllProvinceCasesPaginated", 1)
}

func TestCachedCovidService_GetAllProvinceCasesPaginated_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetAllProvinceCasesPaginated", 10, 0).Return([]models.ProvinceCaseWithDate{}, 0, errors.New("err"))
	_, _, err := svc.GetAllProvinceCasesPaginated(10, 0)
	assert.Error(t, err)
}

func TestCachedCovidService_GetAllProvinceCasesPaginatedSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetAllProvinceCasesPaginatedSorted", 10, 0, sp).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetAllProvinceCasesPaginatedSorted(10, 0, sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetAllProvinceCasesPaginatedSorted(10, 0, sp)
	mockSvc.AssertNumberOfCalls(t, "GetAllProvinceCasesPaginatedSorted", 1)
}

func TestCachedCovidService_GetAllProvinceCasesPaginatedSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetAllProvinceCasesPaginatedSorted", 10, 0, sp).Return([]models.ProvinceCaseWithDate{}, 0, errors.New("err"))
	_, _, err := svc.GetAllProvinceCasesPaginatedSorted(10, 0, sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetAllProvinceCasesByDateRange(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetAllProvinceCasesByDateRange", "2021-01-01", "2021-12-31").Return(expected, nil).Once()

	result, err := svc.GetAllProvinceCasesByDateRange("2021-01-01", "2021-12-31")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetAllProvinceCasesByDateRange("2021-01-01", "2021-12-31")
	mockSvc.AssertNumberOfCalls(t, "GetAllProvinceCasesByDateRange", 1)
}

func TestCachedCovidService_GetAllProvinceCasesByDateRange_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetAllProvinceCasesByDateRange", "2021-01-01", "2021-12-31").Return([]models.ProvinceCaseWithDate{}, errors.New("err"))
	_, err := svc.GetAllProvinceCasesByDateRange("2021-01-01", "2021-12-31")
	assert.Error(t, err)
}

func TestCachedCovidService_GetAllProvinceCasesByDateRangeSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetAllProvinceCasesByDateRangeSorted", "2021-01-01", "2021-12-31", sp).Return(expected, nil).Once()

	result, err := svc.GetAllProvinceCasesByDateRangeSorted("2021-01-01", "2021-12-31", sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)

	svc.GetAllProvinceCasesByDateRangeSorted("2021-01-01", "2021-12-31", sp)
	mockSvc.AssertNumberOfCalls(t, "GetAllProvinceCasesByDateRangeSorted", 1)
}

func TestCachedCovidService_GetAllProvinceCasesByDateRangeSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetAllProvinceCasesByDateRangeSorted", "2021-01-01", "2021-12-31", sp).Return([]models.ProvinceCaseWithDate{}, errors.New("err"))
	_, err := svc.GetAllProvinceCasesByDateRangeSorted("2021-01-01", "2021-12-31", sp)
	assert.Error(t, err)
}

func TestCachedCovidService_GetAllProvinceCasesByDateRangePaginated(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetAllProvinceCasesByDateRangePaginated", "2021-01-01", "2021-12-31", 10, 0).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetAllProvinceCasesByDateRangePaginated("2021-01-01", "2021-12-31", 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetAllProvinceCasesByDateRangePaginated("2021-01-01", "2021-12-31", 10, 0)
	mockSvc.AssertNumberOfCalls(t, "GetAllProvinceCasesByDateRangePaginated", 1)
}

func TestCachedCovidService_GetAllProvinceCasesByDateRangePaginated_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	mockSvc.On("GetAllProvinceCasesByDateRangePaginated", "2021-01-01", "2021-12-31", 10, 0).Return([]models.ProvinceCaseWithDate{}, 0, errors.New("err"))
	_, _, err := svc.GetAllProvinceCasesByDateRangePaginated("2021-01-01", "2021-12-31", 10, 0)
	assert.Error(t, err)
}

func TestCachedCovidService_GetAllProvinceCasesByDateRangePaginatedSorted(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	expected := []models.ProvinceCaseWithDate{{}}
	mockSvc.On("GetAllProvinceCasesByDateRangePaginatedSorted", "2021-01-01", "2021-12-31", 10, 0, sp).Return(expected, 1, nil).Once()

	cases, total, err := svc.GetAllProvinceCasesByDateRangePaginatedSorted("2021-01-01", "2021-12-31", 10, 0, sp)
	assert.NoError(t, err)
	assert.Equal(t, expected, cases)
	assert.Equal(t, 1, total)

	svc.GetAllProvinceCasesByDateRangePaginatedSorted("2021-01-01", "2021-12-31", 10, 0, sp)
	mockSvc.AssertNumberOfCalls(t, "GetAllProvinceCasesByDateRangePaginatedSorted", 1)
}

func TestCachedCovidService_GetAllProvinceCasesByDateRangePaginatedSorted_Error(t *testing.T) {
	mockSvc := new(MockCovidService)
	c := newTestCache()
	svc := NewCachedCovidService(mockSvc, c)

	sp := utils.SortParams{Field: "date", Order: "asc"}
	mockSvc.On("GetAllProvinceCasesByDateRangePaginatedSorted", "2021-01-01", "2021-12-31", 10, 0, sp).Return([]models.ProvinceCaseWithDate{}, 0, errors.New("err"))
	_, _, err := svc.GetAllProvinceCasesByDateRangePaginatedSorted("2021-01-01", "2021-12-31", 10, 0, sp)
	assert.Error(t, err)
}
