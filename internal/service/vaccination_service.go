package service

import (
	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/repository"
)

type VaccinationService struct {
	vaccinationRepo repository.VaccinationRepositoryInterface
}

func NewVaccinationService(vaccinationRepo repository.VaccinationRepositoryInterface) *VaccinationService {
	return &VaccinationService{vaccinationRepo: vaccinationRepo}
}

func (s *VaccinationService) GetNationalVaccinations() ([]models.NationalVaccine, error) {
	return s.vaccinationRepo.GetNationalVaccinations()
}

func (s *VaccinationService) GetProvinceVaccinations() ([]models.ProvinceVaccine, error) {
	return s.vaccinationRepo.GetProvinceVaccinations(72)
}

func (s *VaccinationService) GetVaccineLocations() ([]models.VaccineLocation, error) {
	return s.vaccinationRepo.GetVaccineLocations(72)
}
