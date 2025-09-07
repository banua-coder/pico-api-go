package service

import (
	"fmt"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/repository"
)

type CovidService interface {
	GetNationalCases() ([]models.NationalCase, error)
	GetNationalCasesByDateRange(startDate, endDate string) ([]models.NationalCase, error)
	GetLatestNationalCase() (*models.NationalCase, error)
	GetProvinces() ([]models.Province, error)
	GetProvinceCases(provinceID string) ([]models.ProvinceCaseWithDate, error)
	GetProvinceCasesByDateRange(provinceID, startDate, endDate string) ([]models.ProvinceCaseWithDate, error)
	GetAllProvinceCases() ([]models.ProvinceCaseWithDate, error)
	GetAllProvinceCasesByDateRange(startDate, endDate string) ([]models.ProvinceCaseWithDate, error)
}

type covidService struct {
	nationalCaseRepo  repository.NationalCaseRepository
	provinceRepo      repository.ProvinceRepository
	provinceCaseRepo  repository.ProvinceCaseRepository
}

func NewCovidService(
	nationalCaseRepo repository.NationalCaseRepository,
	provinceRepo repository.ProvinceRepository,
	provinceCaseRepo repository.ProvinceCaseRepository,
) CovidService {
	return &covidService{
		nationalCaseRepo: nationalCaseRepo,
		provinceRepo:     provinceRepo,
		provinceCaseRepo: provinceCaseRepo,
	}
}

func (s *covidService) GetNationalCases() ([]models.NationalCase, error) {
	cases, err := s.nationalCaseRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get national cases: %w", err)
	}
	return cases, nil
}

func (s *covidService) GetNationalCasesByDateRange(startDate, endDate string) ([]models.NationalCase, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, err := s.nationalCaseRepo.GetByDateRange(start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to get national cases by date range: %w", err)
	}
	return cases, nil
}

func (s *covidService) GetLatestNationalCase() (*models.NationalCase, error) {
	nationalCase, err := s.nationalCaseRepo.GetLatest()
	if err != nil {
		return nil, fmt.Errorf("failed to get latest national case: %w", err)
	}
	return nationalCase, nil
}

func (s *covidService) GetProvinces() ([]models.Province, error) {
	provinces, err := s.provinceRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get provinces: %w", err)
	}
	return provinces, nil
}

func (s *covidService) GetProvinceCases(provinceID string) ([]models.ProvinceCaseWithDate, error) {
	cases, err := s.provinceCaseRepo.GetByProvinceID(provinceID)
	if err != nil {
		return nil, fmt.Errorf("failed to get province cases: %w", err)
	}
	return cases, nil
}

func (s *covidService) GetProvinceCasesByDateRange(provinceID, startDate, endDate string) ([]models.ProvinceCaseWithDate, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, err := s.provinceCaseRepo.GetByProvinceIDAndDateRange(provinceID, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to get province cases by date range: %w", err)
	}
	return cases, nil
}

func (s *covidService) GetAllProvinceCases() ([]models.ProvinceCaseWithDate, error) {
	cases, err := s.provinceCaseRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all province cases: %w", err)
	}
	return cases, nil
}

func (s *covidService) GetAllProvinceCasesByDateRange(startDate, endDate string) ([]models.ProvinceCaseWithDate, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, err := s.provinceCaseRepo.GetByDateRange(start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to get all province cases by date range: %w", err)
	}
	return cases, nil
}