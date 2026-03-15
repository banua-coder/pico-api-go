package repository

import (
	"database/sql"
	"fmt"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

// HospitalRepositoryInterface defines the contract for hospital repository operations
type HospitalRepositoryInterface interface {
	GetAll(provinceID int) ([]models.Hospital, error)
	GetByCode(code string) (*models.Hospital, error)
}

// HospitalRepository handles database operations for hospitals
type HospitalRepository struct {
	db *database.DB
}

// NewHospitalRepository creates a new HospitalRepository
func NewHospitalRepository(db *database.DB) *HospitalRepository {
	return &HospitalRepository{db: db}
}

// GetAll returns all hospitals for a province (regency_id LIKE provinceID%)
func (r *HospitalRepository) GetAll(provinceID int) ([]models.Hospital, error) {
	query := `SELECT h.id, h.regency_id, h.name, h.hospital_code, h.address, h.latitude, h.longitude,
		COALESCE((SELECT available FROM hospital_beds WHERE hospital_id = h.id AND hospital_bed_type_id = 1 LIMIT 1), 0) as igd_count
		FROM hospitals h
		WHERE h.regency_id LIKE ?
		ORDER BY h.name`

	likeParam := fmt.Sprintf("%d%%", provinceID)
	rows, err := r.db.Query(query, likeParam)
	if err != nil {
		return nil, fmt.Errorf("failed to query hospitals: %w", err)
	}
	defer rows.Close() //nolint:errcheck

	var hospitals []models.Hospital
	for rows.Next() {
		var h models.Hospital
		if err := rows.Scan(&h.ID, &h.RegencyID, &h.Name, &h.HospitalCode, &h.Address,
			&h.Latitude, &h.Longitude, &h.IGDCount); err != nil {
			return nil, fmt.Errorf("failed to scan hospital: %w", err)
		}
		hospitals = append(hospitals, h)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Load contacts and beds for each hospital
	for i := range hospitals {
		contacts, err := r.getContacts("App\\Models\\Hospital", hospitals[i].ID)
		if err != nil {
			return nil, err
		}
		hospitals[i].Contacts = contacts

		beds, err := r.getBeds(hospitals[i].ID)
		if err != nil {
			return nil, err
		}
		hospitals[i].Beds = beds
	}

	return hospitals, nil
}

// GetByCode returns a hospital by its code
func (r *HospitalRepository) GetByCode(code string) (*models.Hospital, error) {
	query := `SELECT h.id, h.regency_id, h.name, h.hospital_code, h.address, h.latitude, h.longitude,
		COALESCE((SELECT available FROM hospital_beds WHERE hospital_id = h.id AND hospital_bed_type_id = 1 LIMIT 1), 0) as igd_count
		FROM hospitals h
		WHERE h.hospital_code = ?`

	var h models.Hospital
	err := r.db.QueryRow(query, code).Scan(&h.ID, &h.RegencyID, &h.Name, &h.HospitalCode, &h.Address,
		&h.Latitude, &h.Longitude, &h.IGDCount)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get hospital: %w", err)
	}

	contacts, err := r.getContacts("App\\Models\\Hospital", h.ID)
	if err != nil {
		return nil, err
	}
	h.Contacts = contacts

	beds, err := r.getBeds(h.ID)
	if err != nil {
		return nil, err
	}
	h.Beds = beds

	return &h, nil
}

func (r *HospitalRepository) getContacts(contactableType string, contactableID int64) ([]models.Contact, error) {
	query := `SELECT c.id, c.contact_type_id, c.contact, ct.name, ct.icon
		FROM contacts c
		JOIN contact_types ct ON c.contact_type_id = ct.id
		WHERE c.contactable_type = ? AND c.contactable_id = ?`

	rows, err := r.db.Query(query, contactableType, contactableID)
	if err != nil {
		return nil, fmt.Errorf("failed to query contacts: %w", err)
	}
	defer rows.Close() //nolint:errcheck

	var contacts []models.Contact
	for rows.Next() {
		var c models.Contact
		if err := rows.Scan(&c.ID, &c.ContactTypeID, &c.Contact, &c.ContactTypeName, &c.ContactTypeIcon); err != nil {
			return nil, fmt.Errorf("failed to scan contact: %w", err)
		}
		contacts = append(contacts, c)
	}
	return contacts, rows.Err()
}

func (r *HospitalRepository) getBeds(hospitalID int64) ([]models.HospitalBed, error) {
	query := `SELECT hb.id, hb.hospital_id, hb.hospital_bed_type_id, hbt.name, hb.available, hb.total
		FROM hospital_beds hb
		JOIN hospital_bed_types hbt ON hb.hospital_bed_type_id = hbt.id
		WHERE hb.hospital_id = ?`

	rows, err := r.db.Query(query, hospitalID)
	if err != nil {
		return nil, fmt.Errorf("failed to query hospital beds: %w", err)
	}
	defer rows.Close() //nolint:errcheck

	var beds []models.HospitalBed
	for rows.Next() {
		var b models.HospitalBed
		if err := rows.Scan(&b.ID, &b.HospitalID, &b.HospitalBedTypeID, &b.BedTypeName, &b.Available, &b.Total); err != nil {
			return nil, fmt.Errorf("failed to scan hospital bed: %w", err)
		}
		beds = append(beds, b)
	}
	return beds, rows.Err()
}
