package repository

import (
	"fmt"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

// TaskForceRepositoryInterface defines the contract for task force repository operations
type TaskForceRepositoryInterface interface {
	GetAllByProvinceID(provinceID int) ([]models.TaskForceByRegency, error)
}

// TaskForceRepository handles database operations for task forces
type TaskForceRepository struct {
	db *database.DB
}

// NewTaskForceRepository creates a new TaskForceRepository
func NewTaskForceRepository(db *database.DB) *TaskForceRepository {
	return &TaskForceRepository{db: db}
}

// GetAllByProvinceID returns all task forces grouped by regency for a province
func (r *TaskForceRepository) GetAllByProvinceID(provinceID int) ([]models.TaskForceByRegency, error) {
	// Get regencies first
	regQuery := `SELECT id, name FROM regencies WHERE province_id = ? ORDER BY name`
	regRows, err := r.db.Query(regQuery, provinceID)
	if err != nil {
		return nil, fmt.Errorf("failed to query regencies: %w", err)
	}
	defer regRows.Close()

	var result []models.TaskForceByRegency
	for regRows.Next() {
		var reg models.TaskForceByRegency
		if err := regRows.Scan(&reg.RegencyID, &reg.RegencyName); err != nil {
			return nil, fmt.Errorf("failed to scan regency: %w", err)
		}
		result = append(result, reg)
	}
	if err := regRows.Err(); err != nil {
		return nil, err
	}

	// Get task forces with contacts for each regency
	for i := range result {
		tfQuery := `SELECT tf.id, tf.regency_id, tf.name FROM task_forces tf WHERE tf.regency_id = ?`
		tfRows, err := r.db.Query(tfQuery, result[i].RegencyID)
		if err != nil {
			return nil, fmt.Errorf("failed to query task forces: %w", err)
		}

		var taskForces []models.TaskForce
		for tfRows.Next() {
			var tf models.TaskForce
			if err := tfRows.Scan(&tf.ID, &tf.RegencyID, &tf.Name); err != nil {
				tfRows.Close()
				return nil, fmt.Errorf("failed to scan task force: %w", err)
			}

			// Get contacts for this task force
			cQuery := `SELECT c.id, c.contact_type_id, c.contact, ct.name, ct.icon
				FROM contacts c
				JOIN contact_types ct ON c.contact_type_id = ct.id
				WHERE c.contactable_type = 'App\\Models\\TaskForce' AND c.contactable_id = ?`
			cRows, err := r.db.Query(cQuery, tf.ID)
			if err != nil {
				tfRows.Close()
				return nil, fmt.Errorf("failed to query contacts: %w", err)
			}

			var contacts []models.Contact
			for cRows.Next() {
				var c models.Contact
				if err := cRows.Scan(&c.ID, &c.ContactTypeID, &c.Contact, &c.ContactTypeName, &c.ContactTypeIcon); err != nil {
					cRows.Close()
					tfRows.Close()
					return nil, fmt.Errorf("failed to scan contact: %w", err)
				}
				contacts = append(contacts, c)
			}
			cRows.Close()

			tf.Contacts = contacts
			taskForces = append(taskForces, tf)
		}
		tfRows.Close()

		result[i].TaskForces = taskForces
	}

	return result, nil
}
