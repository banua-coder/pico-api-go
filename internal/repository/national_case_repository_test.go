package repository

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/banua-coder/pico-api-go/pkg/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockDB(t *testing.T) (*database.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	
	return &database.DB{DB: db}, mock
}

func TestNationalCaseRepository_GetAll(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewNationalCaseRepository(db)

	now := time.Now()
	rt := 1.2
	rtUpper := 1.5
	rtLower := 0.9

	rows := sqlmock.NewRows([]string{
		"id", "day", "date", "positive", "recovered", "deceased",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"rt", "rt_upper", "rt_lower",
	}).AddRow(1, 1, now, 100, 80, 5, 1000, 800, 50, rt, rtUpper, rtLower)

	mock.ExpectQuery(`SELECT id, day, date, positive, recovered, deceased,`).
		WillReturnRows(rows)

	cases, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, int64(1), cases[0].ID)
	assert.Equal(t, int64(100), cases[0].Positive)
	assert.Equal(t, &rt, cases[0].Rt)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestNationalCaseRepository_GetByDateRange(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewNationalCaseRepository(db)

	startDate := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "day", "date", "positive", "recovered", "deceased",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"rt", "rt_upper", "rt_lower",
	}).AddRow(1, 1, now, 100, 80, 5, 1000, 800, 50, nil, nil, nil)

	mock.ExpectQuery(`SELECT id, day, date, positive, recovered, deceased,`).
		WithArgs(startDate, endDate).
		WillReturnRows(rows)

	cases, err := repo.GetByDateRange(startDate, endDate)

	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, int64(1), cases[0].ID)
	assert.Nil(t, cases[0].Rt)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestNationalCaseRepository_GetLatest(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewNationalCaseRepository(db)

	now := time.Now()
	rt := 1.1

	rows := sqlmock.NewRows([]string{
		"id", "day", "date", "positive", "recovered", "deceased",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"rt", "rt_upper", "rt_lower",
	}).AddRow(1, 1, now, 100, 80, 5, 1000, 800, 50, rt, nil, nil)

	mock.ExpectQuery(`SELECT id, day, date, positive, recovered, deceased,`).
		WillReturnRows(rows)

	nationalCase, err := repo.GetLatest()

	assert.NoError(t, err)
	assert.NotNil(t, nationalCase)
	assert.Equal(t, int64(1), nationalCase.ID)
	assert.Equal(t, &rt, nationalCase.Rt)
	assert.Nil(t, nationalCase.RtUpper)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestNationalCaseRepository_GetLatest_NotFound(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewNationalCaseRepository(db)

	mock.ExpectQuery(`SELECT id, day, date, positive, recovered, deceased,`).
		WillReturnError(sql.ErrNoRows)

	nationalCase, err := repo.GetLatest()

	assert.NoError(t, err)
	assert.Nil(t, nationalCase)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestNationalCaseRepository_GetByDay(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewNationalCaseRepository(db)

	now := time.Now()
	day := int64(1)

	rows := sqlmock.NewRows([]string{
		"id", "day", "date", "positive", "recovered", "deceased",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"rt", "rt_upper", "rt_lower",
	}).AddRow(1, day, now, 100, 80, 5, 1000, 800, 50, nil, nil, nil)

	mock.ExpectQuery(`SELECT id, day, date, positive, recovered, deceased,`).
		WithArgs(day).
		WillReturnRows(rows)

	nationalCase, err := repo.GetByDay(day)

	assert.NoError(t, err)
	assert.NotNil(t, nationalCase)
	assert.Equal(t, day, nationalCase.Day)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestNationalCaseRepository_GetByDay_NotFound(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewNationalCaseRepository(db)

	day := int64(999)

	mock.ExpectQuery(`SELECT id, day, date, positive, recovered, deceased,`).
		WithArgs(day).
		WillReturnError(sql.ErrNoRows)

	nationalCase, err := repo.GetByDay(day)

	assert.NoError(t, err)
	assert.Nil(t, nationalCase)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}
