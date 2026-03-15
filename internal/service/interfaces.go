package service

import "github.com/banua-coder/pico-api-go/internal/models"

// RegencyServiceInterface defines the contract for regency operations
type RegencyServiceInterface interface {
	GetRegencies() ([]models.Regency, error)
	GetRegencyByID(id int) (*models.Regency, error)
	GetRegencyCases(regencyID int) ([]models.RegencyCase, error)
	GetLatestRegencyCases() ([]models.RegencyCase, error)
}

// HospitalServiceInterface defines the contract for hospital operations
type HospitalServiceInterface interface {
	GetHospitals() ([]models.Hospital, error)
	GetHospitalByCode(code string) (*models.Hospital, error)
}

// TaskForceServiceInterface defines the contract for task force operations
type TaskForceServiceInterface interface {
	GetTaskForces() ([]models.TaskForceByRegency, error)
}

// VaccinationServiceInterface defines the contract for vaccination operations
type VaccinationServiceInterface interface {
	GetNationalVaccinations() ([]models.NationalVaccine, error)
	GetProvinceVaccinations() ([]models.ProvinceVaccine, error)
	GetVaccineLocations() ([]models.VaccineLocation, error)
}

// ProvinceStatsServiceInterface defines the contract for province stats operations
type ProvinceStatsServiceInterface interface {
	GetGenderCases() ([]models.ProvinceGenderCase, error)
	GetLatestGenderCase() (*models.ProvinceGenderCase, error)
	GetTests() ([]models.ProvinceTest, error)
	GetTestTypes() ([]models.TestType, error)
}
