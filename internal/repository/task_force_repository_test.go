package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func setupTaskForceRepo(t *testing.T) (*TaskForceRepository, sqlmock.Sqlmock) {
	db, mock := setupMockDB(t)
	return NewTaskForceRepository(db), mock
}

func TestTaskForceRepository_GetAllByProvinceID(t *testing.T) {
	repo, mock := setupTaskForceRepo(t)

	mock.ExpectQuery(`SELECT id, name FROM regencies`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(7201, "Banggai"))

	mock.ExpectQuery(`SELECT tf.id`).
		WithArgs(7201).
		WillReturnRows(sqlmock.NewRows([]string{"id", "regency_id", "name"}).
			AddRow(1, 7201, "Satgas COVID Banggai"))

	mock.ExpectQuery(`SELECT c.id`).
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows(contactCols).
			AddRow(1, 1, "0812-0000-0000", "WhatsApp", "whatsapp"))

	result, err := repo.GetAllByProvinceID(72)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Banggai", result[0].RegencyName)
	assert.Len(t, result[0].TaskForces, 1)
	assert.Len(t, result[0].TaskForces[0].Contacts, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTaskForceRepository_GetAllByProvinceID_NoRegencies(t *testing.T) {
	repo, mock := setupTaskForceRepo(t)

	mock.ExpectQuery(`SELECT id, name FROM regencies`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))

	result, err := repo.GetAllByProvinceID(72)
	assert.NoError(t, err)
	assert.Empty(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTaskForceRepository_GetAllByProvinceID_RegencyQueryError(t *testing.T) {
	repo, mock := setupTaskForceRepo(t)

	mock.ExpectQuery(`SELECT id, name FROM regencies`).
		WithArgs(72).
		WillReturnError(errors.New("db error"))

	_, err := repo.GetAllByProvinceID(72)
	assert.Error(t, err)
}

func TestTaskForceRepository_GetAllByProvinceID_TaskForceQueryError(t *testing.T) {
	repo, mock := setupTaskForceRepo(t)

	mock.ExpectQuery(`SELECT id, name FROM regencies`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(7201, "Banggai"))

	mock.ExpectQuery(`SELECT tf.id`).
		WithArgs(7201).
		WillReturnError(errors.New("task force error"))

	_, err := repo.GetAllByProvinceID(72)
	assert.Error(t, err)
}

func TestTaskForceRepository_GetAllByProvinceID_EmptyTaskForces(t *testing.T) {
	repo, mock := setupTaskForceRepo(t)

	mock.ExpectQuery(`SELECT id, name FROM regencies`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(7201, "Banggai"))

	mock.ExpectQuery(`SELECT tf.id`).
		WithArgs(7201).
		WillReturnRows(sqlmock.NewRows([]string{"id", "regency_id", "name"}))

	result, err := repo.GetAllByProvinceID(72)
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Empty(t, result[0].TaskForces)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTaskForceRepository_GetPaginatedByProvinceID(t *testing.T) {
	repo, mock := setupTaskForceRepo(t)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM regencies`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	mock.ExpectQuery(`SELECT id, name FROM regencies`).
		WithArgs(72, 10, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(7201, "Banggai"))

	mock.ExpectQuery(`SELECT tf.id`).
		WithArgs(7201).
		WillReturnRows(sqlmock.NewRows([]string{"id", "regency_id", "name"}).
			AddRow(1, 7201, "Satgas COVID Banggai"))

	mock.ExpectQuery(`SELECT c.id`).
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "contact_type_id", "contact", "name", "icon"}))

	result, total, err := repo.GetPaginatedByProvinceID(72, 10, 0)
	assert.NoError(t, err)
	assert.Equal(t, 1, total)
	assert.Len(t, result, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTaskForceRepository_GetPaginatedByProvinceID_CountError(t *testing.T) {
	repo, mock := setupTaskForceRepo(t)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM regencies`).
		WithArgs(72).
		WillReturnError(errors.New("db error"))

	_, _, err := repo.GetPaginatedByProvinceID(72, 10, 0)
	assert.Error(t, err)
}

func TestTaskForceRepository_GetPaginatedByProvinceID_QueryError(t *testing.T) {
	repo, mock := setupTaskForceRepo(t)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM regencies`).
		WithArgs(72).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	mock.ExpectQuery(`SELECT id, name FROM regencies`).
		WithArgs(72, 10, 0).
		WillReturnError(errors.New("db error"))

	_, _, err := repo.GetPaginatedByProvinceID(72, 10, 0)
	assert.Error(t, err)
}
