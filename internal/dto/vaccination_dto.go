package dto

import (
	"math"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
)

// CoverageData represents vaccination coverage as percentage.
type CoverageData struct {
	Dose1 float64 `json:"dose_1"`
	Dose2 float64 `json:"dose_2"`
}

// DoseData holds dose_1 and dose_2 counts.
type DoseData struct {
	Dose1 int64 `json:"dose_1"`
	Dose2 int64 `json:"dose_2"`
}

// GroupData holds vaccination data for a specific group.
type GroupData struct {
	Target     int64        `json:"target"`
	Daily      DoseData     `json:"daily"`
	Cumulative DoseData     `json:"cumulative"`
	Coverage   CoverageData `json:"coverage"`
}

// VaccinationTotals holds total daily and cumulative dose data.
type VaccinationTotals struct {
	Daily      DoseData     `json:"daily"`
	Cumulative DoseData     `json:"cumulative"`
	Coverage   CoverageData `json:"coverage"`
}

// VaccinationResponse is the API response for national vaccination data.
type VaccinationResponse struct {
	ID     int64                `json:"id"`
	Day    int64                `json:"day"`
	Date   time.Time            `json:"date"`
	Target int64                `json:"target"`
	Total  VaccinationTotals    `json:"total"`
	Groups map[string]GroupData `json:"groups"`
}

// ProvinceVaccinationResponse is the API response for provincial vaccination data.
type ProvinceVaccinationResponse struct {
	ID         int64                `json:"id"`
	Day        int64                `json:"day"`
	Date       time.Time            `json:"date"`
	ProvinceID int                  `json:"province_id"`
	Target     int64                `json:"target"`
	Total      VaccinationTotals    `json:"total"`
	Groups     map[string]GroupData `json:"groups"`
}

// calcCoverage returns (cumulative / target) * 100, rounded to 2 decimals.
// Returns 0.00 if target is 0.
func calcCoverage(cumulative int64, target int64) float64 {
	if target == 0 {
		return 0.0
	}
	return math.Round(float64(cumulative)/float64(target)*10000) / 100
}

func buildGroup(target int64, daily, cumulative DoseData) GroupData {
	return GroupData{
		Target: target, Daily: daily, Cumulative: cumulative,
		Coverage: CoverageData{
			Dose1: calcCoverage(cumulative.Dose1, target),
			Dose2: calcCoverage(cumulative.Dose2, target),
		},
	}
}

// TransformNationalVaccine converts a NationalVaccine model to VaccinationResponse DTO.
func TransformNationalVaccine(v models.NationalVaccine) VaccinationResponse {
	totalDaily := DoseData{Dose1: v.FirstVaccinationReceived, Dose2: v.SecondVaccinationReceived}
	totalCum := DoseData{Dose1: v.CumulativeFirstVaccinationReceived, Dose2: v.CumulativeSecondVaccinationReceived}

	return VaccinationResponse{
		ID: v.ID, Day: v.Day, Date: v.Date, Target: v.TotalVaccinationTarget,
		Total: VaccinationTotals{
			Daily: totalDaily, Cumulative: totalCum,
			Coverage: CoverageData{
				Dose1: calcCoverage(totalCum.Dose1, v.TotalVaccinationTarget),
				Dose2: calcCoverage(totalCum.Dose2, v.TotalVaccinationTarget),
			},
		},
		Groups: map[string]GroupData{
			"health_worker": buildGroup(v.HealthWorkerVaccinationTarget,
				DoseData{v.HealthWorkerFirstVaccinationReceived, v.HealthWorkerSecondVaccinationReceived},
				DoseData{v.CumulativeHealthWorkerFirstVaccinationReceived, v.CumulativeHealthWorkerSecondVaccinationReceived}),
			"elderly": buildGroup(v.ElderlyVaccinationTarget,
				DoseData{v.ElderlyFirstVaccinationReceived, v.ElderlySecondVaccinationReceived},
				DoseData{v.CumulativeElderlyFirstVaccinationReceived, v.CumulativeElderlySecondVaccinationReceived}),
			"public_officer": buildGroup(v.PublicOfficerVaccinationTarget,
				DoseData{v.PublicOfficerFirstVaccinationReceived, v.PublicOfficerSecondVaccinationReceived},
				DoseData{v.CumulativePublicOfficerFirstVaccinationReceived, v.CumulativePublicOfficerSecondVaccinationReceived}),
			"public": buildGroup(v.PublicVaccinationTarget,
				DoseData{v.PublicFirstVaccinationReceived, v.PublicSecondVaccinationReceived},
				DoseData{v.CumulativePublicFirstVaccinationReceived, v.CumulativePublicSecondVaccinationReceived}),
			"teenager": buildGroup(v.TeenagerVaccinationTarget,
				DoseData{v.TeenagerFirstVaccinationReceived, v.TeenagerSecondVaccinationReceived},
				DoseData{v.CumulativeTeenagerFirstVaccinationReceived, v.CumulativeTeenagerSecondVaccinationReceived}),
		},
	}
}

// TransformProvinceVaccine converts a ProvinceVaccine model to ProvinceVaccinationResponse DTO.
func TransformProvinceVaccine(v models.ProvinceVaccine) ProvinceVaccinationResponse {
	national := TransformNationalVaccine(v.NationalVaccine)
	return ProvinceVaccinationResponse{
		ID: national.ID, Day: national.Day, Date: national.Date,
		ProvinceID: v.ProvinceID, Target: national.Target,
		Total: national.Total, Groups: national.Groups,
	}
}
