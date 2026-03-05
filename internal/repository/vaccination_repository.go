package repository

import (
	"fmt"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

// VaccinationRepository handles database operations for vaccination data
type VaccinationRepository struct {
	db *database.DB
}

// NewVaccinationRepository creates a new VaccinationRepository
func NewVaccinationRepository(db *database.DB) *VaccinationRepository {
	return &VaccinationRepository{db: db}
}

// GetNationalVaccinations returns all national vaccination data
func (r *VaccinationRepository) GetNationalVaccinations() ([]models.NationalVaccine, error) {
	query := `SELECT id, day, date, total_vaccination_target,
		first_vaccination_received, second_vaccination_received,
		cumulative_first_vaccination_received, cumulative_second_vaccination_received,
		health_worker_vaccination_target, health_worker_first_vaccination_received, health_worker_second_vaccination_received,
		cumulative_health_worker_first_vaccination_received, cumulative_health_worker_second_vaccination_received,
		elderly_vaccination_target, elderly_first_vaccination_received, elderly_second_vaccination_received,
		cumulative_elderly_first_vaccination_received, cumulative_elderly_second_vaccination_received,
		public_officer_vaccination_target, public_officer_first_vaccination_received, public_officer_second_vaccination_received,
		cumulative_public_officer_first_vaccination_received, cumulative_public_officer_second_vaccination_received,
		public_vaccination_target, public_first_vaccination_received, public_second_vaccination_received,
		cumulative_public_first_vaccination_received, cumulative_public_second_vaccination_received,
		teenager_vaccination_target, teenager_first_vaccination_received, teenager_second_vaccination_received,
		cumulative_teenager_first_vaccination_received, cumulative_teenager_second_vaccination_received
		FROM national_vaccines ORDER BY day ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query national vaccinations: %w", err)
	}
	defer rows.Close()

	var vaccines []models.NationalVaccine
	for rows.Next() {
		var v models.NationalVaccine
		if err := rows.Scan(&v.ID, &v.Day, &v.Date, &v.TotalVaccinationTarget,
			&v.FirstVaccinationReceived, &v.SecondVaccinationReceived,
			&v.CumulativeFirstVaccinationReceived, &v.CumulativeSecondVaccinationReceived,
			&v.HealthWorkerVaccinationTarget, &v.HealthWorkerFirstVaccinationReceived, &v.HealthWorkerSecondVaccinationReceived,
			&v.CumulativeHealthWorkerFirstVaccinationReceived, &v.CumulativeHealthWorkerSecondVaccinationReceived,
			&v.ElderlyVaccinationTarget, &v.ElderlyFirstVaccinationReceived, &v.ElderlySecondVaccinationReceived,
			&v.CumulativeElderlyFirstVaccinationReceived, &v.CumulativeElderlySecondVaccinationReceived,
			&v.PublicOfficerVaccinationTarget, &v.PublicOfficerFirstVaccinationReceived, &v.PublicOfficerSecondVaccinationReceived,
			&v.CumulativePublicOfficerFirstVaccinationReceived, &v.CumulativePublicOfficerSecondVaccinationReceived,
			&v.PublicVaccinationTarget, &v.PublicFirstVaccinationReceived, &v.PublicSecondVaccinationReceived,
			&v.CumulativePublicFirstVaccinationReceived, &v.CumulativePublicSecondVaccinationReceived,
			&v.TeenagerVaccinationTarget, &v.TeenagerFirstVaccinationReceived, &v.TeenagerSecondVaccinationReceived,
			&v.CumulativeTeenagerFirstVaccinationReceived, &v.CumulativeTeenagerSecondVaccinationReceived,
		); err != nil {
			return nil, fmt.Errorf("failed to scan national vaccine: %w", err)
		}
		vaccines = append(vaccines, v)
	}
	return vaccines, rows.Err()
}

// GetProvinceVaccinations returns vaccination data for a province (default: SulTeng = 72)
func (r *VaccinationRepository) GetProvinceVaccinations(provinceID int) ([]models.ProvinceVaccine, error) {
	query := `SELECT id, day, province_id, date, total_vaccination_target,
		first_vaccination_received, second_vaccination_received,
		cumulative_first_vaccination_received, cumulative_second_vaccination_received,
		health_worker_vaccination_target, health_worker_first_vaccination_received, health_worker_second_vaccination_received,
		cumulative_health_worker_first_vaccination_received, cumulative_health_worker_second_vaccination_received,
		elderly_vaccination_target, elderly_first_vaccination_received, elderly_second_vaccination_received,
		cumulative_elderly_first_vaccination_received, cumulative_elderly_second_vaccination_received,
		public_officer_vaccination_target, public_officer_first_vaccination_received, public_officer_second_vaccination_received,
		cumulative_public_officer_first_vaccination_received, cumulative_public_officer_second_vaccination_received,
		public_vaccination_target, public_first_vaccination_received, public_second_vaccination_received,
		cumulative_public_first_vaccination_received, cumulative_public_second_vaccination_received,
		teenager_vaccination_target, teenager_first_vaccination_received, teenager_second_vaccination_received,
		cumulative_teenager_first_vaccination_received, cumulative_teenager_second_vaccination_received
		FROM province_vaccines WHERE province_id = ? ORDER BY day ASC`

	rows, err := r.db.Query(query, provinceID)
	if err != nil {
		return nil, fmt.Errorf("failed to query province vaccinations: %w", err)
	}
	defer rows.Close()

	var vaccines []models.ProvinceVaccine
	for rows.Next() {
		var v models.ProvinceVaccine
		if err := rows.Scan(&v.ID, &v.Day, &v.ProvinceID, &v.Date, &v.TotalVaccinationTarget,
			&v.FirstVaccinationReceived, &v.SecondVaccinationReceived,
			&v.CumulativeFirstVaccinationReceived, &v.CumulativeSecondVaccinationReceived,
			&v.HealthWorkerVaccinationTarget, &v.HealthWorkerFirstVaccinationReceived, &v.HealthWorkerSecondVaccinationReceived,
			&v.CumulativeHealthWorkerFirstVaccinationReceived, &v.CumulativeHealthWorkerSecondVaccinationReceived,
			&v.ElderlyVaccinationTarget, &v.ElderlyFirstVaccinationReceived, &v.ElderlySecondVaccinationReceived,
			&v.CumulativeElderlyFirstVaccinationReceived, &v.CumulativeElderlySecondVaccinationReceived,
			&v.PublicOfficerVaccinationTarget, &v.PublicOfficerFirstVaccinationReceived, &v.PublicOfficerSecondVaccinationReceived,
			&v.CumulativePublicOfficerFirstVaccinationReceived, &v.CumulativePublicOfficerSecondVaccinationReceived,
			&v.PublicVaccinationTarget, &v.PublicFirstVaccinationReceived, &v.PublicSecondVaccinationReceived,
			&v.CumulativePublicFirstVaccinationReceived, &v.CumulativePublicSecondVaccinationReceived,
			&v.TeenagerVaccinationTarget, &v.TeenagerFirstVaccinationReceived, &v.TeenagerSecondVaccinationReceived,
			&v.CumulativeTeenagerFirstVaccinationReceived, &v.CumulativeTeenagerSecondVaccinationReceived,
		); err != nil {
			return nil, fmt.Errorf("failed to scan province vaccine: %w", err)
		}
		vaccines = append(vaccines, v)
	}
	return vaccines, rows.Err()
}

// GetVaccineLocations returns vaccination centers for SulTeng regencies
func (r *VaccinationRepository) GetVaccineLocations(provinceID int) ([]models.VaccineLocation, error) {
	query := `SELECT id, regency_id, name, address, operational_time,
		is_first_vaccination, is_second_vaccination,
		daily_vaccination_quota, vaccination_stock_remaining, notes
		FROM vaccine_locations WHERE regency_id LIKE ? ORDER BY name`

	likeParam := fmt.Sprintf("%d%%", provinceID)
	rows, err := r.db.Query(query, likeParam)
	if err != nil {
		return nil, fmt.Errorf("failed to query vaccine locations: %w", err)
	}
	defer rows.Close()

	var locations []models.VaccineLocation
	for rows.Next() {
		var l models.VaccineLocation
		if err := rows.Scan(&l.ID, &l.RegencyID, &l.Name, &l.Address, &l.OperationalTime,
			&l.IsFirstVaccination, &l.IsSecondVaccination,
			&l.DailyVaccinationQuota, &l.VaccinationStockRemaining, &l.Notes); err != nil {
			return nil, fmt.Errorf("failed to scan vaccine location: %w", err)
		}
		locations = append(locations, l)
	}
	return locations, rows.Err()
}
