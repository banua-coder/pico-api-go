package service

import (
	"fmt"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/cache"
)

// cachedRegencyService wraps a RegencyServiceInterface with in-memory caching.
type cachedRegencyService struct {
	svc   RegencyServiceInterface
	cache *cache.Cache
}

// NewCachedRegencyService returns a RegencyServiceInterface backed by an in-memory cache.
func NewCachedRegencyService(svc RegencyServiceInterface, c *cache.Cache) RegencyServiceInterface {
	return &cachedRegencyService{svc: svc, cache: c}
}

func (s *cachedRegencyService) GetRegencies() ([]models.Regency, error) {
	const key = "regency:all"
	if v, ok := s.cache.Get(key); ok {
		return v.([]models.Regency), nil
	}
	result, err := s.svc.GetRegencies()
	if err != nil {
		return nil, err
	}
	s.cache.Set(key, result, ttlDefault)
	return result, nil
}

func (s *cachedRegencyService) GetRegenciesPaginated(limit, offset int) ([]models.Regency, int, error) {
	key := fmt.Sprintf("regency:all:page:%d:%d", limit, offset)
	type res struct {
		items []models.Regency
		total int
	}
	if v, ok := s.cache.Get(key); ok {
		r := v.(res)
		return r.items, r.total, nil
	}
	items, total, err := s.svc.GetRegenciesPaginated(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	s.cache.Set(key, res{items, total}, ttlDefault)
	return items, total, nil
}

func (s *cachedRegencyService) GetRegencyByID(id int) (*models.Regency, error) {
	key := fmt.Sprintf("regency:%d", id)
	if v, ok := s.cache.Get(key); ok {
		return v.(*models.Regency), nil
	}
	result, err := s.svc.GetRegencyByID(id)
	if err != nil {
		return nil, err
	}
	s.cache.Set(key, result, ttlDefault)
	return result, nil
}

func (s *cachedRegencyService) GetRegencyCases(regencyID int) ([]models.RegencyCase, error) {
	key := fmt.Sprintf("regency:%d:cases:all", regencyID)
	if v, ok := s.cache.Get(key); ok {
		return v.([]models.RegencyCase), nil
	}
	result, err := s.svc.GetRegencyCases(regencyID)
	if err != nil {
		return nil, err
	}
	s.cache.Set(key, result, ttlDefault)
	return result, nil
}

func (s *cachedRegencyService) GetLatestRegencyCases() ([]models.RegencyCase, error) {
	const key = "regency:cases:latest"
	if v, ok := s.cache.Get(key); ok {
		return v.([]models.RegencyCase), nil
	}
	result, err := s.svc.GetLatestRegencyCases()
	if err != nil {
		return nil, err
	}
	s.cache.Set(key, result, ttlLatest)
	return result, nil
}
