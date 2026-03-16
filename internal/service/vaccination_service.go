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

func (s *VaccinationService) GetNationalVaccinationsPaginated(limit, offset int) ([]models.NationalVaccine, int, error) {
	return s.vaccinationRepo.GetNationalVaccinationsPaginated(limit, offset)
}

func (s *VaccinationService) GetProvinceVaccinations() ([]models.ProvinceVaccine, error) {
	return s.vaccinationRepo.GetProvinceVaccinations(72)
}

func (s *VaccinationService) GetProvinceVaccinationsPaginated(limit, offset int) ([]models.ProvinceVaccine, int, error) {
	return s.vaccinationRepo.GetProvinceVaccinationsPaginated(72, limit, offset)
}

func (s *VaccinationService) GetVaccineLocations() ([]models.VaccineLocation, error) {
	return s.vaccinationRepo.GetVaccineLocations(72)
}

func (s *VaccinationService) GetVaccineLocationsPaginated(limit, offset int) ([]models.VaccineLocation, int, error) {
	return s.vaccinationRepo.GetVaccineLocationsPaginated(72, limit, offset)
}
