package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/cache"
	"github.com/banua-coder/pico-api-go/pkg/utils"
	"github.com/redis/go-redis/v9"
)

// cachedCovidService wraps CovidService with caching.
type cachedCovidService struct {
	inner CovidService
	cache cache.Cache
	ttl   time.Duration
}

// NewCachedCovidService wraps the given CovidService with a caching layer.
func NewCachedCovidService(inner CovidService, c cache.Cache, ttl time.Duration) CovidService {
	return &cachedCovidService{inner: inner, cache: c, ttl: ttl}
}

func (s *cachedCovidService) getOrSet(key string, dest interface{}, fetch func() (interface{}, error)) error {
	ctx := context.Background()

	// Try cache first
	raw, err := s.cache.Get(ctx, key)
	if err == nil {
		if jsonErr := json.Unmarshal([]byte(raw), dest); jsonErr == nil {
			return nil
		}
	}

	// Cache miss — fetch from DB
	result, err := fetch()
	if err != nil {
		return err
	}

	// Store in cache (best-effort)
	b, jsonErr := json.Marshal(result)
	if jsonErr == nil {
		if setErr := s.cache.Set(ctx, key, string(b), s.ttl); setErr != nil {
			log.Printf("cache set error for key %s: %v", key, setErr)
		}
	}

	// Copy result into dest using JSON round-trip
	b2, _ := json.Marshal(result)
	return json.Unmarshal(b2, dest)
}

func isCacheMiss(err error) bool {
	if err == nil {
		return false
	}
	return err == redis.Nil || err.Error() == "cache miss"
}

// ----- Cached methods -----

func (s *cachedCovidService) GetNationalCases() ([]models.NationalCase, error) {
	var result []models.NationalCase
	err := s.getOrSet(cache.Key("national", "all"), &result, func() (interface{}, error) {
		return s.inner.GetNationalCases()
	})
	return result, err
}

