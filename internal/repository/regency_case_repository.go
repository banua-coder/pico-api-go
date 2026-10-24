package repository

import (
	"fmt"
	"log"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

// RegencyCaseRepositoryInterface defines the contract for regency case repository operations
type RegencyCaseRepositoryInterface interface {
	GetByRegencyID(regencyID int) ([]models.RegencyCase, error)
	GetLatest() ([]models.RegencyCase, error)
}

// RegencyCaseRepository handles database operations for regency cases
type RegencyCaseRepository struct {
	db *database.DB
}

// NewRegencyCaseRepository creates a new RegencyCaseRepository
func NewRegencyCaseRepository(db *database.DB) *RegencyCaseRepository {
	return &RegencyCaseRepository{db: db}
}

// GetByRegencyID returns all cases for a specific regency
func (r *RegencyCaseRepository) GetByRegencyID(regencyID int) ([]models.RegencyCase, error) {
	query := `SELECT rc.id, rc.day, rc.regency_id, rc.positive, rc.recovered, rc.deceased,
		rc.person_under_observation, rc.finished_person_under_observation,
		rc.person_under_supervision, rc.finished_person_under_supervision,
		rc.cumulative_positive, rc.cumulative_recovered, rc.cumulative_deceased,
		rc.cumulative_person_under_observation, rc.cumulative_finished_person_under_observation,
		rc.cumulative_person_under_supervision, rc.cumulative_finished_person_under_supervision,
		rc.rt, rc.rt_upper, rc.rt_lower,
		nc.date, reg.id, reg.name
		FROM regency_cases rc
		JOIN national_cases nc ON rc.day = nc.id
		JOIN regencies reg ON rc.regency_id = reg.id
		WHERE rc.regency_id = ?
		ORDER BY rc.day ASC`

	rows, err := r.db.Query(query, regencyID)
	if err != nil {
		return nil, fmt.Errorf("failed to query regency cases: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var cases []models.RegencyCase
	for rows.Next() {
		var c models.RegencyCase
		var regID int
		var regName string

		if err := rows.Scan(&c.ID, &c.Day, &c.RegencyID,
			&c.Positive, &c.Recovered, &c.Deceased,
			&c.PersonUnderObservation, &c.FinishedPersonUnderObservation,
			&c.PersonUnderSupervision, &c.FinishedPersonUnderSupervision,
			&c.CumulativePositive, &c.CumulativeRecovered, &c.CumulativeDeceased,
			&c.CumulativePersonUnderObservation, &c.CumulativeFinishedPersonUnderObservation,
			&c.CumulativePersonUnderSupervision, &c.CumulativeFinishedPersonUnderSupervision,
			&c.Rt, &c.RtUpper, &c.RtLower,
			&c.Date, &regID, &regName); err != nil {
			return nil, fmt.Errorf("failed to scan regency case: %w", err)
		}
		c.Regency = &models.Regency{ID: regID, Name: regName}
		cases = append(cases, c)
	}
	return cases, rows.Err()
}

// GetLatest returns the latest case for every regency that has data, in any province.
func (r *RegencyCaseRepository) GetLatest() ([]models.RegencyCase, error) {
	return r.queryLatest()
}

// queryLatest picks, per regency, the row with the highest day that has any non-zero daily count.
// The MAX(day) is computed once per regency in a non-correlated derived table; the earlier
// correlated subquery re-scanned regency_cases for every outer row (about 18k x 1.4k rows).
func (r *RegencyCaseRepository) queryLatest() ([]models.RegencyCase, error) {
	query := `SELECT rc.id, rc.day, rc.regency_id, rc.positive, rc.recovered, rc.deceased,
		rc.cumulative_positive, rc.cumulative_recovered, rc.cumulative_deceased,
		rc.person_under_observation, rc.finished_person_under_observation,
		rc.person_under_supervision, rc.finished_person_under_supervision,
		rc.cumulative_person_under_observation, rc.cumulative_finished_person_under_observation,
		rc.cumulative_person_under_supervision, rc.cumulative_finished_person_under_supervision,
		rc.rt, rc.rt_upper, rc.rt_lower,
		nc.date, reg.id, reg.name
		FROM regency_cases rc
		JOIN (SELECT regency_id, MAX(day) AS day FROM regency_cases
			WHERE (positive > 0 OR recovered > 0 OR deceased > 0)
			GROUP BY regency_id) m ON m.regency_id = rc.regency_id AND m.day = rc.day
		JOIN national_cases nc ON rc.day = nc.id
		JOIN regencies reg ON rc.regency_id = reg.id
		ORDER BY reg.name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query latest regency cases: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var cases []models.RegencyCase
	for rows.Next() {
		var c models.RegencyCase
		var regID int
		var regName string

		if err := rows.Scan(&c.ID, &c.Day, &c.RegencyID,
			&c.Positive, &c.Recovered, &c.Deceased,
			&c.CumulativePositive, &c.CumulativeRecovered, &c.CumulativeDeceased,
			&c.PersonUnderObservation, &c.FinishedPersonUnderObservation,
			&c.PersonUnderSupervision, &c.FinishedPersonUnderSupervision,
			&c.CumulativePersonUnderObservation, &c.CumulativeFinishedPersonUnderObservation,
			&c.CumulativePersonUnderSupervision, &c.CumulativeFinishedPersonUnderSupervision,
			&c.Rt, &c.RtUpper, &c.RtLower,
			&c.Date, &regID, &regName); err != nil {
			return nil, fmt.Errorf("failed to scan regency case: %w", err)
		}
		c.Regency = &models.Regency{ID: regID, Name: regName}
		cases = append(cases, c)
	}
	return cases, rows.Err()
}
