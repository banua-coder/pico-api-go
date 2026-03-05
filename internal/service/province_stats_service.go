package service

import (
	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/repository"
)

type ProvinceStatsService struct {
	repo *repository.ProvinceStatsRepository
}

func NewProvinceStatsService(repo *repository.ProvinceStatsRepository) *ProvinceStatsService {
	return &ProvinceStatsService{repo: repo}
}

func (s *ProvinceStatsService) GetGenderCases() ([]models.ProvinceGenderCase, error) {
	return s.repo.GetGenderCases(72)
}

func (s *ProvinceStatsService) GetLatestGenderCase() (*models.ProvinceGenderCase, error) {
	return s.repo.GetLatestGenderCase(72)
}

func (s *ProvinceStatsService) GetTests() ([]models.ProvinceTest, error) {
	return s.repo.GetTests(72)
}

func (s *ProvinceStatsService) GetTestTypes() ([]models.TestType, error) {
	return s.repo.GetTestTypes()
}
