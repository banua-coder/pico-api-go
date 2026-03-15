package models

import "time"

// Regency represents a regency/city (kabupaten/kota)
type Regency struct {
	ID         int        `json:"id" db:"id"`
	ProvinceID int        `json:"province_id" db:"province_id"`
	Name       string     `json:"name" db:"name"`
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
