package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

type ProvinceCaseRepository interface {
	GetAll() ([]models.ProvinceCaseWithDate, error)
	GetAllPaginated(limit, offset int) ([]models.ProvinceCaseWithDate, int, error)
	GetByProvinceID(provinceID string) ([]models.ProvinceCaseWithDate, error)
	GetByProvinceIDPaginated(provinceID string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error)
	GetByProvinceIDAndDateRange(provinceID string, startDate, endDate time.Time) ([]models.ProvinceCaseWithDate, error)
	GetByProvinceIDAndDateRangePaginated(provinceID string, startDate, endDate time.Time, limit, offset int) ([]models.ProvinceCaseWithDate, int, error)
	GetByDateRange(startDate, endDate time.Time) ([]models.ProvinceCaseWithDate, error)
	GetByDateRangePaginated(startDate, endDate time.Time, limit, offset int) ([]models.ProvinceCaseWithDate, int, error)
	GetLatestByProvinceID(provinceID string) (*models.ProvinceCaseWithDate, error)
}

type provinceCaseRepository struct {
	db *database.DB
}

func NewProvinceCaseRepository(db *database.DB) ProvinceCaseRepository {
	return &provinceCaseRepository{db: db}
}

func (r *provinceCaseRepository) GetAll() ([]models.ProvinceCaseWithDate, error) {
	query := `SELECT pc.id, pc.day, pc.province_id, pc.positive, pc.recovered, pc.deceased,
			  pc.person_under_observation, pc.finished_person_under_observation,
			  pc.person_under_supervision, pc.finished_person_under_supervision,
			  pc.cumulative_positive, pc.cumulative_recovered, pc.cumulative_deceased,
			  pc.cumulative_person_under_observation, pc.cumulative_finished_person_under_observation,
			  pc.cumulative_person_under_supervision, pc.cumulative_finished_person_under_supervision,
			  pc.rt, pc.rt_upper, pc.rt_lower, nc.date, p.name
			  FROM province_cases pc
			  JOIN national_cases nc ON pc.day = nc.id
			  LEFT JOIN provinces p ON pc.province_id = p.id
			  ORDER BY nc.date DESC, p.name`

	return r.queryProvinceCases(query)
}

func (r *provinceCaseRepository) GetAllPaginated(limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	// First get total count
	countQuery := `SELECT COUNT(*) FROM province_cases pc
				   JOIN national_cases nc ON pc.day = nc.id`
	
	var total int
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count province cases: %w", err)
	}
	
	// Get paginated data
	query := `SELECT pc.id, pc.day, pc.province_id, pc.positive, pc.recovered, pc.deceased,
			  pc.person_under_observation, pc.finished_person_under_observation,
			  pc.person_under_supervision, pc.finished_person_under_supervision,
			  pc.cumulative_positive, pc.cumulative_recovered, pc.cumulative_deceased,
			  pc.cumulative_person_under_observation, pc.cumulative_finished_person_under_observation,
			  pc.cumulative_person_under_supervision, pc.cumulative_finished_person_under_supervision,
			  pc.rt, pc.rt_upper, pc.rt_lower, nc.date, p.name
			  FROM province_cases pc
			  JOIN national_cases nc ON pc.day = nc.id
			  LEFT JOIN provinces p ON pc.province_id = p.id
			  ORDER BY nc.date DESC, p.name
			  LIMIT ? OFFSET ?`
	
	cases, err := r.queryProvinceCases(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	
	return cases, total, nil
}

func (r *provinceCaseRepository) GetByProvinceID(provinceID string) ([]models.ProvinceCaseWithDate, error) {
	query := `SELECT pc.id, pc.day, pc.province_id, pc.positive, pc.recovered, pc.deceased,
			  pc.person_under_observation, pc.finished_person_under_observation,
			  pc.person_under_supervision, pc.finished_person_under_supervision,
			  pc.cumulative_positive, pc.cumulative_recovered, pc.cumulative_deceased,
			  pc.cumulative_person_under_observation, pc.cumulative_finished_person_under_observation,
			  pc.cumulative_person_under_supervision, pc.cumulative_finished_person_under_supervision,
			  pc.rt, pc.rt_upper, pc.rt_lower, nc.date, p.name
			  FROM province_cases pc
			  JOIN national_cases nc ON pc.day = nc.id
			  LEFT JOIN provinces p ON pc.province_id = p.id
			  WHERE pc.province_id = ?
			  ORDER BY nc.date DESC`

	return r.queryProvinceCases(query, provinceID)
}

