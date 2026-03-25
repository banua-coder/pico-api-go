package dto

import (
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
)

// DoseData holds dose_1 and dose_2 counts.
type DoseData struct {
	Dose1 int64 `json:"dose_1"`
	Dose2 int64 `json:"dose_2"`
}

// GroupData holds vaccination data for a specific group.
type GroupData struct {
	Target     int64    `json:"target"`
	Daily      DoseData `json:"daily"`
	Cumulative DoseData `json:"cumulative"`
}

// VaccinationTotals holds total daily and cumulative dose data.
type VaccinationTotals struct {
	Daily      DoseData `json:"daily"`
	Cumulative DoseData `json:"cumulative"`
}

// VaccinationResponse is the API response for national vaccination data.
type VaccinationResponse struct {
	ID     int64             `json:"id"`
	Day    int64             `json:"day"`
	Date   time.Time         `json:"date"`
	Target int64             `json:"target"`
	Total  VaccinationTotals `json:"total"`
	Groups map[string]GroupData `json:"groups"`
}

// ProvinceVaccinationResponse is the API response for provincial vaccination data.
type ProvinceVaccinationResponse struct {
	ID         int64             `json:"id"`
	Day        int64             `json:"day"`
	Date       time.Time         `json:"date"`
	ProvinceID int               `json:"province_id"`
	Target     int64             `json:"target"`
	Total      VaccinationTotals `json:"total"`
	Groups     map[string]GroupData `json:"groups"`
}

// TransformNationalVaccine converts a NationalVaccine model to VaccinationResponse DTO.
func TransformNationalVaccine(v models.NationalVaccine) VaccinationResponse {
	return VaccinationResponse{
		ID:     v.ID,
		Day:    v.Day,
		Date:   v.Date,
		Target: v.TotalVaccinationTarget,
		Total: VaccinationTotals{
			Daily: DoseData{
				Dose1: v.FirstVaccinationReceived,
				Dose2: v.SecondVaccinationReceived,
			},
			Cumulative: DoseData{
				Dose1: v.CumulativeFirstVaccinationReceived,
				Dose2: v.CumulativeSecondVaccinationReceived,
			},
		},
		Groups: map[string]GroupData{
			"health_worker": {
				Target: v.HealthWorkerVaccinationTarget,
				Daily: DoseData{
					Dose1: v.HealthWorkerFirstVaccinationReceived,
					Dose2: v.HealthWorkerSecondVaccinationReceived,
				},
				Cumulative: DoseData{
					Dose1: v.CumulativeHealthWorkerFirstVaccinationReceived,
					Dose2: v.CumulativeHealthWorkerSecondVaccinationReceived,
				},
			},
			"elderly": {
				Target: v.ElderlyVaccinationTarget,
				Daily: DoseData{
					Dose1: v.ElderlyFirstVaccinationReceived,
					Dose2: v.ElderlySecondVaccinationReceived,
				},
				Cumulative: DoseData{
					Dose1: v.CumulativeElderlyFirstVaccinationReceived,
					Dose2: v.CumulativeElderlySecondVaccinationReceived,
				},
			},
			"public_officer": {
				Target: v.PublicOfficerVaccinationTarget,
				Daily: DoseData{
					Dose1: v.PublicOfficerFirstVaccinationReceived,
					Dose2: v.PublicOfficerSecondVaccinationReceived,
				},
				Cumulative: DoseData{
					Dose1: v.CumulativePublicOfficerFirstVaccinationReceived,
					Dose2: v.CumulativePublicOfficerSecondVaccinationReceived,
				},
			},
			"public": {
				Target: v.PublicVaccinationTarget,
				Daily: DoseData{
					Dose1: v.PublicFirstVaccinationReceived,
					Dose2: v.PublicSecondVaccinationReceived,
				},
				Cumulative: DoseData{
					Dose1: v.CumulativePublicFirstVaccinationReceived,
					Dose2: v.CumulativePublicSecondVaccinationReceived,
				},
			},
			"teenager": {
				Target: v.TeenagerVaccinationTarget,
				Daily: DoseData{
					Dose1: v.TeenagerFirstVaccinationReceived,
					Dose2: v.TeenagerSecondVaccinationReceived,
				},
				Cumulative: DoseData{
					Dose1: v.CumulativeTeenagerFirstVaccinationReceived,
					Dose2: v.CumulativeTeenagerSecondVaccinationReceived,
				},
			},
		},
	}
}

// TransformProvinceVaccine converts a ProvinceVaccine model to ProvinceVaccinationResponse DTO.
func TransformProvinceVaccine(v models.ProvinceVaccine) ProvinceVaccinationResponse {
	national := TransformNationalVaccine(v.NationalVaccine)
	return ProvinceVaccinationResponse{
		ID:         national.ID,
		Day:        national.Day,
		Date:       national.Date,
		ProvinceID: v.ProvinceID,
		Target:     national.Target,
		Total:      national.Total,
		Groups:     national.Groups,
	}
}
