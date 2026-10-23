package service

import (
	"fmt"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/cache"
)

// cachedContentService wraps a ContentServiceInterface with in-memory caching (1h).
type cachedContentService struct {
	svc   ContentServiceInterface
	cache *cache.Cache
}

// NewCachedContentService returns a ContentServiceInterface backed by an in-memory cache.
func NewCachedContentService(svc ContentServiceInterface, c *cache.Cache) ContentServiceInterface {
	return &cachedContentService{svc: svc, cache: c}
}

type pageResult[T any] struct {
	items []T
	total int
}

// cachedPage returns the cached page for key or loads and caches it.
func cachedPage[T any](c *cache.Cache, key string, load func() ([]T, int, error)) ([]T, int, error) {
	if v, ok := c.Get(key); ok {
		r := v.(pageResult[T])
		return r.items, r.total, nil
	}
	items, total, err := load()
	if err != nil {
		return nil, 0, err
	}
	c.Set(key, pageResult[T]{items, total}, ttlDefault)
	return items, total, nil
}

func (s *cachedContentService) GetBanners(includeInactive bool, limit, offset int) ([]models.Banner, int, error) {
	key := fmt.Sprintf("content:banners:%t:%d:%d", includeInactive, limit, offset)
	return cachedPage(s.cache, key, func() ([]models.Banner, int, error) {
		return s.svc.GetBanners(includeInactive, limit, offset)
	})
}

func (s *cachedContentService) GetInfographics(limit, offset int) ([]models.Infographic, int, error) {
	key := fmt.Sprintf("content:infographics:%d:%d", limit, offset)
	return cachedPage(s.cache, key, func() ([]models.Infographic, int, error) {
		return s.svc.GetInfographics(limit, offset)
	})
}

func (s *cachedContentService) GetDonations(limit, offset int) ([]models.Donation, int, error) {
	key := fmt.Sprintf("content:donations:%d:%d", limit, offset)
	return cachedPage(s.cache, key, func() ([]models.Donation, int, error) {
		return s.svc.GetDonations(limit, offset)
	})
}

func (s *cachedContentService) GetTelemedicines(limit, offset int) ([]models.Telemedicine, int, error) {
	key := fmt.Sprintf("content:telemedicines:%d:%d", limit, offset)
	return cachedPage(s.cache, key, func() ([]models.Telemedicine, int, error) {
		return s.svc.GetTelemedicines(limit, offset)
	})
}

func (s *cachedContentService) GetPartners(inFooter *bool, limit, offset int) ([]models.Partner, int, error) {
	f := "any"
	if inFooter != nil {
		f = fmt.Sprintf("%t", *inFooter)
	}
	key := fmt.Sprintf("content:partners:%s:%d:%d", f, limit, offset)
	return cachedPage(s.cache, key, func() ([]models.Partner, int, error) {
		return s.svc.GetPartners(inFooter, limit, offset)
	})
}

func (s *cachedContentService) GetContacts(ownerType string, ownerID int64, limit, offset int) ([]models.ContactEntry, int, error) {
	key := fmt.Sprintf("content:contacts:%s:%d:%d:%d", ownerType, ownerID, limit, offset)
	return cachedPage(s.cache, key, func() ([]models.ContactEntry, int, error) {
		return s.svc.GetContacts(ownerType, ownerID, limit, offset)
	})
}