func (s *cachedCovidService) GetLatestNationalCase() (*models.NationalCase, error) {
	var result models.NationalCase
	err := s.getOrSet(cache.Key("national", "latest"), &result, func() (interface{}, error) {
		return s.inner.GetLatestNationalCase()
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *cachedCovidService) GetProvinces() ([]models.Province, error) {
	var result []models.Province
	err := s.getOrSet(cache.Key("provinces", "all"), &result, func() (interface{}, error) {
		return s.inner.GetProvinces()
	})
	return result, err
}

func (s *cachedCovidService) GetProvincesWithLatestCase() ([]models.ProvinceWithLatestCase, error) {
	var result []models.ProvinceWithLatestCase
	err := s.getOrSet(cache.Key("provinces", "latest-cases"), &result, func() (interface{}, error) {
		return s.inner.GetProvincesWithLatestCase()
	})
	return result, err
}

func (s *cachedCovidService) GetProvinceByID(id string) (*models.Province, error) {
	var result models.Province
	err := s.getOrSet(cache.Key("province", id), &result, func() (interface{}, error) {
		return s.inner.GetProvinceByID(id)
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *cachedCovidService) GetAllProvinceCases() ([]models.ProvinceCaseWithDate, error) {
	var result []models.ProvinceCaseWithDate
	err := s.getOrSet(cache.Key("province-cases", "all"), &result, func() (interface{}, error) {
		return s.inner.GetAllProvinceCases()
	})
	return result, err
}

// ----- Pass-through methods (not cached — too many variants / params) -----

func (s *cachedCovidService) GetNationalCasesSorted(sortParams utils.SortParams) ([]models.NationalCase, error) {
	return s.inner.GetNationalCasesSorted(sortParams)
}

func (s *cachedCovidService) GetNationalCasesPaginated(limit, offset int) ([]models.NationalCase, int, error) {
	return s.inner.GetNationalCasesPaginated(limit, offset)
}

func (s *cachedCovidService) GetNationalCasesPaginatedSorted(limit, offset int, sortParams utils.SortParams) ([]models.NationalCase, int, error) {
	return s.inner.GetNationalCasesPaginatedSorted(limit, offset, sortParams)
}

func (s *cachedCovidService) GetNationalCasesByDateRange(startDate, endDate string) ([]models.NationalCase, error) {
	return s.inner.GetNationalCasesByDateRange(startDate, endDate)
}

func (s *cachedCovidService) GetNationalCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.NationalCase, error) {
	return s.inner.GetNationalCasesByDateRangeSorted(startDate, endDate, sortParams)
}

func (s *cachedCovidService) GetNationalCasesByDateRangePaginated(startDate, endDate string, limit, offset int) ([]models.NationalCase, int, error) {
	return s.inner.GetNationalCasesByDateRangePaginated(startDate, endDate, limit, offset)
}

func (s *cachedCovidService) GetNationalCasesByDateRangePaginatedSorted(startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.NationalCase, int, error) {
	return s.inner.GetNationalCasesByDateRangePaginatedSorted(startDate, endDate, limit, offset, sortParams)
}

func (s *cachedCovidService) GetNationalCaseByDay(day int64) (*models.NationalCase, error) {
	return s.inner.GetNationalCaseByDay(day)
}

func (s *cachedCovidService) GetProvinceCases(provinceID string) ([]models.ProvinceCaseWithDate, error) {
	return s.inner.GetProvinceCases(provinceID)
}

func (s *cachedCovidService) GetProvinceCasesSorted(provinceID string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	return s.inner.GetProvinceCasesSorted(provinceID, sortParams)
}

func (s *cachedCovidService) GetProvinceCasesPaginated(provinceID string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	return s.inner.GetProvinceCasesPaginated(provinceID, limit, offset)
}

func (s *cachedCovidService) GetProvinceCasesPaginatedSorted(provinceID string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	return s.inner.GetProvinceCasesPaginatedSorted(provinceID, limit, offset, sortParams)
}

func (s *cachedCovidService) GetProvinceCasesByDateRange(provinceID, startDate, endDate string) ([]models.ProvinceCaseWithDate, error) {
	return s.inner.GetProvinceCasesByDateRange(provinceID, startDate, endDate)
}

func (s *cachedCovidService) GetProvinceCasesByDateRangeSorted(provinceID, startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	return s.inner.GetProvinceCasesByDateRangeSorted(provinceID, startDate, endDate, sortParams)
}

func (s *cachedCovidService) GetProvinceCasesByDateRangePaginated(provinceID, startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	return s.inner.GetProvinceCasesByDateRangePaginated(provinceID, startDate, endDate, limit, offset)
}

func (s *cachedCovidService) GetProvinceCasesByDateRangePaginatedSorted(provinceID, startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	return s.inner.GetProvinceCasesByDateRangePaginatedSorted(provinceID, startDate, endDate, limit, offset, sortParams)
}

func (s *cachedCovidService) GetAllProvinceCasesSorted(sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	return s.inner.GetAllProvinceCasesSorted(sortParams)
}

func (s *cachedCovidService) GetAllProvinceCasesPaginated(limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	return s.inner.GetAllProvinceCasesPaginated(limit, offset)
}

func (s *cachedCovidService) GetAllProvinceCasesPaginatedSorted(limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	return s.inner.GetAllProvinceCasesPaginatedSorted(limit, offset, sortParams)
}

func (s *cachedCovidService) GetAllProvinceCasesByDateRange(startDate, endDate string) ([]models.ProvinceCaseWithDate, error) {
	return s.inner.GetAllProvinceCasesByDateRange(startDate, endDate)
}

func (s *cachedCovidService) GetAllProvinceCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	return s.inner.GetAllProvinceCasesByDateRangeSorted(startDate, endDate, sortParams)
}

func (s *cachedCovidService) GetAllProvinceCasesByDateRangePaginated(startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	return s.inner.GetAllProvinceCasesByDateRangePaginated(startDate, endDate, limit, offset)
}

func (s *cachedCovidService) GetAllProvinceCasesByDateRangePaginatedSorted(startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	return s.inner.GetAllProvinceCasesByDateRangePaginatedSorted(startDate, endDate, limit, offset, sortParams)
}

// Ensure unused import doesn't error — isCacheMiss may be used in future.
var _ = fmt.Sprintf
var _ = isCacheMiss
