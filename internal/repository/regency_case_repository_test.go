package repository

import (
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// cols for GetByRegencyID (23 cols: includes rt fields)
var regencyCaseCols = []string{
	"id", "day", "regency_id", "positive", "recovered", "deceased",
	"person_under_observation", "finished_person_under_observation",
	"person_under_supervision", "finished_person_under_supervision",
	"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
	"cumulative_person_under_observation", "cumulative_finished_person_under_observation",
	"cumulative_person_under_supervision", "cumulative_finished_person_under_supervision",
	"rt", "rt_upper", "rt_lower",
	"date", "reg_id", "reg_name",
}

// cols for GetLatestByProvinceID (cumulative order differs from GetByRegencyID)
var regencyCaseLatestCols = []string{
	"id", "day", "regency_id", "positive", "recovered", "deceased",
	"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
	"person_under_observation", "finished_person_under_observation",
	"person_under_supervision", "finished_person_under_supervision",
	"cumulative_person_under_observation", "cumulative_finished_person_under_observation",
	"cumulative_person_under_supervision", "cumulative_finished_person_under_supervision",
	"rt", "rt_upper", "rt_lower",
	"date", "reg_id", "reg_name",
}

func regencyCaseRow(regencyID int) []driver.Value {
	now := time.Now()
	return []driver.Value{
		1, 1, regencyID, int64(10), int64(8), int64(1),
		int64(0), int64(0), int64(0), int64(0),
		int64(100), int64(80), int64(10),
		int64(0), int64(0), int64(0), int64(0),
		nil, nil, nil,
		now, regencyID, "Kabupaten Test",
	}
}

func regencyCaseLatestRow(regencyID int) []driver.Value {
	now := time.Now()
	return []driver.Value{
		1, 1, regencyID, int64(10), int64(8), int64(1),
		int64(100), int64(80), int64(10),
		int64(0), int64(0), int64(0), int64(0),
		int64(0), int64(0), int64(0), int64(0),
		1.2, 1.5, 0.9,
		now, regencyID, "Kabupaten Test",
	}
}

func setupRegencyCaseRepo(t *testing.T) (*RegencyCaseRepository, sqlmock.Sqlmock) {
	db, mock := setupMockDB(t)
	return NewRegencyCaseRepository(db), mock
}

func TestRegencyCaseRepository_GetByRegencyID(t *testing.T) {
	repo, mock := setupRegencyCaseRepo(t)

	mock.ExpectQuery(`SELECT rc.id`).
		WithArgs(7201).
		WillReturnRows(sqlmock.NewRows(regencyCaseCols).AddRow(regencyCaseRow(7201)...))

	result, err := repo.GetByRegencyID(7201)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, 7201, result[0].RegencyID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegencyCaseRepository_GetByRegencyID_Error(t *testing.T) {
	repo, mock := setupRegencyCaseRepo(t)

	mock.ExpectQuery(`SELECT rc.id`).
		WithArgs(7201).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetByRegencyID(7201)
	assert.Error(t, err)
}

func TestRegencyCaseRepository_GetByRegencyID_Empty(t *testing.T) {
	repo, mock := setupRegencyCaseRepo(t)

	mock.ExpectQuery(`SELECT rc.id`).
		WithArgs(9999).
		WillReturnRows(sqlmock.NewRows(regencyCaseCols))

	result, err := repo.GetByRegencyID(9999)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestRegencyCaseRepository_GetLatest_AllProvinces(t *testing.T) {
	repo, mock := setupRegencyCaseRepo(t)

	// No province filter and no args: regencies from several provinces come back.
	mock.ExpectQuery(`SELECT rc.id`).
		WithArgs().
		WillReturnRows(sqlmock.NewRows(regencyCaseLatestCols).
			AddRow(regencyCaseLatestRow(7201)...).
			AddRow(regencyCaseLatestRow(1101)...))

	result, err := repo.GetLatest()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegencyCaseRepository_GetLatest_UsesNonCorrelatedJoin(t *testing.T) {
	repo, mock := setupRegencyCaseRepo(t)
	mock.ExpectQuery(`JOIN \(SELECT regency_id, MAX\(day\) AS day FROM regency_cases\s+WHERE \(positive > 0 OR recovered > 0 OR deceased > 0\)\s+GROUP BY regency_id\) m ON m.regency_id = rc.regency_id AND m.day = rc.day`).
		WithArgs().
		WillReturnRows(sqlmock.NewRows(regencyCaseLatestCols).AddRow(regencyCaseLatestRow(7201)...))
	got, err := repo.GetLatest()
	assert.NoError(t, err)
	assert.Len(t, got, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegencyCaseRepository_GetLatest_Error(t *testing.T) {
	repo, mock := setupRegencyCaseRepo(t)
	mock.ExpectQuery(`SELECT rc.id`).WillReturnError(errors.New("db error"))
	_, err := repo.GetLatest()
	assert.Error(t, err)
}
