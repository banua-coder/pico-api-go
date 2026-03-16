package repository

import (
	"database/sql/driver"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var regencyCols = []string{"id", "province_id", "name", "created_at", "updated_at"}

func regencyRow() []driver.Value {
	now := time.Now()
	return []driver.Value{7201, 72, "Kabupaten Banggai", &now, &now}
}

func setupRegencyRepo(t *testing.T) (*RegencyRepository, sqlmock.Sqlmock) {
	db, mock := setupMockDB(t)
	return NewRegencyRepository(db), mock
}

func TestRegencyRepository_GetAll(t *testing.T) {
	repo, mock := setupRegencyRepo(t)

	mock.ExpectQuery(`SELECT id, province_id`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows(regencyCols).AddRow(regencyRow()...))

	result, err := repo.GetAll(72)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Kabupaten Banggai", result[0].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegencyRepository_GetAll_Error(t *testing.T) {
	repo, mock := setupRegencyRepo(t)

	mock.ExpectQuery(`SELECT id, province_id`).
		WithArgs(72).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetAll(72)
	assert.Error(t, err)
}

func TestRegencyRepository_GetPaginated(t *testing.T) {
	repo, mock := setupRegencyRepo(t)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM regencies`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	mock.ExpectQuery(`SELECT id, province_id`).
		WithArgs(72, 10, 0).
		WillReturnRows(sqlmock.NewRows(regencyCols).AddRow(regencyRow()...))

	result, total, err := repo.GetPaginated(72, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, 1, total)
	assert.Len(t, result, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegencyRepository_GetPaginated_CountError(t *testing.T) {
	repo, mock := setupRegencyRepo(t)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM regencies`).
		WithArgs(72).
		WillReturnError(errors.New("db error"))

	_, _, err := repo.GetPaginated(72, 10, 0)
	assert.Error(t, err)
}

func TestRegencyRepository_GetPaginated_QueryError(t *testing.T) {
	repo, mock := setupRegencyRepo(t)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM regencies`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	mock.ExpectQuery(`SELECT id, province_id`).
		WithArgs(72, 10, 0).
		WillReturnError(errors.New("db error"))

	_, _, err := repo.GetPaginated(72, 10, 0)
	assert.Error(t, err)
}

func TestRegencyRepository_GetByID(t *testing.T) {
	repo, mock := setupRegencyRepo(t)

	mock.ExpectQuery(`SELECT id, province_id`).
		WithArgs(7201).
		WillReturnRows(sqlmock.NewRows(regencyCols).AddRow(regencyRow()...))

	result, err := repo.GetByID(7201)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Kabupaten Banggai", result.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRegencyRepository_GetByID_NotFound(t *testing.T) {
	repo, mock := setupRegencyRepo(t)

	mock.ExpectQuery(`SELECT id, province_id`).
		WithArgs(9999).
		WillReturnRows(sqlmock.NewRows(regencyCols))

	result, err := repo.GetByID(9999)
	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestRegencyRepository_GetByID_Error(t *testing.T) {
	repo, mock := setupRegencyRepo(t)

	mock.ExpectQuery(`SELECT id, province_id`).
		WithArgs(7201).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetByID(7201)
	assert.Error(t, err)
}
