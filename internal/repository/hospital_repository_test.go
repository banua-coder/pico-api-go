package repository

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var hospitalCols = []string{"id", "regency_id", "name", "hospital_code", "address", "latitude", "longitude", "igd_count"}
var contactCols = []string{"id", "contact_type_id", "contact", "name", "icon"}
var bedCols = []string{"id", "hospital_id", "hospital_bed_type_id", "name", "available", "total"}

func setupHospitalRepo(t *testing.T) (*HospitalRepository, sqlmock.Sqlmock) {
	db, mock := setupMockDB(t)
	return NewHospitalRepository(db), mock
}

func expectEmptyContacts(mock sqlmock.Sqlmock, hospitalID int64) {
	mock.ExpectQuery(`SELECT c.id, c.contact_type_id`).
		WithArgs("App\\Models\\Hospital", hospitalID).
		WillReturnRows(sqlmock.NewRows(contactCols))
}

func expectEmptyBeds(mock sqlmock.Sqlmock, hospitalID int64) {
	mock.ExpectQuery(`SELECT hb.id`).
		WithArgs(hospitalID).
		WillReturnRows(sqlmock.NewRows(bedCols))
}

func TestHospitalRepository_GetAll(t *testing.T) {
	repo, mock := setupHospitalRepo(t)
	code := "7201001"

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("72%").
		WillReturnRows(sqlmock.NewRows(hospitalCols).
			AddRow(1, 7201, "RSUD Palu", code, "Jl. Test", 0.1, 119.1, 5))

	expectEmptyContacts(mock, 1)
	expectEmptyBeds(mock, 1)

	hospitals, err := repo.GetAll(72)
	assert.NoError(t, err)
	assert.Len(t, hospitals, 1)
	assert.Equal(t, "RSUD Palu", hospitals[0].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHospitalRepository_GetAll_WithContactsAndBeds(t *testing.T) {
	repo, mock := setupHospitalRepo(t)

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("72%").
		WillReturnRows(sqlmock.NewRows(hospitalCols).
			AddRow(1, 7201, "RSUD Palu", nil, "Jl. Test", 0.1, 119.1, 3))

	mock.ExpectQuery(`SELECT c.id, c.contact_type_id`).
		WithArgs("App\\Models\\Hospital", int64(1)).
		WillReturnRows(sqlmock.NewRows(contactCols).
			AddRow(1, 1, "0451-123456", "Telepon", "phone"))

	mock.ExpectQuery(`SELECT hb.id`).
		WithArgs(int64(1)).
		WillReturnRows(sqlmock.NewRows(bedCols).
			AddRow(1, 1, 1, "ICU", 5, 10))

	hospitals, err := repo.GetAll(72)
	assert.NoError(t, err)
	assert.Len(t, hospitals, 1)
	assert.Len(t, hospitals[0].Contacts, 1)
	assert.Len(t, hospitals[0].Beds, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHospitalRepository_GetAll_QueryError(t *testing.T) {
	repo, mock := setupHospitalRepo(t)

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("72%").
		WillReturnError(errors.New("db error"))

	_, err := repo.GetAll(72)
	assert.Error(t, err)
}

func TestHospitalRepository_GetAll_ContactsError(t *testing.T) {
	repo, mock := setupHospitalRepo(t)

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("72%").
		WillReturnRows(sqlmock.NewRows(hospitalCols).
			AddRow(1, 7201, "RSUD Palu", nil, "Jl. Test", 0.1, 119.1, 0))

	mock.ExpectQuery(`SELECT c.id, c.contact_type_id`).
		WithArgs("App\\Models\\Hospital", int64(1)).
		WillReturnError(errors.New("contact query error"))

	_, err := repo.GetAll(72)
	assert.Error(t, err)
}

func TestHospitalRepository_GetAll_BedsError(t *testing.T) {
	repo, mock := setupHospitalRepo(t)

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("72%").
		WillReturnRows(sqlmock.NewRows(hospitalCols).
			AddRow(1, 7201, "RSUD Palu", nil, "Jl. Test", 0.1, 119.1, 0))

	expectEmptyContacts(mock, 1)
	mock.ExpectQuery(`SELECT hb.id`).
		WithArgs(int64(1)).
		WillReturnError(errors.New("beds query error"))

	_, err := repo.GetAll(72)
	assert.Error(t, err)
}

func TestHospitalRepository_GetByCode(t *testing.T) {
	repo, mock := setupHospitalRepo(t)
	code := "7201001"

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("7201001").
		WillReturnRows(sqlmock.NewRows(hospitalCols).
			AddRow(1, 7201, "RSUD Palu", code, "Jl. Test", 0.1, 119.1, 5))

	expectEmptyContacts(mock, 1)
	expectEmptyBeds(mock, 1)

	hospital, err := repo.GetByCode("7201001")
	assert.NoError(t, err)
	assert.NotNil(t, hospital)
	assert.Equal(t, "RSUD Palu", hospital.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHospitalRepository_GetByCode_NotFound(t *testing.T) {
	repo, mock := setupHospitalRepo(t)

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("invalid").
		WillReturnError(sqlmock.ErrCancelled)

	_, err := repo.GetByCode("invalid")
	assert.Error(t, err)
}

func TestHospitalRepository_GetByCode_DBError(t *testing.T) {
	repo, mock := setupHospitalRepo(t)

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("7201001").
		WillReturnError(errors.New("db error"))

	_, err := repo.GetByCode("7201001")
	assert.Error(t, err)
}

func TestHospitalRepository_GetPaginated(t *testing.T) {
	repo, mock := setupHospitalRepo(t)
	code := "7201001"

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM hospitals`).
		WithArgs("72%").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("72%", 10, 0).
		WillReturnRows(sqlmock.NewRows(hospitalCols).
			AddRow(1, 7201, "RSUD Palu", code, "Jl. Test", 0.1, 119.1, 5))

	expectEmptyContacts(mock, 1)
	expectEmptyBeds(mock, 1)

	hospitals, total, err := repo.GetPaginated(72, 10, 0)
	assert.NoError(t, err)
	assert.Len(t, hospitals, 1)
	assert.Equal(t, 1, total)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHospitalRepository_GetPaginated_CountError(t *testing.T) {
	repo, mock := setupHospitalRepo(t)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM hospitals`).
		WithArgs("72%").
		WillReturnError(errors.New("db error"))

	_, _, err := repo.GetPaginated(72, 10, 0)
	assert.Error(t, err)
}

func TestHospitalRepository_GetPaginated_QueryError(t *testing.T) {
	repo, mock := setupHospitalRepo(t)

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM hospitals`).
		WithArgs("72%").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

	mock.ExpectQuery(`SELECT h.id`).
		WithArgs("72%", 10, 0).
		WillReturnError(errors.New("db error"))

	_, _, err := repo.GetPaginated(72, 10, 0)
	assert.Error(t, err)
}
