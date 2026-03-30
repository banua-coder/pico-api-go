package service

import (
	"fmt"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/cache"
	"github.com/banua-coder/pico-api-go/pkg/utils"
)

const (
	ttlLatest     = 15 * time.Minute
	ttlHistorical = 24 * time.Hour
	ttlDefault    = time.Hour
)

// CacheInvalidator is the interface for cache invalidation.
type CacheInvalidator interface {
	Clear()
}

// cachedCovidService wraps a CovidService with in-memory caching.
type cachedCovidService struct {
	svc   CovidService
	cache *cache.Cache
}

// NewCachedCovidService returns a CovidService backed by an in-memory cache.
func NewCachedCovidService(svc CovidService, c *cache.Cache) CovidService {
	return &cachedCovidService{svc: svc, cache: c}
}

// -- helper ----------------------------------------------------------

func (s *cachedCovidService) getOrSet(key string, ttl time.Duration, fn func() (interface{}, error)) (interface{}, error) {
	if v, ok := s.cache.Get(key); ok {
		return v, nil
	}
	v, err := fn()
	if err != nil {
		return nil, err
	}
	s.cache.Set(key, v, ttl)
	return v, nil
}

// -- national cases --------------------------------------------------

func (s *cachedCovidService) GetNationalCases() ([]models.NationalCase, error) {
	v, err := s.getOrSet("national:all", ttlDefault, func() (interface{}, error) {
		return s.svc.GetNationalCases()
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.NationalCase), nil
}

func (s *cachedCovidService) GetNationalCasesSorted(sortParams utils.SortParams) ([]models.NationalCase, error) {
	key := fmt.Sprintf("national:all:sort:%s:%s", sortParams.Field, sortParams.Order)
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		return s.svc.GetNationalCasesSorted(sortParams)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.NationalCase), nil
}

func (s *cachedCovidService) GetNationalCasesPaginated(limit, offset int) ([]models.NationalCase, int, error) {
	key := fmt.Sprintf("national:all:page:%d:%d", limit, offset)
	type result struct {
		cases []models.NationalCase
		total int
	}
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		cases, total, err := s.svc.GetNationalCasesPaginated(limit, offset)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetNationalCasesPaginatedSorted(limit, offset int, sortParams utils.SortParams) ([]models.NationalCase, int, error) {
	key := fmt.Sprintf("national:all:page:%d:%d:sort:%s:%s", limit, offset, sortParams.Field, sortParams.Order)
	type result struct {
		cases []models.NationalCase
		total int
	}
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		cases, total, err := s.svc.GetNationalCasesPaginatedSorted(limit, offset, sortParams)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetNationalCasesByDateRange(startDate, endDate string) ([]models.NationalCase, error) {
	key := fmt.Sprintf("national:date:%s:%s", startDate, endDate)
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		return s.svc.GetNationalCasesByDateRange(startDate, endDate)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.NationalCase), nil
}

func (s *cachedCovidService) GetNationalCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.NationalCase, error) {
	key := fmt.Sprintf("national:date:%s:%s:sort:%s:%s", startDate, endDate, sortParams.Field, sortParams.Order)
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		return s.svc.GetNationalCasesByDateRangeSorted(startDate, endDate, sortParams)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.NationalCase), nil
}

func (s *cachedCovidService) GetNationalCasesByDateRangePaginated(startDate, endDate string, limit, offset int) ([]models.NationalCase, int, error) {
	key := fmt.Sprintf("national:date:%s:%s:page:%d:%d", startDate, endDate, limit, offset)
	type result struct {
		cases []models.NationalCase
		total int
	}
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		cases, total, err := s.svc.GetNationalCasesByDateRangePaginated(startDate, endDate, limit, offset)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetNationalCasesByDateRangePaginatedSorted(startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.NationalCase, int, error) {
	key := fmt.Sprintf("national:date:%s:%s:page:%d:%d:sort:%s:%s", startDate, endDate, limit, offset, sortParams.Field, sortParams.Order)
	type result struct {
		cases []models.NationalCase
		total int
	}
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		cases, total, err := s.svc.GetNationalCasesByDateRangePaginatedSorted(startDate, endDate, limit, offset, sortParams)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetLatestNationalCase() (*models.NationalCase, error) {
	v, err := s.getOrSet("national:latest", ttlLatest, func() (interface{}, error) {
		return s.svc.GetLatestNationalCase()
	})
	if err != nil {
		return nil, err
	}
	return v.(*models.NationalCase), nil
}

func (s *cachedCovidService) GetNationalCaseByDay(day int64) (*models.NationalCase, error) {
	key := fmt.Sprintf("national:day:%d", day)
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		return s.svc.GetNationalCaseByDay(day)
	})
	if err != nil {
		return nil, err
	}
	return v.(*models.NationalCase), nil
}

