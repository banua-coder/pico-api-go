package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/banua-coder/pico-api-go/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestProvinceCaseRepository_GetAll(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewProvinceCaseRepository(db)

	now := time.Now()
	rt := 1.1

	rows := sqlmock.NewRows([]string{
		"id", "day", "province_id", "positive", "recovered", "deceased",
		"person_under_observation", "finished_person_under_observation",
		"person_under_supervision", "finished_person_under_supervision",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"cumulative_person_under_observation", "cumulative_finished_person_under_observation",
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
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewProvinceCaseRepository(db)

	provinceID := "11"
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "day", "province_id", "positive", "recovered", "deceased",
		"person_under_observation", "finished_person_under_observation",
		"person_under_supervision", "finished_person_under_supervision",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"cumulative_person_under_observation", "cumulative_finished_person_under_observation",
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
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

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
		"cumulative_person_under_observation", "cumulative_finished_person_under_observation",
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
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewProvinceCaseRepository(db)

	provinceID := "11"
	now := time.Now()
	rt := 1.2

	rows := sqlmock.NewRows([]string{
		"id", "day", "province_id", "positive", "recovered", "deceased",
		"person_under_observation", "finished_person_under_observation",
		"person_under_supervision", "finished_person_under_supervision",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"cumulative_person_under_observation", "cumulative_finished_person_under_observation",
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
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewProvinceCaseRepository(db)

	provinceID := "999"

	rows := sqlmock.NewRows([]string{
		"id", "day", "province_id", "positive", "recovered", "deceased",
		"person_under_observation", "finished_person_under_observation",
		"person_under_supervision", "finished_person_under_supervision",
		"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
		"cumulative_person_under_observation", "cumulative_finished_person_under_observation",
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

var provinceCaseColumns = []string{
	"id", "day", "province_id", "positive", "recovered", "deceased",
	"person_under_observation", "finished_person_under_observation",
	"person_under_supervision", "finished_person_under_supervision",
	"cumulative_positive", "cumulative_recovered", "cumulative_deceased",
	"cumulative_person_under_observation", "cumulative_finished_person_under_observation",
	"cumulative_person_under_supervision", "cumulative_finished_person_under_supervision",
	"rt", "rt_upper", "rt_lower", "date", "name",
}

func addProvinceCaseRow(rows *sqlmock.Rows, provinceID string, now time.Time) *sqlmock.Rows {
	return rows.AddRow(1, 1, provinceID, 50, 40, 2, 10, 8, 5, 3, 500, 400, 20, 100, 80, 50, 30, nil, nil, nil, now, "Aceh")
}

func TestProvinceCaseRepository_GetAllPaginated(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	now := time.Now()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM province_cases`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(100))
	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), "11", now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	cases, total, err := repo.GetAllPaginated(10, 0)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, 100, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetAllPaginatedSorted(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	now := time.Now()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM province_cases`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(50))
	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), "11", now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	sortParams := utils.SortParams{Field: "date", Order: "asc"}
	cases, total, err := repo.GetAllPaginatedSorted(10, 0, sortParams)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, 50, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByProvinceIDPaginated(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	provinceID := "11"
	now := time.Now()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM province_cases`).
		WithArgs(provinceID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(20))
	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), provinceID, now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	cases, total, err := repo.GetByProvinceIDPaginated(provinceID, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, 20, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByProvinceIDAndDateRangePaginated(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	provinceID := "11"
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM province_cases`).
		WithArgs(provinceID, start, end).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))
	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), provinceID, now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	cases, total, err := repo.GetByProvinceIDAndDateRangePaginated(provinceID, start, end, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, 5, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByDateRange(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), "11", now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WithArgs(start, end).
		WillReturnRows(rows)

	cases, err := repo.GetByDateRange(start, end)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByDateRangePaginated(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM province_cases`).
		WithArgs(start, end).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(15))
	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), "11", now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	cases, total, err := repo.GetByDateRangePaginated(start, end, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, 15, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByProvinceIDSorted(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	provinceID := "11"
	now := time.Now()

	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), provinceID, now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	sortParams := utils.SortParams{Field: "date", Order: "desc"}
	cases, err := repo.GetByProvinceIDSorted(provinceID, sortParams)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByProvinceIDPaginatedSorted(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	provinceID := "11"
	now := time.Now()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM province_cases`).
		WithArgs(provinceID).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(10))
	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), provinceID, now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	sortParams := utils.SortParams{Field: "positive", Order: "desc"}
	cases, total, err := repo.GetByProvinceIDPaginatedSorted(provinceID, 10, 0, sortParams)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, 10, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByProvinceIDAndDateRangeSorted(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	provinceID := "11"
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), provinceID, now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	sortParams := utils.SortParams{Field: "date", Order: "asc"}
	cases, err := repo.GetByProvinceIDAndDateRangeSorted(provinceID, start, end, sortParams)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByProvinceIDAndDateRangePaginatedSorted(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	provinceID := "11"
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM province_cases`).
		WithArgs(provinceID, start, end).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(8))
	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), provinceID, now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	sortParams := utils.SortParams{Field: "date", Order: "asc"}
	cases, total, err := repo.GetByProvinceIDAndDateRangePaginatedSorted(provinceID, start, end, 10, 0, sortParams)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, 8, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByDateRangeSorted(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), "11", now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WithArgs(start, end).
		WillReturnRows(rows)

	sortParams := utils.SortParams{Field: "date", Order: "asc"}
	cases, err := repo.GetByDateRangeSorted(start, end, sortParams)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetByDateRangePaginatedSorted(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	start := time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
	now := time.Now()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM province_cases`).
		WithArgs(start, end).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(12))
	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), "11", now)
	mock.ExpectQuery(`SELECT pc\.id, pc\.day, pc\.province_id`).
		WillReturnRows(rows)

	sortParams := utils.SortParams{Field: "date", Order: "desc"}
	cases, total, err := repo.GetByDateRangePaginatedSorted(start, end, 10, 0, sortParams)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, 12, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetAllSorted_ByProvinceName(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	now := time.Now()

	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), "11", now)
	mock.ExpectQuery(`SELECT pc\.id`).
		WillReturnRows(rows)

	cases, err := repo.GetAllSorted(utils.SortParams{Field: "province_name", Order: "desc"})
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceCaseRepository_GetAllSorted_UnknownField(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing db: %v", err)
		}
	}()
	repo := NewProvinceCaseRepository(db)
	now := time.Now()

	rows := addProvinceCaseRow(sqlmock.NewRows(provinceCaseColumns), "11", now)
	mock.ExpectQuery(`SELECT pc\.id`).
		WillReturnRows(rows)

	cases, err := repo.GetAllSorted(utils.SortParams{Field: "unknown_field", Order: "asc"})
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}
