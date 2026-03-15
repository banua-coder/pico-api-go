package repository

import (
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var genderCols = []string{
	"id", "day", "province_id",
	"positive_male", "positive_female", "pdp_male", "pdp_female",
	"positive_male_0_14", "positive_male_15_19", "positive_male_20_24", "positive_male_25_49", "positive_male_50_54", "positive_male_55",
	"positive_female_0_14", "positive_female_15_19", "positive_female_20_24", "positive_female_25_49", "positive_female_50_54", "positive_female_55",
	"pdp_male_0_14", "pdp_male_15_19", "pdp_male_20_24", "pdp_male_25_49", "pdp_male_50_54", "pdp_male_55",
	"pdp_female_0_14", "pdp_female_15_19", "pdp_female_20_24", "pdp_female_25_49", "pdp_female_50_54", "pdp_female_55",
}

var testTypeCols = []string{"id", "key", "name", "sample", "duration", "is_recommended"}
var provinceTestCols = []string{"id", "test_type_id", "day", "province_id", "date_from", "process", "invalid", "positive", "negative",
	"tt_id", "tt_key", "tt_name", "tt_sample", "tt_duration", "tt_is_recommended"}

func setupProvinceStatsRepo(t *testing.T) (*ProvinceStatsRepository, sqlmock.Sqlmock) {
	db, mock := setupMockDB(t)
	return NewProvinceStatsRepository(db), mock
}

func genderRow() []driver.Value {
	vals := []driver.Value{int64(1), int64(100), 72, 50, 40, 10, 8}
	for i := 0; i < 24; i++ {
		vals = append(vals, 0)
	}
	return vals
}

func TestProvinceStatsRepository_GetGenderCases(t *testing.T) {
	repo, mock := setupProvinceStatsRepo(t)

	mock.ExpectQuery(`SELECT id, day, province_id`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows(genderCols).AddRow(genderRow()...))

	cases, err := repo.GetGenderCases(72)
	assert.NoError(t, err)
	assert.Len(t, cases, 1)
	assert.Equal(t, 72, cases[0].ProvinceID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceStatsRepository_GetGenderCases_Empty(t *testing.T) {
	repo, mock := setupProvinceStatsRepo(t)

	mock.ExpectQuery(`SELECT id, day, province_id`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows(genderCols))

	cases, err := repo.GetGenderCases(72)
	assert.NoError(t, err)
	assert.Empty(t, cases)
}

func TestProvinceStatsRepository_GetGenderCases_Error(t *testing.T) {
	repo, mock := setupProvinceStatsRepo(t)

	mock.ExpectQuery(`SELECT id, day, province_id`).
		WithArgs(72).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetGenderCases(72)
	assert.Error(t, err)
}

func TestProvinceStatsRepository_GetLatestGenderCase(t *testing.T) {
	repo, mock := setupProvinceStatsRepo(t)

	mock.ExpectQuery(`SELECT id, day, province_id`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows(genderCols).AddRow(genderRow()...))

	result, err := repo.GetLatestGenderCase(72)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 72, result.ProvinceID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceStatsRepository_GetLatestGenderCase_Error(t *testing.T) {
	repo, mock := setupProvinceStatsRepo(t)

	mock.ExpectQuery(`SELECT id, day, province_id`).
		WithArgs(72).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetLatestGenderCase(72)
	assert.Error(t, err)
}

func TestProvinceStatsRepository_GetTests(t *testing.T) {
	repo, mock := setupProvinceStatsRepo(t)

	mock.ExpectQuery(`SELECT pt.id`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows(provinceTestCols).
			AddRow(1, 1, 100, 72, time.Now(), 50, 5, 30, 15, 1, "pcr", "PCR", "Swab", "1-2 hari", true))

	tests, err := repo.GetTests(72)
	assert.NoError(t, err)
	assert.Len(t, tests, 1)
	assert.NotNil(t, tests[0].TestType)
	assert.Equal(t, "PCR", tests[0].TestType.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceStatsRepository_GetTests_Error(t *testing.T) {
	repo, mock := setupProvinceStatsRepo(t)

	mock.ExpectQuery(`SELECT pt.id`).
		WithArgs(72).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetTests(72)
	assert.Error(t, err)
}

func TestProvinceStatsRepository_GetTestTypes(t *testing.T) {
	repo, mock := setupProvinceStatsRepo(t)

	mock.ExpectQuery(`SELECT id`).
		WillReturnRows(sqlmock.NewRows(testTypeCols).
			AddRow(1, "pcr", "PCR", "Swab", "1-2 hari", true).
			AddRow(2, "antigen", "Antigen", "Swab", "30 menit", false))

	types, err := repo.GetTestTypes()
	assert.NoError(t, err)
	assert.Len(t, types, 2)
	assert.Equal(t, "pcr", types[0].Key)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceStatsRepository_GetTestTypes_Error(t *testing.T) {
	repo, mock := setupProvinceStatsRepo(t)

	mock.ExpectQuery(`SELECT id`).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetTestTypes()
	assert.Error(t, err)
}
