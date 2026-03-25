package dto

import (
	"math"

	"github.com/banua-coder/pico-api-go/internal/models"
)

// CoverageData holds vaccination coverage percentages per dose
type CoverageData struct {
	Dose1 float64 `json:"dose_1"`
	Dose2 float64 `json:"dose_2"`
}

// DoseData holds vaccination counts per dose
type DoseData struct {
	Dose1 int64 `json:"dose_1"`
	Dose2 int64 `json:"dose_2"`
}

// GroupData represents vaccination data for a specific group
type GroupData struct {
	Target     int64        `json:"target"`
	Daily      DoseData     `json:"daily"`
	Cumulative DoseData     `json:"cumulative"`
	Coverage   CoverageData `json:"coverage"`
}

// VaccinationTotals represents total national vaccination summary
type VaccinationTotals struct {
	Target     int64        `json:"target"`
	Daily      DoseData     `json:"daily"`
	Cumulative DoseData     `json:"cumulative"`
	Coverage   CoverageData `json:"coverage"`
}

// VaccinationResponse is the DTO for the vaccination summary response
type VaccinationResponse struct {
	Total         VaccinationTotals    `json:"total"`
	HealthWorker  GroupData            `json:"health_worker"`
	Elderly       GroupData            `json:"elderly"`
	PublicOfficer GroupData            `json:"public_officer"`
	Public        GroupData            `json:"public"`
	Teenager      GroupData            `json:"teenager"`
}

// calcCoverage calculates coverage percentage, rounded to 2 decimal places.
// Returns 0.00 if target is 0.
func calcCoverage(cumulative, target int64) float64 {
	if target == 0 {
		return 0.00
	}
	result := (float64(cumulative) / float64(target)) * 100
	return math.Round(result*100) / 100
}

// FromNationalVaccine transforms a NationalVaccine model to VaccinationResponse DTO
func FromNationalVaccine(v models.NationalVaccine) VaccinationResponse {
	return VaccinationResponse{
		Total: VaccinationTotals{
			Target: v.TotalVaccinationTarget,
			Daily: DoseData{
				Dose1: v.FirstVaccinationReceived,
				Dose2: v.SecondVaccinationReceived,
			},
			Cumulative: DoseData{
				Dose1: v.CumulativeFirstVaccinationReceived,
				Dose2: v.CumulativeSecondVaccinationReceived,
			},
			Coverage: CoverageData{
				Dose1: calcCoverage(v.CumulativeFirstVaccinationReceived, v.TotalVaccinationTarget),
				Dose2: calcCoverage(v.CumulativeSecondVaccinationReceived, v.TotalVaccinationTarget),
			},
		},
		HealthWorker: GroupData{
			Target: v.HealthWorkerVaccinationTarget,
			Daily: DoseData{
				Dose1: v.HealthWorkerFirstVaccinationReceived,
				Dose2: v.HealthWorkerSecondVaccinationReceived,
			},
			Cumulative: DoseData{
				Dose1: v.CumulativeHealthWorkerFirstVaccinationReceived,
				Dose2: v.CumulativeHealthWorkerSecondVaccinationReceived,
			},
			Coverage: CoverageData{
				Dose1: calcCoverage(v.CumulativeHealthWorkerFirstVaccinationReceived, v.HealthWorkerVaccinationTarget),
				Dose2: calcCoverage(v.CumulativeHealthWorkerSecondVaccinationReceived, v.HealthWorkerVaccinationTarget),
			},
		},
		Elderly: GroupData{
			Target: v.ElderlyVaccinationTarget,
			Daily: DoseData{
				Dose1: v.ElderlyFirstVaccinationReceived,
				Dose2: v.ElderlySecondVaccinationReceived,
			},
			Cumulative: DoseData{
				Dose1: v.CumulativeElderlyFirstVaccinationReceived,
				Dose2: v.CumulativeElderlySecondVaccinationReceived,
			},
			Coverage: CoverageData{
				Dose1: calcCoverage(v.CumulativeElderlyFirstVaccinationReceived, v.ElderlyVaccinationTarget),
				Dose2: calcCoverage(v.CumulativeElderlySecondVaccinationReceived, v.ElderlyVaccinationTarget),
			},
		},
		PublicOfficer: GroupData{
			Target: v.PublicOfficerVaccinationTarget,
			Daily: DoseData{
				Dose1: v.PublicOfficerFirstVaccinationReceived,
				Dose2: v.PublicOfficerSecondVaccinationReceived,
			},
			Cumulative: DoseData{
				Dose1: v.CumulativePublicOfficerFirstVaccinationReceived,
				Dose2: v.CumulativePublicOfficerSecondVaccinationReceived,
			},
			Coverage: CoverageData{
				Dose1: calcCoverage(v.CumulativePublicOfficerFirstVaccinationReceived, v.PublicOfficerVaccinationTarget),
				Dose2: calcCoverage(v.CumulativePublicOfficerSecondVaccinationReceived, v.PublicOfficerVaccinationTarget),
			},
		},
		Public: GroupData{
			Target: v.PublicVaccinationTarget,
			Daily: DoseData{
				Dose1: v.PublicFirstVaccinationReceived,
				Dose2: v.PublicSecondVaccinationReceived,
			},
			Cumulative: DoseData{
				Dose1: v.CumulativePublicFirstVaccinationReceived,
				Dose2: v.CumulativePublicSecondVaccinationReceived,
			},
			Coverage: CoverageData{
				Dose1: calcCoverage(v.CumulativePublicFirstVaccinationReceived, v.PublicVaccinationTarget),
				Dose2: calcCoverage(v.CumulativePublicSecondVaccinationReceived, v.PublicVaccinationTarget),
			},
		},
		Teenager: GroupData{
			Target: v.TeenagerVaccinationTarget,
			Daily: DoseData{
				Dose1: v.TeenagerFirstVaccinationReceived,
				Dose2: v.TeenagerSecondVaccinationReceived,
			},
			Cumulative: DoseData{
				Dose1: v.CumulativeTeenagerFirstVaccinationReceived,
				Dose2: v.CumulativeTeenagerSecondVaccinationReceived,
			},
			Coverage: CoverageData{
				Dose1: calcCoverage(v.CumulativeTeenagerFirstVaccinationReceived, v.TeenagerVaccinationTarget),
				Dose2: calcCoverage(v.CumulativeTeenagerSecondVaccinationReceived, v.TeenagerVaccinationTarget),
			},
		},
	}
}
