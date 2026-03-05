package repository

import (
	"fmt"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

// ProvinceStatsRepositoryInterface defines the contract for province stats repository operations
type ProvinceStatsRepositoryInterface interface {
	GetGenderCases(provinceID int) ([]models.ProvinceGenderCase, error)
	GetLatestGenderCase(provinceID int) (*models.ProvinceGenderCase, error)
	GetTests(provinceID int) ([]models.ProvinceTest, error)
	GetTestTypes() ([]models.TestType, error)
}

// ProvinceStatsRepository handles province gender cases and test data
type ProvinceStatsRepository struct {
	db *database.DB
}

func NewProvinceStatsRepository(db *database.DB) *ProvinceStatsRepository {
	return &ProvinceStatsRepository{db: db}
}

// GetGenderCases returns all gender/age-based case data for a province
func (r *ProvinceStatsRepository) GetGenderCases(provinceID int) ([]models.ProvinceGenderCase, error) {
	query := `SELECT id, day, province_id,
		positive_male, positive_female, pdp_male, pdp_female,
		positive_male_0_14, positive_male_15_19, positive_male_20_24, positive_male_25_49, positive_male_50_54, positive_male_55,
		positive_female_0_14, positive_female_15_19, positive_female_20_24, positive_female_25_49, positive_female_50_54, positive_female_55,
		pdp_male_0_14, pdp_male_15_19, pdp_male_20_24, pdp_male_25_49, pdp_male_50_54, pdp_male_55,
		pdp_female_0_14, pdp_female_15_19, pdp_female_20_24, pdp_female_25_49, pdp_female_50_54, pdp_female_55
		FROM province_gender_cases WHERE province_id = ? ORDER BY day ASC`

	rows, err := r.db.Query(query, provinceID)
	if err != nil {
		return nil, fmt.Errorf("failed to query gender cases: %w", err)
	}
	defer rows.Close()

	var cases []models.ProvinceGenderCase
	for rows.Next() {
		var c models.ProvinceGenderCase
		if err := rows.Scan(&c.ID, &c.Day, &c.ProvinceID,
			&c.PositiveMale, &c.PositiveFemale, &c.PDPMale, &c.PDPFemale,
			&c.PositiveMale0_14, &c.PositiveMale15_19, &c.PositiveMale20_24, &c.PositiveMale25_49, &c.PositiveMale50_54, &c.PositiveMale55,
			&c.PositiveFemale0_14, &c.PositiveFemale15_19, &c.PositiveFemale20_24, &c.PositiveFemale25_49, &c.PositiveFemale50_54, &c.PositiveFemale55,
			&c.PDPMale0_14, &c.PDPMale15_19, &c.PDPMale20_24, &c.PDPMale25_49, &c.PDPMale50_54, &c.PDPMale55,
			&c.PDPFemale0_14, &c.PDPFemale15_19, &c.PDPFemale20_24, &c.PDPFemale25_49, &c.PDPFemale50_54, &c.PDPFemale55,
		); err != nil {
			return nil, fmt.Errorf("failed to scan gender case: %w", err)
		}
		cases = append(cases, c)
	}
	return cases, rows.Err()
}

// GetLatestGenderCase returns the latest gender case for a province
func (r *ProvinceStatsRepository) GetLatestGenderCase(provinceID int) (*models.ProvinceGenderCase, error) {
	query := `SELECT id, day, province_id,
		positive_male, positive_female, pdp_male, pdp_female,
		positive_male_0_14, positive_male_15_19, positive_male_20_24, positive_male_25_49, positive_male_50_54, positive_male_55,
		positive_female_0_14, positive_female_15_19, positive_female_20_24, positive_female_25_49, positive_female_50_54, positive_female_55,
		pdp_male_0_14, pdp_male_15_19, pdp_male_20_24, pdp_male_25_49, pdp_male_50_54, pdp_male_55,
		pdp_female_0_14, pdp_female_15_19, pdp_female_20_24, pdp_female_25_49, pdp_female_50_54, pdp_female_55
		FROM province_gender_cases WHERE province_id = ? ORDER BY day DESC LIMIT 1`

	var c models.ProvinceGenderCase
	err := r.db.QueryRow(query, provinceID).Scan(&c.ID, &c.Day, &c.ProvinceID,
		&c.PositiveMale, &c.PositiveFemale, &c.PDPMale, &c.PDPFemale,
		&c.PositiveMale0_14, &c.PositiveMale15_19, &c.PositiveMale20_24, &c.PositiveMale25_49, &c.PositiveMale50_54, &c.PositiveMale55,
		&c.PositiveFemale0_14, &c.PositiveFemale15_19, &c.PositiveFemale20_24, &c.PositiveFemale25_49, &c.PositiveFemale50_54, &c.PositiveFemale55,
		&c.PDPMale0_14, &c.PDPMale15_19, &c.PDPMale20_24, &c.PDPMale25_49, &c.PDPMale50_54, &c.PDPMale55,
		&c.PDPFemale0_14, &c.PDPFemale15_19, &c.PDPFemale20_24, &c.PDPFemale25_49, &c.PDPFemale50_54, &c.PDPFemale55,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get latest gender case: %w", err)
	}
	return &c, nil
}

// GetTests returns all test data for a province with test type info
func (r *ProvinceStatsRepository) GetTests(provinceID int) ([]models.ProvinceTest, error) {
	query := `SELECT pt.id, pt.test_type_id, pt.day, pt.province_id, pt.date_from,
		pt.process, pt.invalid, pt.positive, pt.negative,
		tt.id, tt.key, tt.name, tt.sample, tt.duration, tt.is_recommended
		FROM province_tests pt
		JOIN test_types tt ON pt.test_type_id = tt.id
		WHERE pt.province_id = ?
		ORDER BY pt.day ASC, pt.test_type_id ASC`

	rows, err := r.db.Query(query, provinceID)
	if err != nil {
		return nil, fmt.Errorf("failed to query province tests: %w", err)
	}
	defer rows.Close()

	var tests []models.ProvinceTest
	for rows.Next() {
		var t models.ProvinceTest
		var tt models.TestType
		if err := rows.Scan(&t.ID, &t.TestTypeID, &t.Day, &t.ProvinceID, &t.DateFrom,
			&t.Process, &t.Invalid, &t.Positive, &t.Negative,
			&tt.ID, &tt.Key, &tt.Name, &tt.Sample, &tt.Duration, &tt.IsRecommended,
		); err != nil {
			return nil, fmt.Errorf("failed to scan province test: %w", err)
		}
		t.TestType = &tt
		tests = append(tests, t)
	}
	return tests, rows.Err()
}

// GetTestTypes returns all available test types
func (r *ProvinceStatsRepository) GetTestTypes() ([]models.TestType, error) {
	query := `SELECT id, ` + "`key`" + `, name, sample, duration, is_recommended FROM test_types ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query test types: %w", err)
	}
	defer rows.Close()

	var types []models.TestType
	for rows.Next() {
		var t models.TestType
		if err := rows.Scan(&t.ID, &t.Key, &t.Name, &t.Sample, &t.Duration, &t.IsRecommended); err != nil {
			return nil, fmt.Errorf("failed to scan test type: %w", err)
		}
		types = append(types, t)
	}
	return types, rows.Err()
}
