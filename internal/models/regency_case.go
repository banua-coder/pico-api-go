package models

import "time"

// RegencyCase represents COVID-19 case data for a regency
type RegencyCase struct {
	ID                                       int64     `json:"id" db:"id"`
	Day                                      int64     `json:"day" db:"day"`
	RegencyID                                int       `json:"regency_id" db:"regency_id"`
	Positive                                 int64     `json:"positive" db:"positive"`
	Recovered                                int64     `json:"recovered" db:"recovered"`
	Deceased                                 int64     `json:"deceased" db:"deceased"`
	PersonUnderObservation                   *float64  `json:"person_under_observation" db:"person_under_observation"`
	FinishedPersonUnderObservation           *int64    `json:"finished_person_under_observation" db:"finished_person_under_observation"`
	PersonUnderSupervision                   *float64  `json:"person_under_supervision" db:"person_under_supervision"`
	FinishedPersonUnderSupervision           *int64    `json:"finished_person_under_supervision" db:"finished_person_under_supervision"`
	CumulativePositive                       int64     `json:"cumulative_positive" db:"cumulative_positive"`
	CumulativeRecovered                      int64     `json:"cumulative_recovered" db:"cumulative_recovered"`
	CumulativeDeceased                       int64     `json:"cumulative_deceased" db:"cumulative_deceased"`
	CumulativePersonUnderObservation         *float64  `json:"cumulative_person_under_observation" db:"cumulative_person_under_observation"`
	CumulativeFinishedPersonUnderObservation *int64    `json:"cumulative_finished_person_under_observation" db:"cumulative_finished_person_under_observation"`
	CumulativePersonUnderSupervision         *float64  `json:"cumulative_person_under_supervision" db:"cumulative_person_under_supervision"`
	CumulativeFinishedPersonUnderSupervision *int64    `json:"cumulative_finished_person_under_supervision" db:"cumulative_finished_person_under_supervision"`
	Rt                                       *float64  `json:"rt" db:"rt"`
	RtUpper                                  *float64  `json:"rt_upper" db:"rt_upper"`
	RtLower                                  *float64  `json:"rt_lower" db:"rt_lower"`
	Regency                                  *Regency  `json:"regency,omitempty"`
	Date                                     *time.Time `json:"date,omitempty" db:"date"`
}
