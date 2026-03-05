package models

// TaskForce represents a COVID-19 task force post (posko/gugus tugas)
type TaskForce struct {
	ID        int64     `json:"id" db:"id"`
	RegencyID int       `json:"regency_id" db:"regency_id"`
	Name      string    `json:"name" db:"name"`
	Contacts  []Contact `json:"contacts,omitempty"`
}

// TaskForceByRegency groups task forces by regency
type TaskForceByRegency struct {
	RegencyID   int         `json:"regency_id"`
	RegencyName string      `json:"regency_name"`
	TaskForces  []TaskForce `json:"task_forces"`
}
