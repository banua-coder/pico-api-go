package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
	"github.com/banua-coder/pico-api-go/pkg/utils"
)

type NationalCaseRepository interface {
	GetAll() ([]models.NationalCase, error)
	GetAllSorted(sortParams utils.SortParams) ([]models.NationalCase, error)
	GetByDateRange(startDate, endDate time.Time) ([]models.NationalCase, error)
	GetByDateRangeSorted(startDate, endDate time.Time, sortParams utils.SortParams) ([]models.NationalCase, error)
	GetLatest() (*models.NationalCase, error)
	GetByDay(day int64) (*models.NationalCase, error)
}

type nationalCaseRepository struct {
	db *database.DB
}

func NewNationalCaseRepository(db *database.DB) NationalCaseRepository {
	return &nationalCaseRepository{db: db}
}

func (r *nationalCaseRepository) GetAll() ([]models.NationalCase, error) {
	// Default sorting by date ascending
	return r.GetAllSorted(utils.SortParams{Field: "date", Order: "asc"})
}

func (r *nationalCaseRepository) GetAllSorted(sortParams utils.SortParams) ([]models.NationalCase, error) {
	query := `SELECT id, day, date, positive, recovered, deceased, 
			  cumulative_positive, cumulative_recovered, cumulative_deceased,
			  rt, rt_upper, rt_lower 
			  FROM national_cases ORDER BY ` + sortParams.GetSQLOrderClause()

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query national cases: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("Error closing rows: %v\n", err)
		}
	}()

	var cases []models.NationalCase
	for rows.Next() {
		var c models.NationalCase
		err := rows.Scan(&c.ID, &c.Day, &c.Date, &c.Positive, &c.Recovered, &c.Deceased,
			&c.CumulativePositive, &c.CumulativeRecovered, &c.CumulativeDeceased,
			&c.Rt, &c.RtUpper, &c.RtLower)
		if err != nil {
			return nil, fmt.Errorf("failed to scan national case: %w", err)
		}
		cases = append(cases, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return cases, nil
}

func (r *nationalCaseRepository) GetByDateRange(startDate, endDate time.Time) ([]models.NationalCase, error) {
	// Default sorting by date ascending
	return r.GetByDateRangeSorted(startDate, endDate, utils.SortParams{Field: "date", Order: "asc"})
}

func (r *nationalCaseRepository) GetByDateRangeSorted(startDate, endDate time.Time, sortParams utils.SortParams) ([]models.NationalCase, error) {
	query := `SELECT id, day, date, positive, recovered, deceased, 
			  cumulative_positive, cumulative_recovered, cumulative_deceased,
			  rt, rt_upper, rt_lower 
			  FROM national_cases 
			  WHERE date BETWEEN ? AND ? 
			  ORDER BY ` + sortParams.GetSQLOrderClause()

	rows, err := r.db.Query(query, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to query national cases by date range: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("Error closing rows: %v\n", err)
		}
	}()

	var cases []models.NationalCase
	for rows.Next() {
		var c models.NationalCase
		err := rows.Scan(&c.ID, &c.Day, &c.Date, &c.Positive, &c.Recovered, &c.Deceased,
			&c.CumulativePositive, &c.CumulativeRecovered, &c.CumulativeDeceased,
			&c.Rt, &c.RtUpper, &c.RtLower)
		if err != nil {
			return nil, fmt.Errorf("failed to scan national case: %w", err)
		}
		cases = append(cases, c)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return cases, nil
}

func (r *nationalCaseRepository) GetLatest() (*models.NationalCase, error) {
	query := `SELECT id, day, date, positive, recovered, deceased, 
			  cumulative_positive, cumulative_recovered, cumulative_deceased,
			  rt, rt_upper, rt_lower 
			  FROM national_cases 
			  ORDER BY date DESC LIMIT 1`

	var c models.NationalCase
	err := r.db.QueryRow(query).Scan(&c.ID, &c.Day, &c.Date, &c.Positive, &c.Recovered, &c.Deceased,
		&c.CumulativePositive, &c.CumulativeRecovered, &c.CumulativeDeceased,
		&c.Rt, &c.RtUpper, &c.RtLower)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get latest national case: %w", err)
	}

	return &c, nil
}

func (r *nationalCaseRepository) GetByDay(day int64) (*models.NationalCase, error) {
	query := `SELECT id, day, date, positive, recovered, deceased, 
			  cumulative_positive, cumulative_recovered, cumulative_deceased,
			  rt, rt_upper, rt_lower 
			  FROM national_cases 
			  WHERE day = ?`

	var c models.NationalCase
	err := r.db.QueryRow(query, day).Scan(&c.ID, &c.Day, &c.Date, &c.Positive, &c.Recovered, &c.Deceased,
		&c.CumulativePositive, &c.CumulativeRecovered, &c.CumulativeDeceased,
		&c.Rt, &c.RtUpper, &c.RtLower)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get national case by day: %w", err)
	}

	return &c, nil
}
