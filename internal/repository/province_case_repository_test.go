package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestProvinceCaseRepository_GetAll(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	repo := NewProvinceCaseRepository(db)

	now := time.Now()
	rt := 1.1

	rows := sqlmock.NewRows([]string{
		"id", "day", "province_id", "positive", "recovered", "deceased",
		"person_under_observation", "finished_person_under_observation",
		"person_under_supervision", "finished_person_under_supervision",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"cumulative_person_under_observation", "cumulative_finished_persoon_under_observation",
		"cumulative_person_under_supervision", "cumulative_finished_person_under_supervision",
		"rt", "rt_upper", "rt_lower", "date", "name",
	}).AddRow(1, 1, "11", 50, 40, 2, 10, 8, 5, 3, 500, 400, 20, 100, 80, 50, 30, rt, nil, nil, now, "Aceh")

	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	cases, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, int64(1), cases[0].ID)
	assert.Equal(t, "11", cases[0].ProvinceID)
	assert.Equal(t, int64(50), cases[0].Positive)
	assert.NotNil(t, cases[0].Province)
	assert.Equal(t, "Aceh", cases[0].Province.Name)
	assert.Equal(t, &rt, cases[0].Rt)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByProvinceID(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	repo := NewProvinceCaseRepository(db)

	provinceID := "11"
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "day", "province_id", "positive", "recovered", "deceased",
		"person_under_observation", "finished_person_under_observation",
		"person_under_supervision", "finished_person_under_supervision",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"cumulative_person_under_observation", "cumulative_finished_persoon_under_observation",
		"cumulative_person_under_supervision", "cumulative_finished_person_under_supervision",
		"rt", "rt_upper", "rt_lower", "date", "name",
	}).AddRow(1, 1, provinceID, 50, 40, 2, 10, 8, 5, 3, 500, 400, 20, 100, 80, 50, 30, nil, nil, nil, now, "Aceh")

	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WithArgs(provinceID).
		WillReturnRows(rows)

	cases, err := repo.GetByProvinceID(provinceID)

	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, provinceID, cases[0].ProvinceID)
	assert.Nil(t, cases[0].Rt)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByProvinceIDAndDateRange(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	repo := NewProvinceCaseRepository(db)

	provinceID := "11"
	startDate := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "day", "province_id", "positive", "recovered", "deceased",
		"person_under_observation", "finished_person_under_observation",
		"person_under_supervision", "finished_person_under_supervision",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"cumulative_person_under_observation", "cumulative_finished_persoon_under_observation",
		"cumulative_person_under_supervision", "cumulative_finished_person_under_supervision",
		"rt", "rt_upper", "rt_lower", "date", "name",
	}).AddRow(1, 1, provinceID, 50, 40, 2, 10, 8, 5, 3, 500, 400, 20, 100, 80, 50, 30, nil, nil, nil, now, "Aceh")

	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WithArgs(provinceID, startDate, endDate).
		WillReturnRows(rows)

	cases, err := repo.GetByProvinceIDAndDateRange(provinceID, startDate, endDate)

	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, provinceID, cases[0].ProvinceID)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetLatestByProvinceID(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	repo := NewProvinceCaseRepository(db)

	provinceID := "11"
	now := time.Now()
	rt := 1.2

	rows := sqlmock.NewRows([]string{
		"id", "day", "province_id", "positive", "recovered", "deceased",
		"person_under_observation", "finished_person_under_observation",
		"person_under_supervision", "finished_person_under_supervision",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"cumulative_person_under_observation", "cumulative_finished_persoon_under_observation",
		"cumulative_person_under_supervision", "cumulative_finished_person_under_supervision",
		"rt", "rt_upper", "rt_lower", "date", "name",
	}).AddRow(1, 1, provinceID, 50, 40, 2, 10, 8, 5, 3, 500, 400, 20, 100, 80, 50, 30, rt, nil, nil, now, "Aceh")

	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WithArgs(provinceID).
		WillReturnRows(rows)

	provinceCase, err := repo.GetLatestByProvinceID(provinceID)

	assert.NoError(t, err)
	assert.NotNil(t, provinceCase)
	assert.Equal(t, provinceID, provinceCase.ProvinceID)
	assert.Equal(t, &rt, provinceCase.Rt)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetLatestByProvinceID_NotFound(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.Close()

	repo := NewProvinceCaseRepository(db)

	provinceID := "999"

	rows := sqlmock.NewRows([]string{
		"id", "day", "province_id", "positive", "recovered", "deceased",
		"person_under_observation", "finished_person_under_observation",
		"person_under_supervision", "finished_person_under_supervision",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"cumulative_person_under_observation", "cumulative_finished_persoon_under_observation",
		"cumulative_person_under_supervision", "cumulative_finished_person_under_supervision",
		"rt", "rt_upper", "rt_lower", "date", "name",
	})

	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WithArgs(provinceID).
		WillReturnRows(rows)

	provinceCase, err := repo.GetLatestByProvinceID(provinceID)

	assert.NoError(t, err)
	assert.Nil(t, provinceCase)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}