func (r *provinceCaseRepository) GetByProvinceIDPaginated(provinceID string, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	// First get total count
	countQuery := `SELECT COUNT(*) FROM province_cases pc
				   JOIN national_cases nc ON pc.day = nc.id
				   WHERE pc.province_id = ?`
	
	var total int
	err := r.db.QueryRow(countQuery, provinceID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count province cases for province %s: %w", provinceID, err)
	}
	
	// Get paginated data
	query := `SELECT pc.id, pc.day, pc.province_id, pc.positive, pc.recovered, pc.deceased,
			  pc.person_under_observation, pc.finished_person_under_observation,
			  pc.person_under_supervision, pc.finished_person_under_supervision,
			  pc.cumulative_positive, pc.cumulative_recovered, pc.cumulative_deceased,
			  pc.cumulative_person_under_observation, pc.cumulative_finished_person_under_observation,
			  pc.cumulative_person_under_supervision, pc.cumulative_finished_person_under_supervision,
			  pc.rt, pc.rt_upper, pc.rt_lower, nc.date, p.name
			  FROM province_cases pc
			  JOIN national_cases nc ON pc.day = nc.id
			  LEFT JOIN provinces p ON pc.province_id = p.id
			  WHERE pc.province_id = ?
			  ORDER BY nc.date DESC
			  LIMIT ? OFFSET ?`
	
	cases, err := r.queryProvinceCases(query, provinceID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	
	return cases, total, nil
}

func (r *provinceCaseRepository) GetByProvinceIDAndDateRange(provinceID string, startDate, endDate time.Time) ([]models.ProvinceCaseWithDate, error) {
	query := `SELECT pc.id, pc.day, pc.province_id, pc.positive, pc.recovered, pc.deceased,
			  pc.person_under_observation, pc.finished_person_under_observation,
			  pc.person_under_supervision, pc.finished_person_under_supervision,
			  pc.cumulative_positive, pc.cumulative_recovered, pc.cumulative_deceased,
			  pc.cumulative_person_under_observation, pc.cumulative_finished_person_under_observation,
			  pc.cumulative_person_under_supervision, pc.cumulative_finished_person_under_supervision,
			  pc.rt, pc.rt_upper, pc.rt_lower, nc.date, p.name
			  FROM province_cases pc
			  JOIN national_cases nc ON pc.day = nc.id
			  LEFT JOIN provinces p ON pc.province_id = p.id
			  WHERE pc.province_id = ? AND nc.date BETWEEN ? AND ?
			  ORDER BY nc.date DESC`

	return r.queryProvinceCases(query, provinceID, startDate, endDate)
}

func (r *provinceCaseRepository) GetByProvinceIDAndDateRangePaginated(provinceID string, startDate, endDate time.Time, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	// First get total count
	countQuery := `SELECT COUNT(*) FROM province_cases pc
				   JOIN national_cases nc ON pc.day = nc.id
				   WHERE pc.province_id = ? AND nc.date BETWEEN ? AND ?`
	
	var total int
	err := r.db.QueryRow(countQuery, provinceID, startDate, endDate).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count province cases for province %s in date range: %w", provinceID, err)
	}
	
	// Get paginated data
	query := `SELECT pc.id, pc.day, pc.province_id, pc.positive, pc.recovered, pc.deceased,
			  pc.person_under_observation, pc.finished_person_under_observation,
			  pc.person_under_supervision, pc.finished_person_under_supervision,
			  pc.cumulative_positive, pc.cumulative_recovered, pc.cumulative_deceased,
			  pc.cumulative_person_under_observation, pc.cumulative_finished_person_under_observation,
			  pc.cumulative_person_under_supervision, pc.cumulative_finished_person_under_supervision,
			  pc.rt, pc.rt_upper, pc.rt_lower, nc.date, p.name
			  FROM province_cases pc
			  JOIN national_cases nc ON pc.day = nc.id
			  LEFT JOIN provinces p ON pc.province_id = p.id
			  WHERE pc.province_id = ? AND nc.date BETWEEN ? AND ?
			  ORDER BY nc.date DESC
			  LIMIT ? OFFSET ?`
	
	cases, err := r.queryProvinceCases(query, provinceID, startDate, endDate, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	
	return cases, total, nil
}

func (r *provinceCaseRepository) GetByDateRange(startDate, endDate time.Time) ([]models.ProvinceCaseWithDate, error) {
	query := `SELECT pc.id, pc.day, pc.province_id, pc.positive, pc.recovered, pc.deceased,
			  pc.person_under_observation, pc.finished_person_under_observation,
			  pc.person_under_supervision, pc.finished_person_under_supervision,
			  pc.cumulative_positive, pc.cumulative_recovered, pc.cumulative_deceased,
			  pc.cumulative_person_under_observation, pc.cumulative_finished_person_under_observation,
			  pc.cumulative_person_under_supervision, pc.cumulative_finished_person_under_supervision,
			  pc.rt, pc.rt_upper, pc.rt_lower, nc.date, p.name
			  FROM province_cases pc
			  JOIN national_cases nc ON pc.day = nc.id
			  LEFT JOIN provinces p ON pc.province_id = p.id
			  WHERE nc.date BETWEEN ? AND ?
			  ORDER BY nc.date DESC, p.name`

	return r.queryProvinceCases(query, startDate, endDate)
}

func (r *provinceCaseRepository) GetByDateRangePaginated(startDate, endDate time.Time, limit, offset int) ([]models.ProvinceCaseWithDate, int, error) {
	// First get total count
	countQuery := `SELECT COUNT(*) FROM province_cases pc
				   JOIN national_cases nc ON pc.day = nc.id
				   WHERE nc.date BETWEEN ? AND ?`
	
	var total int
	err := r.db.QueryRow(countQuery, startDate, endDate).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count province cases in date range: %w", err)
	}
	
	// Get paginated data
	query := `SELECT pc.id, pc.day, pc.province_id, pc.positive, pc.recovered, pc.deceased,
			  pc.person_under_observation, pc.finished_person_under_observation,
			  pc.person_under_supervision, pc.finished_person_under_supervision,
			  pc.cumulative_positive, pc.cumulative_recovered, pc.cumulative_deceased,
			  pc.cumulative_person_under_observation, pc.cumulative_finished_person_under_observation,
			  pc.cumulative_person_under_supervision, pc.cumulative_finished_person_under_supervision,
			  pc.rt, pc.rt_upper, pc.rt_lower, nc.date, p.name
			  FROM province_cases pc
			  JOIN national_cases nc ON pc.day = nc.id
			  LEFT JOIN provinces p ON pc.province_id = p.id
			  WHERE nc.date BETWEEN ? AND ?
			  ORDER BY nc.date DESC, p.name
			  LIMIT ? OFFSET ?`
	
	cases, err := r.queryProvinceCases(query, startDate, endDate, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	
	return cases, total, nil
}

func (r *provinceCaseRepository) GetLatestByProvinceID(provinceID string) (*models.ProvinceCaseWithDate, error) {
	query := `SELECT pc.id, pc.day, pc.province_id, pc.positive, pc.recovered, pc.deceased,
			  pc.person_under_observation, pc.finished_person_under_observation,
			  pc.person_under_supervision, pc.finished_person_under_supervision,
			  pc.cumulative_positive, pc.cumulative_recovered, pc.cumulative_deceased,
			  pc.cumulative_person_under_observation, pc.cumulative_finished_person_under_observation,
			  pc.cumulative_person_under_supervision, pc.cumulative_finished_person_under_supervision,
			  pc.rt, pc.rt_upper, pc.rt_lower, nc.date, p.name
			  FROM province_cases pc
			  JOIN national_cases nc ON pc.day = nc.id
			  LEFT JOIN provinces p ON pc.province_id = p.id
			  WHERE pc.province_id = ?
			  ORDER BY nc.date DESC LIMIT 1`

	cases, err := r.queryProvinceCases(query, provinceID)
	if err != nil {
		return nil, err
	}

	if len(cases) == 0 {
		return nil, nil
	}

	return &cases[0], nil
}

func (r *provinceCaseRepository) queryProvinceCases(query string, args ...interface{}) ([]models.ProvinceCaseWithDate, error) {
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query province cases: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("Error closing rows: %v\n", err)
		}
	}()

	var cases []models.ProvinceCaseWithDate
	for rows.Next() {
		var c models.ProvinceCaseWithDate
		var provinceName sql.NullString

		err := rows.Scan(&c.ID, &c.Day, &c.ProvinceID, &c.Positive, &c.Recovered, &c.Deceased,
			&c.PersonUnderObservation, &c.FinishedPersonUnderObservation,
			&c.PersonUnderSupervision, &c.FinishedPersonUnderSupervision,
			&c.CumulativePositive, &c.CumulativeRecovered, &c.CumulativeDeceased,
			&c.CumulativePersonUnderObservation, &c.CumulativeFinishedPersonUnderObservation,
			&c.CumulativePersonUnderSupervision, &c.CumulativeFinishedPersonUnderSupervision,
			&c.Rt, &c.RtUpper, &c.RtLower, &c.Date, &provinceName)
		if err != nil {
			return nil, fmt.Errorf("failed to scan province case: %w", err)
		}

		if provinceName.Valid {
			c.Province = &models.Province{
				ID:   c.ProvinceID,
				Name: provinceName.String,
			}
		}

		cases = append(cases, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return cases, nil
}