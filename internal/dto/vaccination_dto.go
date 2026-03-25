package dto

import (
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
)

// VaccinationDose represents daily and cumulative dose counts
type VaccinationDose struct {
	Daily      int64 `json:"daily"`
	Cumulative int64 `json:"cumulative"`
}

// VaccinationGroup represents vaccination data for a specific group
type VaccinationGroup struct {
	Target int64           `json:"target"`
	Dose1  VaccinationDose `json:"dose_1"`
	Dose2  VaccinationDose `json:"dose_2"`
}

// VaccinationResponse is the structured API response for vaccination data
// @Description Nested vaccination response with groups
type VaccinationResponse struct {
	ID     int64     `json:"id"`
	Day    int64     `json:"day"`
	Date   time.Time `json:"date"`
	Target int64     `json:"total_vaccination_target"`
	Dose1  VaccinationDose `json:"dose_1"`
	Dose2  VaccinationDose `json:"dose_2"`
	Groups map[string]VaccinationGroup `json:"groups"`
}

// ProvinceVaccinationResponse extends VaccinationResponse with province info
type ProvinceVaccinationResponse struct {
	VaccinationResponse
	ProvinceID int `json:"province_id"`
}

// TransformNationalVaccine transforms a NationalVaccine model to VaccinationResponse DTO
func TransformNationalVaccine(v models.NationalVaccine) VaccinationResponse {
	return VaccinationResponse{
		ID:     v.ID,
		Day:    v.Day,
		Date:   v.Date,
		Target: v.TotalVaccinationTarget,
		Dose1: VaccinationDose{
			Daily:      v.FirstVaccinationReceived,
			Cumulative: v.CumulativeFirstVaccinationReceived,
		},
		Dose2: VaccinationDose{
			Daily:      v.SecondVaccinationReceived,
			Cumulative: v.CumulativeSecondVaccinationReceived,
		},
		Groups: map[string]VaccinationGroup{
			"health_worker": {
				Target: v.HealthWorkerVaccinationTarget,
				Dose1: VaccinationDose{
					Daily:      v.HealthWorkerFirstVaccinationReceived,
					Cumulative: v.CumulativeHealthWorkerFirstVaccinationReceived,
				},
				Dose2: VaccinationDose{
					Daily:      v.HealthWorkerSecondVaccinationReceived,
					Cumulative: v.CumulativeHealthWorkerSecondVaccinationReceived,
				},
			},
			"elderly": {
				Target: v.ElderlyVaccinationTarget,
				Dose1: VaccinationDose{
					Daily:      v.ElderlyFirstVaccinationReceived,
					Cumulative: v.CumulativeElderlyFirstVaccinationReceived,
				},
				Dose2: VaccinationDose{
					Daily:      v.ElderlySecondVaccinationReceived,
					Cumulative: v.CumulativeElderlySecondVaccinationReceived,
				},
			},
			"public_officer": {
				Target: v.PublicOfficerVaccinationTarget,
				Dose1: VaccinationDose{
					Daily:      v.PublicOfficerFirstVaccinationReceived,
					Cumulative: v.CumulativePublicOfficerFirstVaccinationReceived,
				},
				Dose2: VaccinationDose{
					Daily:      v.PublicOfficerSecondVaccinationReceived,
					Cumulative: v.CumulativePublicOfficerSecondVaccinationReceived,
				},
			},
			"public": {
				Target: v.PublicVaccinationTarget,
				Dose1: VaccinationDose{
					Daily:      v.PublicFirstVaccinationReceived,
					Cumulative: v.CumulativePublicFirstVaccinationReceived,
				},
				Dose2: VaccinationDose{
					Daily:      v.PublicSecondVaccinationReceived,
					Cumulative: v.CumulativePublicSecondVaccinationReceived,
				},
			},
			"teenager": {
				Target: v.TeenagerVaccinationTarget,
				Dose1: VaccinationDose{
					Daily:      v.TeenagerFirstVaccinationReceived,
					Cumulative: v.CumulativeTeenagerFirstVaccinationReceived,
				},
				Dose2: VaccinationDose{
					Daily:      v.TeenagerSecondVaccinationReceived,
					Cumulative: v.CumulativeTeenagerSecondVaccinationReceived,
				},
			},
		},
	}
}

// TransformProvinceVaccine transforms a ProvinceVaccine model to ProvinceVaccinationResponse DTO
func TransformProvinceVaccine(v models.ProvinceVaccine) ProvinceVaccinationResponse {
	return ProvinceVaccinationResponse{
		VaccinationResponse: TransformNationalVaccine(v.NationalVaccine),
		ProvinceID:          v.ProvinceID,
	}
}
