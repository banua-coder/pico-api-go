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

// cols for GetLatestByProvinceID (20 cols: cumulative order differs, no rt fields)
var regencyCaseLatestCols = []string{
	"id", "day", "regency_id", "positive", "recovered", "deceased",
	"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
	"person_under_observation", "finished_person_under_observation",
	"person_under_supervision", "finished_person_under_supervision",
	"cumulative_person_under_observation", "cumulative_finished_person_under_observation",
	"cumulative_person_under_supervision", "cumulative_finished_person_under_supervision",
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

func TestRegencyCaseRepository_GetLatestByProvinceID(t *testing.T) {
	repo, mock := setupRegencyCaseRepo(t)

	mock.ExpectQuery(`SELECT rc.id`).
		WithArgs("72%").
		WillReturnRows(sqlmock.NewRows(regencyCaseLatestCols).AddRow(regencyCaseLatestRow(7201)...))

	result, err := repo.GetLatestByProvinceID(72)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegencyCaseRepository_GetLatestByProvinceID_Error(t *testing.T) {
	repo, mock := setupRegencyCaseRepo(t)

	mock.ExpectQuery(`SELECT rc.id`).
		WithArgs("72%").
		WillReturnError(errors.New("db error"))

	_, err := repo.GetLatestByProvinceID(72)
	assert.Error(t, err)
}