// -- provinces -------------------------------------------------------

func (s *cachedCovidService) GetProvinces() ([]models.Province, error) {
	v, err := s.getOrSet("province:all", ttlDefault, func() (interface{}, error) {
		return s.svc.GetProvinces()
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.Province), nil
}

func (s *cachedCovidService) GetProvinceByID(id string) (*models.Province, error) {
	key := fmt.Sprintf("province:%s", id)
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		return s.svc.GetProvinceByID(id)
	})
	if err != nil {
		return nil, err
	}
	return v.(*models.Province), nil
}

func (s *cachedCovidService) GetProvincesWithLatestCase() ([]models.ProvinceWithLatestCase, error) {
	v, err := s.getOrSet("province:all:with_latest", ttlLatest, func() (interface{}, error) {
		return s.svc.GetProvincesWithLatestCase()
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.ProvinceWithLatestCase), nil
}

func (s *cachedCovidService) GetProvinceCases(provinceID string) ([]models.ProvinceCaseWithDate, error) {
	key := fmt.Sprintf("province:%s:cases:all", provinceID)
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		return s.svc.GetProvinceCases(provinceID)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.ProvinceCaseWithDate), nil
}

func (s *cachedCovidService) GetProvinceCasesSorted(provinceID string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	key := fmt.Sprintf("province:%s:cases:all:sort:%s:%s", provinceID, sortParams.Field, sortParams.Order)
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		return s.svc.GetProvinceCasesSorted(provinceID, sortParams)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.ProvinceCaseWithDate), nil
}

func (s *cachedCovidService) GetProvinceCasesPaginated(provinceID string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	key := fmt.Sprintf("province:%s:cases:page:%d:%d", provinceID, limit, offset)
	type result struct {
		cases []models.ProvinceCaseWithDate
		total int
	}
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		cases, total, err := s.svc.GetProvinceCasesPaginated(provinceID, limit, offset)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetProvinceCasesPaginatedSorted(provinceID string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	key := fmt.Sprintf("province:%s:cases:page:%d:%d:sort:%s:%s", provinceID, limit, offset, sortParams.Field, sortParams.Order)
	type result struct {
		cases []models.ProvinceCaseWithDate
		total int
	}
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		cases, total, err := s.svc.GetProvinceCasesPaginatedSorted(provinceID, limit, offset, sortParams)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetProvinceCasesByDateRange(provinceID, startDate, endDate string) ([]models.ProvinceCaseWithDate, error) {
	key := fmt.Sprintf("province:%s:cases:date:%s:%s", provinceID, startDate, endDate)
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		return s.svc.GetProvinceCasesByDateRange(provinceID, startDate, endDate)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.ProvinceCaseWithDate), nil
}

func (s *cachedCovidService) GetProvinceCasesByDateRangeSorted(provinceID, startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	key := fmt.Sprintf("province:%s:cases:date:%s:%s:sort:%s:%s", provinceID, startDate, endDate, sortParams.Field, sortParams.Order)
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		return s.svc.GetProvinceCasesByDateRangeSorted(provinceID, startDate, endDate, sortParams)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.ProvinceCaseWithDate), nil
}

