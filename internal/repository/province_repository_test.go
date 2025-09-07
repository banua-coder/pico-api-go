package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestProvinceRepository_GetAll(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewProvinceRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow("11", "Aceh").
		AddRow("72", "Sulawesi Tengah").
		AddRow("31", "DKI Jakarta")

	mock.ExpectQuery(`SELECT id, name FROM provinces ORDER BY name`).
		WillReturnRows(rows)

	provinces, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, provinces, 3)
	assert.Equal(t, "11", provinces[0].ID)
	assert.Equal(t, "Aceh", provinces[0].Name)
	assert.Equal(t, "72", provinces[1].ID)
	assert.Equal(t, "Sulawesi Tengah", provinces[1].Name)
	assert.Equal(t, "31", provinces[2].ID)
	assert.Equal(t, "DKI Jakarta", provinces[2].Name)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceRepository_GetAll_Empty(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewProvinceRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name"})

	mock.ExpectQuery(`SELECT id, name FROM provinces ORDER BY name`).
		WillReturnRows(rows)

	provinces, err := repo.GetAll()

	assert.NoError(t, err)
	assert.Len(t, provinces, 0)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceRepository_GetByID(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewProvinceRepository(db)

	provinceID := "11"
	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(provinceID, "Aceh")

	mock.ExpectQuery(`SELECT id, name FROM provinces WHERE id = \?`).
		WithArgs(provinceID).
		WillReturnRows(rows)

	province, err := repo.GetByID(provinceID)

	assert.NoError(t, err)
	assert.NotNil(t, province)
	assert.Equal(t, provinceID, province.ID)
	assert.Equal(t, "Aceh", province.Name)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceRepository_GetByID_NotFound(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewProvinceRepository(db)

	provinceID := "999"

	mock.ExpectQuery(`SELECT id, name FROM provinces WHERE id = \?`).
		WithArgs(provinceID).
		WillReturnError(sql.ErrNoRows)

	province, err := repo.GetByID(provinceID)

	assert.NoError(t, err)
	assert.Nil(t, province)
	
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestProvinceRepository_GetByID_DatabaseError(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		if err := db.Close(); err != nil {
			t.Logf("Error closing database: %v", err)
		}
	}()

	repo := NewProvinceRepository(db)

	provinceID := "11"

	mock.ExpectQuery(`SELECT id, name FROM provinces WHERE id = \?`).
		WithArgs(provinceID).
		WillReturnError(sql.ErrConnDone)

	province, err := repo.GetByID(provinceID)

	assert.Error(t, err)
	assert.Nil(t, province)
	assert.Contains(t, err.Error(), "failed to get province by ID")
	
	assert.NoError(t, mock.ExpectationsWereMet())
}