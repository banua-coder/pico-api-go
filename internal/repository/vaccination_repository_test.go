package repository

import (
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var nationalVaccineColumns = []string{
	"id", "day", "date", "total_vaccination_target",
	"first_vaccination_received", "second_vaccination_received",
	"cumulative_first_vaccination_received", "cumulative_second_vaccination_received",
	"health_worker_vaccination_target", "health_worker_first_vaccination_received", "health_worker_second_vaccination_received",
	"cumulative_health_worker_first_vaccination_received", "cumulative_health_worker_second_vaccination_received",
	"elderly_vaccination_target", "elderly_first_vaccination_received", "elderly_second_vaccination_received",
	"cumulative_elderly_first_vaccination_received", "cumulative_elderly_second_vaccination_received",
	"public_officer_vaccination_target", "public_officer_first_vaccination_received", "public_officer_second_vaccination_received",
	"cumulative_public_officer_first_vaccination_received", "cumulative_public_officer_second_vaccination_received",
	"public_vaccination_target", "public_first_vaccination_received", "public_second_vaccination_received",
	"cumulative_public_first_vaccination_received", "cumulative_public_second_vaccination_received",
	"teenager_vaccination_target", "teenager_first_vaccination_received", "teenager_second_vaccination_received",
	"cumulative_teenager_first_vaccination_received", "cumulative_teenager_second_vaccination_received",
}

func addNationalVaccineRow(rows *sqlmock.Rows, now time.Time) *sqlmock.Rows {
	vals := []driver.Value{1, 1, now}
	for i := 0; i < 30; i++ {
		vals = append(vals, int64(100))
	}
	return rows.AddRow(vals...)
}

func TestVaccinationRepository_GetNationalVaccinations(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()
	repo := NewVaccinationRepository(db)
	now := time.Now()

	rows := addNationalVaccineRow(sqlmock.NewRows(nationalVaccineColumns), now)
	mock.ExpectQuery(`SELECT id, day, date`).
		WillReturnRows(rows)

	vaccines, err := repo.GetNationalVaccinations()
	assert.NoError(t, err)
	assert.Len(t, vaccines, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVaccinationRepository_GetProvinceVaccinations(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()
	repo := NewVaccinationRepository(db)
	now := time.Now()

	provinceCols := []string{"id", "day", "province_id", "date", "total_vaccination_target",
		"first_vaccination_received", "second_vaccination_received",
		"cumulative_first_vaccination_received", "cumulative_second_vaccination_received",
		"health_worker_vaccination_target", "health_worker_first_vaccination_received", "health_worker_second_vaccination_received",
		"cumulative_health_worker_first_vaccination_received", "cumulative_health_worker_second_vaccination_received",
		"elderly_vaccination_target", "elderly_first_vaccination_received", "elderly_second_vaccination_received",
		"cumulative_elderly_first_vaccination_received", "cumulative_elderly_second_vaccination_received",
		"public_officer_vaccination_target", "public_officer_first_vaccination_received", "public_officer_second_vaccination_received",
		"cumulative_public_officer_first_vaccination_received", "cumulative_public_officer_second_vaccination_received",
		"public_vaccination_target", "public_first_vaccination_received", "public_second_vaccination_received",
		"cumulative_public_first_vaccination_received", "cumulative_public_second_vaccination_received",
		"teenager_vaccination_target", "teenager_first_vaccination_received", "teenager_second_vaccination_received",
		"cumulative_teenager_first_vaccination_received", "cumulative_teenager_second_vaccination_received",
	}

	vals := []driver.Value{1, 1, 72, now}
	for i := 0; i < 30; i++ {
		vals = append(vals, int64(50))
	}
	rows := sqlmock.NewRows(provinceCols).AddRow(vals...)
	mock.ExpectQuery(`SELECT id, day, province_id, date`).
		WithArgs(72).
		WillReturnRows(rows)

	vaccines, err := repo.GetProvinceVaccinations(72)
	assert.NoError(t, err)
	assert.Len(t, vaccines, 1)
	assert.Equal(t, 72, vaccines[0].ProvinceID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestVaccinationRepository_GetVaccineLocations(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()
	repo := NewVaccinationRepository(db)

	cols := []string{"id", "regency_id", "name", "address", "operational_time",
		"is_first_vaccination", "is_second_vaccination",
		"daily_vaccination_quota", "vaccination_stock_remaining", "notes"}
	rows := sqlmock.NewRows(cols).
		AddRow(1, 7201, "Puskesmas A", "Jl. Raya", "08:00-16:00", true, true, 100, 50, "")

	mock.ExpectQuery(`SELECT id, regency_id, name`).
		WithArgs("72%").
		WillReturnRows(rows)

	locs, err := repo.GetVaccineLocations(72)
	assert.NoError(t, err)
	assert.Len(t, locs, 1)
	assert.Equal(t, "Puskesmas A", locs[0].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}
