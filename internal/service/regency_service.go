package service

import (
	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/repository"
)

// RegencyService handles business logic for regencies
type RegencyService struct {
	regencyRepo     repository.RegencyRepositoryInterface
	regencyCaseRepo repository.RegencyCaseRepositoryInterface
}

// NewRegencyService creates a new RegencyService
func NewRegencyService(regencyRepo repository.RegencyRepositoryInterface, regencyCaseRepo repository.RegencyCaseRepositoryInterface) *RegencyService {
	return &RegencyService{
		regencyRepo:     regencyRepo,
		regencyCaseRepo: regencyCaseRepo,
	}
}

// GetRegencies returns all regencies for SulTeng (province_id=72)
func (s *RegencyService) GetRegencies() ([]models.Regency, error) {
	return s.regencyRepo.GetAll(72)
}

// GetRegencyByID returns a single regency
func (s *RegencyService) GetRegencyByID(id int) (*models.Regency, error) {
	return s.regencyRepo.GetByID(id)
}

// GetRegencyCases returns all cases for a regency
func (s *RegencyService) GetRegencyCases(regencyID int) ([]models.RegencyCase, error) {
	return s.regencyCaseRepo.GetByRegencyID(regencyID)
}

// GetLatestRegencyCases returns latest case for each regency
func (s *RegencyService) GetLatestRegencyCases() ([]models.RegencyCase, error) {
	return s.regencyCaseRepo.GetLatestByProvinceID(72)
}
