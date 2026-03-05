package service

import (
	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/repository"
)

// HospitalService handles business logic for hospitals
type HospitalService struct {
	hospitalRepo *repository.HospitalRepository
}

// NewHospitalService creates a new HospitalService
func NewHospitalService(hospitalRepo *repository.HospitalRepository) *HospitalService {
	return &HospitalService{hospitalRepo: hospitalRepo}
}

// GetHospitals returns all hospitals in SulTeng (province_id=72)
func (s *HospitalService) GetHospitals() ([]models.Hospital, error) {
	return s.hospitalRepo.GetAll(72)
}

// GetHospitalByCode returns a single hospital by code
func (s *HospitalService) GetHospitalByCode(code string) (*models.Hospital, error) {
	return s.hospitalRepo.GetByCode(code)
}
