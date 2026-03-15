package models

import "time"

// Hospital represents a hospital in the system
type Hospital struct {
	ID           int64            `json:"id" db:"id"`
	RegencyID    int              `json:"regency_id" db:"regency_id"`
	Name         string           `json:"name" db:"name"`
	HospitalCode *string          `json:"hospital_code" db:"hospital_code"`
	Address      string           `json:"address" db:"address"`
	Latitude     float64          `json:"latitude" db:"latitude"`
	Longitude    float64          `json:"longitude" db:"longitude"`
	CreatedAt    *time.Time       `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    *time.Time       `json:"updated_at,omitempty" db:"updated_at"`
	Contacts     []Contact        `json:"contacts,omitempty"`
	Beds         []HospitalBed    `json:"beds,omitempty"`
	IGDCount     int              `json:"igd_count" db:"igd_count"`
}

// HospitalBedType represents a type of hospital bed
type HospitalBedType struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// HospitalBed represents hospital bed availability
type HospitalBed struct {
	ID                int64  `json:"id" db:"id"`
	HospitalID        int64  `json:"hospital_id" db:"hospital_id"`
	HospitalBedTypeID int64  `json:"hospital_bed_type_id" db:"hospital_bed_type_id"`
	BedTypeName       string `json:"bed_type_name,omitempty"`
	Available         int    `json:"available" db:"available"`
	Total             int    `json:"total" db:"total"`
}
