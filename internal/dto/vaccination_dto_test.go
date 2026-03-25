package dto

import (
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
)

func TestTransformNationalVaccine(t *testing.T) {
	v := models.NationalVaccine{
		ID:   1,
		Day:  93,
		Date: time.Date(2021, 5, 4, 0, 0, 0, 0, time.UTC),
		TotalVaccinationTarget:                          2240548,
		FirstVaccinationReceived:                        473,
		SecondVaccinationReceived:                       1535,
		CumulativeFirstVaccinationReceived:              81885,
		CumulativeSecondVaccinationReceived:             61974,
		HealthWorkerVaccinationTarget:                   24698,
		HealthWorkerFirstVaccinationReceived:            1,
		HealthWorkerSecondVaccinationReceived:           36,
		CumulativeHealthWorkerFirstVaccinationReceived:  24788,
		CumulativeHealthWorkerSecondVaccinationReceived: 23075,
		ElderlyVaccinationTarget:                        25292,
		ElderlyFirstVaccinationReceived:                 37,
		ElderlySecondVaccinationReceived:                295,
		CumulativeElderlyFirstVaccinationReceived:       6084,
		CumulativeElderlySecondVaccinationReceived:      3790,
		PublicOfficerVaccinationTarget:                   65471,
		PublicOfficerFirstVaccinationReceived:            435,
		PublicOfficerSecondVaccinationReceived:           1204,
		CumulativePublicOfficerFirstVaccinationReceived:  51013,
		CumulativePublicOfficerSecondVaccinationReceived: 35109,
		PublicVaccinationTarget:                          1374811,
		TeenagerVaccinationTarget:                        314609,
	}

	result := TransformNationalVaccine(v)

	// Basic fields
	if result.ID != 1 { t.Errorf("ID = %d, want 1", result.ID) }
	if result.Day != 93 { t.Errorf("Day = %d, want 93", result.Day) }
	if result.Target != 2240548 { t.Errorf("Target = %d, want 2240548", result.Target) }

	// Total
	if result.Total.Daily.Dose1 != 473 { t.Errorf("Total.Daily.Dose1 = %d, want 473", result.Total.Daily.Dose1) }
	if result.Total.Daily.Dose2 != 1535 { t.Errorf("Total.Daily.Dose2 = %d, want 1535", result.Total.Daily.Dose2) }
	if result.Total.Cumulative.Dose1 != 81885 { t.Errorf("Total.Cumulative.Dose1 = %d, want 81885", result.Total.Cumulative.Dose1) }
	if result.Total.Cumulative.Dose2 != 61974 { t.Errorf("Total.Cumulative.Dose2 = %d, want 61974", result.Total.Cumulative.Dose2) }

	// Groups - all 5 should exist
	expectedGroups := []string{"health_worker", "elderly", "public_officer", "public", "teenager"}
	for _, g := range expectedGroups {
		if _, ok := result.Groups[g]; !ok {
			t.Errorf("Missing group: %s", g)
		}
	}

	// Health worker details
	hw := result.Groups["health_worker"]
	if hw.Target != 24698 { t.Errorf("hw.Target = %d, want 24698", hw.Target) }
	if hw.Daily.Dose1 != 1 { t.Errorf("hw.Daily.Dose1 = %d, want 1", hw.Daily.Dose1) }
	if hw.Cumulative.Dose1 != 24788 { t.Errorf("hw.Cumulative.Dose1 = %d, want 24788", hw.Cumulative.Dose1) }

	// Elderly
	el := result.Groups["elderly"]
	if el.Target != 25292 { t.Errorf("el.Target = %d, want 25292", el.Target) }
	if el.Daily.Dose2 != 295 { t.Errorf("el.Daily.Dose2 = %d, want 295", el.Daily.Dose2) }
}

func TestTransformNationalVaccine_ZeroValues(t *testing.T) {
	v := models.NationalVaccine{}
	result := TransformNationalVaccine(v)

	if result.Target != 0 { t.Errorf("Target = %d, want 0", result.Target) }
	if result.Total.Daily.Dose1 != 0 { t.Errorf("Total.Daily.Dose1 = %d, want 0", result.Total.Daily.Dose1) }
	if len(result.Groups) != 5 { t.Errorf("Groups count = %d, want 5", len(result.Groups)) }
}

func TestTransformProvinceVaccine(t *testing.T) {
	v := models.ProvinceVaccine{
		NationalVaccine: models.NationalVaccine{
			ID: 1, Day: 93, TotalVaccinationTarget: 100000,
			FirstVaccinationReceived: 500, CumulativeFirstVaccinationReceived: 5000,
		},
		ProvinceID: 72,
	}

	result := TransformProvinceVaccine(v)
	if result.ProvinceID != 72 { t.Errorf("ProvinceID = %d, want 72", result.ProvinceID) }
	if result.Target != 100000 { t.Errorf("Target = %d, want 100000", result.Target) }
	if result.Total.Daily.Dose1 != 500 { t.Errorf("Total.Daily.Dose1 = %d, want 500", result.Total.Daily.Dose1) }
}
