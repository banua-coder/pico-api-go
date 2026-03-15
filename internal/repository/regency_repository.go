package repository

import (
	"log"
	"database/sql"
	"fmt"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

// RegencyRepositoryInterface defines the contract for regency repository operations
type RegencyRepositoryInterface interface {
	GetAll(provinceID int) ([]models.Regency, error)
	GetByID(id int) (*models.Regency, error)
}

// RegencyRepository handles database operations for regencies
type RegencyRepository struct {
	db *database.DB
}

// NewRegencyRepository creates a new RegencyRepository
func NewRegencyRepository(db *database.DB) *RegencyRepository {
	return &RegencyRepository{db: db}
}

// GetAll returns all regencies for a province
func (r *RegencyRepository) GetAll(provinceID int) ([]models.Regency, error) {
	query := `SELECT id, province_id, name, created_at, updated_at FROM regencies WHERE province_id = ? ORDER BY name`

	rows, err := r.db.Query(query, provinceID)
	if err != nil {
		return nil, fmt.Errorf("failed to query regencies: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var regencies []models.Regency
	for rows.Next() {
		var reg models.Regency
		if err := rows.Scan(&reg.ID, &reg.ProvinceID, &reg.Name, &reg.CreatedAt, &reg.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan regency: %w", err)
		}
		regencies = append(regencies, reg)
	}
	return regencies, rows.Err()
}

// GetByID returns a single regency by ID
func (r *RegencyRepository) GetByID(id int) (*models.Regency, error) {
	query := `SELECT id, province_id, name, created_at, updated_at FROM regencies WHERE id = ?`

	var reg models.Regency
	err := r.db.QueryRow(query, id).Scan(&reg.ID, &reg.ProvinceID, &reg.Name, &reg.CreatedAt, &reg.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get regency: %w", err)
	}
	return &reg, nil
}
