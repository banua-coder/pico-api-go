package dto

import (
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
)

func sampleNationalVaccine() models.NationalVaccine {
	return models.NationalVaccine{
		ID:   1,
		Day:  100,
		Date: time.Now(),

		TotalVaccinationTarget:              10000,
		FirstVaccinationReceived:            50,
		SecondVaccinationReceived:           40,
		CumulativeFirstVaccinationReceived:  365,
		CumulativeSecondVaccinationReceived: 277,

		HealthWorkerVaccinationTarget:                   1000,
		HealthWorkerFirstVaccinationReceived:            10,
		HealthWorkerSecondVaccinationReceived:           8,
		CumulativeHealthWorkerFirstVaccinationReceived:  500,
		CumulativeHealthWorkerSecondVaccinationReceived: 400,

		ElderlyVaccinationTarget:                   2000,
		ElderlyFirstVaccinationReceived:            20,
		ElderlySecondVaccinationReceived:           15,
		CumulativeElderlyFirstVaccinationReceived:  1000,
		CumulativeElderlySecondVaccinationReceived: 800,

		PublicOfficerVaccinationTarget:                   500,
		PublicOfficerFirstVaccinationReceived:            5,
		PublicOfficerSecondVaccinationReceived:           4,
		CumulativePublicOfficerFirstVaccinationReceived:  250,
		CumulativePublicOfficerSecondVaccinationReceived: 200,

		PublicVaccinationTarget:                   5000,
		PublicFirstVaccinationReceived:            30,
		PublicSecondVaccinationReceived:           25,
		CumulativePublicFirstVaccinationReceived:  2500,
		CumulativePublicSecondVaccinationReceived: 2000,

		TeenagerVaccinationTarget:                   1500,
		TeenagerFirstVaccinationReceived:            15,
		TeenagerSecondVaccinationReceived:           10,
		CumulativeTeenagerFirstVaccinationReceived:  750,
		CumulativeTeenagerSecondVaccinationReceived: 600,
	}
}

func TestFromNationalVaccine_TotalCoverage(t *testing.T) {
	v := sampleNationalVaccine()
	dto := FromNationalVaccine(v)

	// coverage = (365 / 10000) * 100 = 3.65
	assert.Equal(t, 3.65, dto.Total.Coverage.Dose1)
	// coverage = (277 / 10000) * 100 = 2.77
	assert.Equal(t, 2.77, dto.Total.Coverage.Dose2)
}

func TestFromNationalVaccine_GroupCoverage(t *testing.T) {
	v := sampleNationalVaccine()
	dto := FromNationalVaccine(v)

	// HealthWorker: (500/1000)*100 = 50.00
	assert.Equal(t, 50.0, dto.HealthWorker.Coverage.Dose1)
	assert.Equal(t, 40.0, dto.HealthWorker.Coverage.Dose2)

	// Elderly: (1000/2000)*100 = 50.00
	assert.Equal(t, 50.0, dto.Elderly.Coverage.Dose1)
	assert.Equal(t, 40.0, dto.Elderly.Coverage.Dose2)

	// PublicOfficer: (250/500)*100 = 50.00
	assert.Equal(t, 50.0, dto.PublicOfficer.Coverage.Dose1)
	assert.Equal(t, 40.0, dto.PublicOfficer.Coverage.Dose2)

	// Public: (2500/5000)*100 = 50.00
	assert.Equal(t, 50.0, dto.Public.Coverage.Dose1)
	assert.Equal(t, 40.0, dto.Public.Coverage.Dose2)

	// Teenager: (750/1500)*100 = 50.00
	assert.Equal(t, 50.0, dto.Teenager.Coverage.Dose1)
	assert.Equal(t, 40.0, dto.Teenager.Coverage.Dose2)
}

func TestFromNationalVaccine_ZeroTarget(t *testing.T) {
	v := models.NationalVaccine{
		TotalVaccinationTarget:             0,
		CumulativeFirstVaccinationReceived: 100,
	}
	dto := FromNationalVaccine(v)

	// target is 0, should return 0.00 (no division by zero)
	assert.Equal(t, 0.00, dto.Total.Coverage.Dose1)
	assert.Equal(t, 0.00, dto.Total.Coverage.Dose2)
}

func TestFromNationalVaccine_TotalFields(t *testing.T) {
	v := sampleNationalVaccine()
	dto := FromNationalVaccine(v)

	assert.Equal(t, int64(10000), dto.Total.Target)
	assert.Equal(t, int64(50), dto.Total.Daily.Dose1)
	assert.Equal(t, int64(40), dto.Total.Daily.Dose2)
	assert.Equal(t, int64(365), dto.Total.Cumulative.Dose1)
	assert.Equal(t, int64(277), dto.Total.Cumulative.Dose2)
}

func TestCalcCoverage_Rounding(t *testing.T) {
	// 1/3 * 100 = 33.33333... → should round to 33.33
	result := calcCoverage(1, 3)
	assert.Equal(t, 33.33, result)
}

func TestCalcCoverage_ZeroTarget(t *testing.T) {
	assert.Equal(t, 0.00, calcCoverage(100, 0))
}
