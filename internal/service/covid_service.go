package service

import (
	"fmt"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/repository"
	"github.com/banua-coder/pico-api-go/pkg/utils"
)

type CovidService interface {
	GetNationalCases() ([]models.NationalCase, error)
	GetNationalCasesSorted(sortParams utils.SortParams) ([]models.NationalCase, error)
	GetNationalCasesByDateRange(startDate, endDate string) ([]models.NationalCase, error)
	GetNationalCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.NationalCase, error)
	GetLatestNationalCase() (*models.NationalCase, error)
	GetProvinces() ([]models.Province, error)
	GetProvincesWithLatestCase() ([]models.ProvinceWithLatestCase, error)
	GetProvinceCases(provinceID string) ([]models.ProvinceCaseWithDate, error)
	GetProvinceCasesSorted(provinceID string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error)
	GetProvinceCasesPaginated(provinceID string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error)
	GetProvinceCasesPaginatedSorted(provinceID string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error)
	GetProvinceCasesByDateRange(provinceID, startDate, endDate string) ([]models.ProvinceCaseWithDate, error)
	GetProvinceCasesByDateRangeSorted(provinceID, startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error)
	GetProvinceCasesByDateRangePaginated(provinceID, startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error)
	GetProvinceCasesByDateRangePaginatedSorted(provinceID, startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error)
	GetAllProvinceCases() ([]models.ProvinceCaseWithDate, error)
	GetAllProvinceCasesSorted(sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error)
	GetAllProvinceCasesPaginated(limit, offset int) ([]models.ProvinceCaseWithDate, int, error)
	GetAllProvinceCasesPaginatedSorted(limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error)
	GetAllProvinceCasesByDateRange(startDate, endDate string) ([]models.ProvinceCaseWithDate, error)
	GetAllProvinceCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error)
	GetAllProvinceCasesByDateRangePaginated(startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error)
	GetAllProvinceCasesByDateRangePaginatedSorted(startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error)
}

type covidService struct {
	nationalCaseRepo repository.NationalCaseRepository
	provinceRepo     repository.ProvinceRepository
	provinceCaseRepo repository.ProvinceCaseRepository
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

func (s *covidService) GetNationalCasesSorted(sortParams utils.SortParams) ([]models.NationalCase, error) {
	cases, err := s.nationalCaseRepo.GetAllSorted(sortParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get sorted national cases: %w", err)
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

func (s *covidService) GetNationalCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.NationalCase, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, err := s.nationalCaseRepo.GetByDateRangeSorted(start, end, sortParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get sorted national cases by date range: %w", err)
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

func (s *covidService) GetProvincesWithLatestCase() ([]models.ProvinceWithLatestCase, error) {
	provinces, err := s.provinceRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get provinces: %w", err)
	}

	result := make([]models.ProvinceWithLatestCase, len(provinces))

	for i, province := range provinces {
		result[i] = models.ProvinceWithLatestCase{
			Province: province,
		}

		// Get latest case for this province
		latestCase, err := s.provinceCaseRepo.GetLatestByProvinceID(province.ID)
		if err != nil {
			// If error or no data, continue without latest case
			continue
		}

		if latestCase != nil {
			// Transform to response format
			caseResponse := latestCase.TransformToResponse()
			result[i].LatestCase = &caseResponse
		}
	}

	return result, nil
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

func (s *covidService) GetAllProvinceCasesSorted(sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	cases, err := s.provinceCaseRepo.GetAllSorted(sortParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get sorted province cases: %w", err)
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

func (s *covidService) GetProvinceCasesPaginated(provinceID string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	cases, total, err := s.provinceCaseRepo.GetByProvinceIDPaginated(provinceID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get province cases paginated: %w", err)
	}
	return cases, total, nil
}

func (s *covidService) GetProvinceCasesByDateRangePaginated(provinceID, startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, total, err := s.provinceCaseRepo.GetByProvinceIDAndDateRangePaginated(provinceID, start, end, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get province cases by date range paginated: %w", err)
	}
	return cases, total, nil
}

func (s *covidService) GetAllProvinceCasesPaginated(limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	cases, total, err := s.provinceCaseRepo.GetAllPaginated(limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all province cases paginated: %w", err)
	}
	return cases, total, nil
}

func (s *covidService) GetAllProvinceCasesByDateRangePaginated(startDate, endDate string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, total, err := s.provinceCaseRepo.GetByDateRangePaginated(start, end, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all province cases by date range paginated: %w", err)
	}
	return cases, total, nil
}

func (s *covidService) GetAllProvinceCasesPaginatedSorted(limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	cases, total, err := s.provinceCaseRepo.GetAllPaginatedSorted(limit, offset, sortParams)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get sorted province cases paginated: %w", err)
	}
	return cases, total, nil
}

func (s *covidService) GetAllProvinceCasesByDateRangeSorted(startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, err := s.provinceCaseRepo.GetByDateRangeSorted(start, end, sortParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get sorted province cases by date range: %w", err)
	}
	return cases, nil
}

func (s *covidService) GetAllProvinceCasesByDateRangePaginatedSorted(startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, total, err := s.provinceCaseRepo.GetByDateRangePaginatedSorted(start, end, limit, offset, sortParams)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get sorted province cases by date range paginated: %w", err)
	}
	return cases, total, nil
}

func (s *covidService) GetProvinceCasesSorted(provinceID string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	cases, err := s.provinceCaseRepo.GetByProvinceIDSorted(provinceID, sortParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get sorted province cases: %w", err)
	}
	return cases, nil
}

func (s *covidService) GetProvinceCasesPaginatedSorted(provinceID string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	cases, total, err := s.provinceCaseRepo.GetByProvinceIDPaginatedSorted(provinceID, limit, offset, sortParams)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get sorted province cases paginated: %w", err)
	}
	return cases, total, nil
}

func (s *covidService) GetProvinceCasesByDateRangeSorted(provinceID, startDate, endDate string, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, err := s.provinceCaseRepo.GetByProvinceIDAndDateRangeSorted(provinceID, start, end, sortParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get sorted province cases by date range: %w", err)
	}
	return cases, nil
}

func (s *covidService) GetProvinceCasesByDateRangePaginatedSorted(provinceID, startDate, endDate string, limit, offset int, sortParams utils.SortParams) ([]models.ProvinceCaseWithDate, int, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid start date format: %w", err)
	}

	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid end date format: %w", err)
	}

	cases, total, err := s.provinceCaseRepo.GetByProvinceIDAndDateRangePaginatedSorted(provinceID, start, end, limit, offset, sortParams)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get sorted province cases by date range paginated: %w", err)
	}
	return cases, total, nil
}
