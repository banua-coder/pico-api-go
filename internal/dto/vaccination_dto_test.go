package dto

import (
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
)

func sampleVaccine() models.NationalVaccine {
	return models.NationalVaccine{
		ID: 1, Day: 93, Date: time.Date(2021, 5, 4, 0, 0, 0, 0, time.UTC),
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
	}
}

func TestTransformNationalVaccine(t *testing.T) {
	r := TransformNationalVaccine(sampleVaccine())

	if r.Target != 2240548 {
		t.Errorf("Target = %d, want 2240548", r.Target)
	}
	if r.Total.Daily.Dose1 != 473 {
		t.Errorf("Total.Daily.Dose1 = %d, want 473", r.Total.Daily.Dose1)
	}
	if r.Total.Cumulative.Dose1 != 81885 {
		t.Errorf("Total.Cumulative.Dose1 = %d, want 81885", r.Total.Cumulative.Dose1)
	}
	if r.Total.Coverage.Dose1 != 3.65 {
		t.Errorf("Total.Coverage.Dose1 = %f, want 3.65", r.Total.Coverage.Dose1)
	}
	if r.Total.Coverage.Dose2 != 2.77 {
		t.Errorf("Total.Coverage.Dose2 = %f, want 2.77", r.Total.Coverage.Dose2)
	}

	// All 5 groups
	for _, g := range []string{"health_worker", "elderly", "public_officer", "public", "teenager"} {
		if _, ok := r.Groups[g]; !ok {
			t.Errorf("Missing group: %s", g)
		}
	}

	hw := r.Groups["health_worker"]
	if hw.Coverage.Dose1 != 100.36 {
		t.Errorf("hw.Coverage.Dose1 = %f, want 100.36", hw.Coverage.Dose1)
	}
	if hw.Coverage.Dose2 != 93.43 {
		t.Errorf("hw.Coverage.Dose2 = %f, want 93.43", hw.Coverage.Dose2)
	}
}

func TestTransformNationalVaccine_ZeroValues(t *testing.T) {
	r := TransformNationalVaccine(models.NationalVaccine{})
	if r.Total.Coverage.Dose1 != 0.0 {
		t.Errorf("Coverage.Dose1 = %f, want 0.0", r.Total.Coverage.Dose1)
	}
	if len(r.Groups) != 5 {
		t.Errorf("Groups count = %d, want 5", len(r.Groups))
	}
}

func TestTransformProvinceVaccine(t *testing.T) {
	v := models.ProvinceVaccine{NationalVaccine: sampleVaccine(), ProvinceID: 72}
	r := TransformProvinceVaccine(v)
	if r.ProvinceID != 72 {
		t.Errorf("ProvinceID = %d, want 72", r.ProvinceID)
	}
	if r.Total.Coverage.Dose1 != 3.65 {
		t.Errorf("Coverage.Dose1 = %f, want 3.65", r.Total.Coverage.Dose1)
	}
}

func TestCalcCoverage_ZeroTarget(t *testing.T) {
	if r := calcCoverage(5000, 0); r != 0.0 {
		t.Errorf("calcCoverage(5000, 0) = %f, want 0.0", r)
	}
}

func TestCalcCoverage_Rounding(t *testing.T) {
	if r := calcCoverage(1, 3); r != 33.33 {
		t.Errorf("calcCoverage(1, 3) = %f, want 33.33", r)
	}
}
