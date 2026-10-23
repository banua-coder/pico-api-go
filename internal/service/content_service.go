package service

import (
	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/repository"
)

// ContentService serves read-only site content (banners, infographics, ...).
type ContentService struct {
	repo repository.ContentRepositoryInterface
}

// NewContentService creates a new ContentService.
func NewContentService(repo repository.ContentRepositoryInterface) *ContentService {
	return &ContentService{repo: repo}
}

func (s *ContentService) GetBanners(includeInactive bool, limit, offset int) ([]models.Banner, int, error) {
	return s.repo.GetBanners(includeInactive, limit, offset)
}

func (s *ContentService) GetInfographics(limit, offset int) ([]models.Infographic, int, error) {
	return s.repo.GetInfographics(limit, offset)
}

func (s *ContentService) GetDonations(limit, offset int) ([]models.Donation, int, error) {
	return s.repo.GetDonations(limit, offset)
}

func (s *ContentService) GetTelemedicines(limit, offset int) ([]models.Telemedicine, int, error) {
	return s.repo.GetTelemedicines(limit, offset)
}

func (s *ContentService) GetPartners(inFooter *bool, limit, offset int) ([]models.Partner, int, error) {
	return s.repo.GetPartners(inFooter, limit, offset)
}

func (s *ContentService) GetContacts(ownerType string, ownerID int64, limit, offset int) ([]models.ContactEntry, int, error) {
	return s.repo.GetContacts(ownerType, ownerID, limit, offset)
}