func (s *cachedCovidService) GetProvinceCasesByDateRangePaginated(provinceID, startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	key := fmt.Sprintf("province:%s:cases:date:%s:%s:page:%d:%d", provinceID, startDate, endDate, limit, offset)
	type result struct {
		cases []models.ProvinceCaseWithDate
		total int
	}
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		cases, total, err := s.svc.GetProvinceCasesByDateRangePaginated(provinceID, startDate, endDate, limit, offset)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetProvinceCasesByDateRangePaginatedSorted(provinceID, startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	key := fmt.Sprintf("province:%s:cases:date:%s:%s:page:%d:%d:sort:%s:%s", provinceID, startDate, endDate, limit, offset, sortParams.Field, sortParams.Order)
	type result struct {
		cases []models.ProvinceCaseWithDate
		total int
	}
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		cases, total, err := s.svc.GetProvinceCasesByDateRangePaginatedSorted(provinceID, startDate, endDate, limit, offset, sortParams)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

// -- all province cases ----------------------------------------------

func (s *cachedCovidService) GetAllProvinceCases() ([]models.ProvinceCaseWithDate, error) {
	v, err := s.getOrSet("province:cases:all", ttlDefault, func() (interface{}, error) {
		return s.svc.GetAllProvinceCases()
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.ProvinceCaseWithDate), nil
}

func (s *cachedCovidService) GetAllProvinceCasesSorted(sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	key := fmt.Sprintf("province:cases:all:sort:%s:%s", sortParams.Field, sortParams.Order)
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		return s.svc.GetAllProvinceCasesSorted(sortParams)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.ProvinceCaseWithDate), nil
}

func (s *cachedCovidService) GetAllProvinceCasesPaginated(limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	key := fmt.Sprintf("province:cases:all:page:%d:%d", limit, offset)
	type result struct {
		cases []models.ProvinceCaseWithDate
		total int
	}
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		cases, total, err := s.svc.GetAllProvinceCasesPaginated(limit, offset)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetAllProvinceCasesPaginatedSorted(limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	key := fmt.Sprintf("province:cases:all:page:%d:%d:sort:%s:%s", limit, offset, sortParams.Field, sortParams.Order)
	type result struct {
		cases []models.ProvinceCaseWithDate
		total int
	}
	v, err := s.getOrSet(key, ttlDefault, func() (interface{}, error) {
		cases, total, err := s.svc.GetAllProvinceCasesPaginatedSorted(limit, offset, sortParams)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetAllProvinceCasesByDateRange(startDate, endDate string) ([]models.ProvinceCaseWithDate, error) {
	key := fmt.Sprintf("province:cases:date:%s:%s", startDate, endDate)
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		return s.svc.GetAllProvinceCasesByDateRange(startDate, endDate)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.ProvinceCaseWithDate), nil
}

func (s *cachedCovidService) GetAllProvinceCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	key := fmt.Sprintf("province:cases:date:%s:%s:sort:%s:%s", startDate, endDate, sortParams.Field, sortParams.Order)
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		return s.svc.GetAllProvinceCasesByDateRangeSorted(startDate, endDate, sortParams)
	})
	if err != nil {
		return nil, err
	}
	return v.([]models.ProvinceCaseWithDate), nil
}

func (s *cachedCovidService) GetAllProvinceCasesByDateRangePaginated(startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	key := fmt.Sprintf("province:cases:date:%s:%s:page:%d:%d", startDate, endDate, limit, offset)
	type result struct {
		cases []models.ProvinceCaseWithDate
		total int
	}
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		cases, total, err := s.svc.GetAllProvinceCasesByDateRangePaginated(startDate, endDate, limit, offset)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}

func (s *cachedCovidService) GetAllProvinceCasesByDateRangePaginatedSorted(startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	key := fmt.Sprintf("province:cases:date:%s:%s:page:%d:%d:sort:%s:%s", startDate, endDate, limit, offset, sortParams.Field, sortParams.Order)
	type result struct {
		cases []models.ProvinceCaseWithDate
		total int
	}
	v, err := s.getOrSet(key, ttlHistorical, func() (interface{}, error) {
		cases, total, err := s.svc.GetAllProvinceCasesByDateRangePaginatedSorted(startDate, endDate, limit, offset, sortParams)
		return result{cases, total}, err
	})
	if err != nil {
		return nil, 0, err
	}
	r := v.(result)
	return r.cases, r.total, nil
}
