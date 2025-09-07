package repository

import (
	"database/sql"
	"fmt"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

type ProvinceRepository interface {
	GetAll() ([]models.Province, error)
	GetByID(id string) (*models.Province, error)
}

type provinceRepository struct {
	db *database.DB
}

func NewProvinceRepository(db *database.DB) ProvinceRepository {
	return &provinceRepository{db: db}
}

func (r *provinceRepository) GetAll() ([]models.Province, error) {
	query := `SELECT id, name FROM provinces ORDER BY name`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query provinces: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			fmt.Printf("Error closing rows: %v\n", err)
		}
	}()

	var provinces []models.Province
	for rows.Next() {
		var p models.Province
		err := rows.Scan(&p.ID, &p.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan province: %w", err)
		}
		provinces = append(provinces, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return provinces, nil
}

func (r *provinceRepository) GetByID(id string) (*models.Province, error) {
	query := `SELECT id, name FROM provinces WHERE id = ?`

	var p models.Province
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get province by ID: %w", err)
	}

	return &p, nil
}