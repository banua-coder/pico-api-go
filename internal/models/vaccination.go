package models

import "time"

// NationalVaccine represents national vaccination data
type NationalVaccine struct {
	ID   int64     `json:"id" db:"id"`
	Day  int64     `json:"day" db:"day"`
	Date time.Time `json:"date" db:"date"`

	TotalVaccinationTarget int64 `json:"total_vaccination_target" db:"total_vaccination_target"`

	FirstVaccinationReceived            int64 `json:"first_vaccination_received" db:"first_vaccination_received"`
	SecondVaccinationReceived           int64 `json:"second_vaccination_received" db:"second_vaccination_received"`
	CumulativeFirstVaccinationReceived  int64 `json:"cumulative_first_vaccination_received" db:"cumulative_first_vaccination_received"`
	CumulativeSecondVaccinationReceived int64 `json:"cumulative_second_vaccination_received" db:"cumulative_second_vaccination_received"`

	HealthWorkerVaccinationTarget                    int64 `json:"health_worker_vaccination_target" db:"health_worker_vaccination_target"`
	HealthWorkerFirstVaccinationReceived             int64 `json:"health_worker_first_vaccination_received" db:"health_worker_first_vaccination_received"`
	HealthWorkerSecondVaccinationReceived            int64 `json:"health_worker_second_vaccination_received" db:"health_worker_second_vaccination_received"`
	CumulativeHealthWorkerFirstVaccinationReceived   int64 `json:"cumulative_health_worker_first_vaccination_received" db:"cumulative_health_worker_first_vaccination_received"`
	CumulativeHealthWorkerSecondVaccinationReceived  int64 `json:"cumulative_health_worker_second_vaccination_received" db:"cumulative_health_worker_second_vaccination_received"`

	ElderlyVaccinationTarget                    int64 `json:"elderly_vaccination_target" db:"elderly_vaccination_target"`
	ElderlyFirstVaccinationReceived             int64 `json:"elderly_first_vaccination_received" db:"elderly_first_vaccination_received"`
	ElderlySecondVaccinationReceived            int64 `json:"elderly_second_vaccination_received" db:"elderly_second_vaccination_received"`
	CumulativeElderlyFirstVaccinationReceived   int64 `json:"cumulative_elderly_first_vaccination_received" db:"cumulative_elderly_first_vaccination_received"`
	CumulativeElderlySecondVaccinationReceived  int64 `json:"cumulative_elderly_second_vaccination_received" db:"cumulative_elderly_second_vaccination_received"`

	PublicOfficerVaccinationTarget                    int64 `json:"public_officer_vaccination_target" db:"public_officer_vaccination_target"`
	PublicOfficerFirstVaccinationReceived             int64 `json:"public_officer_first_vaccination_received" db:"public_officer_first_vaccination_received"`
	PublicOfficerSecondVaccinationReceived            int64 `json:"public_officer_second_vaccination_received" db:"public_officer_second_vaccination_received"`
	CumulativePublicOfficerFirstVaccinationReceived   int64 `json:"cumulative_public_officer_first_vaccination_received" db:"cumulative_public_officer_first_vaccination_received"`
	CumulativePublicOfficerSecondVaccinationReceived  int64 `json:"cumulative_public_officer_second_vaccination_received" db:"cumulative_public_officer_second_vaccination_received"`

	PublicVaccinationTarget                    int64 `json:"public_vaccination_target" db:"public_vaccination_target"`
	PublicFirstVaccinationReceived             int64 `json:"public_first_vaccination_received" db:"public_first_vaccination_received"`
	PublicSecondVaccinationReceived            int64 `json:"public_second_vaccination_received" db:"public_second_vaccination_received"`
	CumulativePublicFirstVaccinationReceived   int64 `json:"cumulative_public_first_vaccination_received" db:"cumulative_public_first_vaccination_received"`
	CumulativePublicSecondVaccinationReceived  int64 `json:"cumulative_public_second_vaccination_received" db:"cumulative_public_second_vaccination_received"`

	TeenagerVaccinationTarget                    int64 `json:"teenager_vaccination_target" db:"teenager_vaccination_target"`
	TeenagerFirstVaccinationReceived             int64 `json:"teenager_first_vaccination_received" db:"teenager_first_vaccination_received"`
	TeenagerSecondVaccinationReceived            int64 `json:"teenager_second_vaccination_received" db:"teenager_second_vaccination_received"`
	CumulativeTeenagerFirstVaccinationReceived   int64 `json:"cumulative_teenager_first_vaccination_received" db:"cumulative_teenager_first_vaccination_received"`
	CumulativeTeenagerSecondVaccinationReceived  int64 `json:"cumulative_teenager_second_vaccination_received" db:"cumulative_teenager_second_vaccination_received"`
}

// ProvinceVaccine represents provincial vaccination data
type ProvinceVaccine struct {
	NationalVaccine
	ProvinceID int `json:"province_id" db:"province_id"`
}

// VaccineLocation represents a vaccination center
type VaccineLocation struct {
	ID                        int64   `json:"id" db:"id"`
	RegencyID                 int     `json:"regency_id" db:"regency_id"`
	Name                      string  `json:"name" db:"name"`
	Address                   string  `json:"address" db:"address"`
	OperationalTime           string  `json:"operational_time" db:"operational_time"`
	IsFirstVaccination        bool    `json:"is_first_vaccination" db:"is_first_vaccination"`
	IsSecondVaccination       bool    `json:"is_second_vaccination" db:"is_second_vaccination"`
	DailyVaccinationQuota     *int    `json:"daily_vaccination_quota" db:"daily_vaccination_quota"`
	VaccinationStockRemaining *int    `json:"vaccination_stock_remaining" db:"vaccination_stock_remaining"`
	Notes                     *string `json:"notes" db:"notes"`
}
