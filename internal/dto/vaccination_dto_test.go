package dto

import (
	"math"
	"testing"
	"time"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
)

func sampleNationalVaccine() models.NationalVaccine {
	return models.NationalVaccine{
		ID:   1,
		Day:  100,
		Date: time.Date(2021, 8, 17, 0, 0, 0, 0, time.UTC),

		TotalVaccinationTarget: 208_000_000,

		FirstVaccinationReceived:            500_000,
		SecondVaccinationReceived:           300_000,
		CumulativeFirstVaccinationReceived:  50_000_000,
		CumulativeSecondVaccinationReceived: 30_000_000,

		HealthWorkerVaccinationTarget:                   1_468_764,
		HealthWorkerFirstVaccinationReceived:            10_000,
		HealthWorkerSecondVaccinationReceived:           8_000,
		CumulativeHealthWorkerFirstVaccinationReceived:  1_400_000,
		CumulativeHealthWorkerSecondVaccinationReceived: 1_200_000,

		ElderlyVaccinationTarget:                   21_553_118,
		ElderlyFirstVaccinationReceived:            50_000,
		ElderlySecondVaccinationReceived:           40_000,
		CumulativeElderlyFirstVaccinationReceived:  15_000_000,
		CumulativeElderlySecondVaccinationReceived: 10_000_000,

		PublicOfficerVaccinationTarget:                   17_327_167,
		PublicOfficerFirstVaccinationReceived:            30_000,
		PublicOfficerSecondVaccinationReceived:           20_000,
		CumulativePublicOfficerFirstVaccinationReceived:  12_000_000,
		CumulativePublicOfficerSecondVaccinationReceived: 9_000_000,

		PublicVaccinationTarget:                   141_211_180,
		PublicFirstVaccinationReceived:            400_000,
		PublicSecondVaccinationReceived:           230_000,
		CumulativePublicFirstVaccinationReceived:  20_000_000,
		CumulativePublicSecondVaccinationReceived: 9_800_000,

		TeenagerVaccinationTarget:                   26_705_490,
		TeenagerFirstVaccinationReceived:            10_000,
		TeenagerSecondVaccinationReceived:           2_000,
		CumulativeTeenagerFirstVaccinationReceived:  1_600_000,
		CumulativeTeenagerSecondVaccinationReceived: 0,
	}
}

func TestTransformNationalVaccine(t *testing.T) {
	testDate := time.Date(2021, 8, 17, 0, 0, 0, 0, time.UTC)
	maxVal := int64(math.MaxInt64)

	tests := []struct {
		name  string
		input models.NationalVaccine
	}{
		{
			name:  "zero values",
			input: models.NationalVaccine{},
		},
		{
			name:  "sample real-world data",
			input: sampleNationalVaccine(),
		},
		{
			name: "max int64 values",
			input: models.NationalVaccine{
				ID:                                              maxVal,
				Day:                                             maxVal,
				Date:                                            testDate,
				TotalVaccinationTarget:                          maxVal,
				FirstVaccinationReceived:                        maxVal,
				SecondVaccinationReceived:                       maxVal,
				CumulativeFirstVaccinationReceived:              maxVal,
				CumulativeSecondVaccinationReceived:             maxVal,
				HealthWorkerVaccinationTarget:                   maxVal,
				HealthWorkerFirstVaccinationReceived:            maxVal,
				HealthWorkerSecondVaccinationReceived:           maxVal,
				CumulativeHealthWorkerFirstVaccinationReceived:  maxVal,
				CumulativeHealthWorkerSecondVaccinationReceived: maxVal,
				ElderlyVaccinationTarget:                        maxVal,
				ElderlyFirstVaccinationReceived:                 maxVal,
				ElderlySecondVaccinationReceived:                maxVal,
				CumulativeElderlyFirstVaccinationReceived:       maxVal,
				CumulativeElderlySecondVaccinationReceived:      maxVal,
				PublicOfficerVaccinationTarget:                  maxVal,
				PublicOfficerFirstVaccinationReceived:           maxVal,
				PublicOfficerSecondVaccinationReceived:          maxVal,
				CumulativePublicOfficerFirstVaccinationReceived:  maxVal,
				CumulativePublicOfficerSecondVaccinationReceived: maxVal,
				PublicVaccinationTarget:                         maxVal,
				PublicFirstVaccinationReceived:                  maxVal,
				PublicSecondVaccinationReceived:                 maxVal,
				CumulativePublicFirstVaccinationReceived:        maxVal,
				CumulativePublicSecondVaccinationReceived:       maxVal,
				TeenagerVaccinationTarget:                       maxVal,
				TeenagerFirstVaccinationReceived:                maxVal,
				TeenagerSecondVaccinationReceived:               maxVal,
				CumulativeTeenagerFirstVaccinationReceived:      maxVal,
				CumulativeTeenagerSecondVaccinationReceived:     maxVal,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TransformNationalVaccine(tt.input)

			// Basic fields
			assert.Equal(t, tt.input.ID, result.ID)
			assert.Equal(t, tt.input.Day, result.Day)
			assert.Equal(t, tt.input.Date, result.Date)
			assert.Equal(t, tt.input.TotalVaccinationTarget, result.Target)

			// Overall dose_1 and dose_2
			assert.Equal(t, tt.input.FirstVaccinationReceived, result.Dose1.Daily, "dose_1 daily should match FirstVaccinationReceived")
			assert.Equal(t, tt.input.CumulativeFirstVaccinationReceived, result.Dose1.Cumulative, "dose_1 cumulative should match CumulativeFirstVaccinationReceived")
			assert.Equal(t, tt.input.SecondVaccinationReceived, result.Dose2.Daily, "dose_2 daily should match SecondVaccinationReceived")
			assert.Equal(t, tt.input.CumulativeSecondVaccinationReceived, result.Dose2.Cumulative, "dose_2 cumulative should match CumulativeSecondVaccinationReceived")

			// Groups presence
			requiredGroups := []string{"health_worker", "elderly", "public_officer", "public", "teenager"}
			assert.Len(t, result.Groups, 5, "groups map should have exactly 5 keys")
			for _, key := range requiredGroups {
				_, ok := result.Groups[key]
				assert.True(t, ok, "groups should contain key: %s", key)
			}

			// health_worker group
			hw := result.Groups["health_worker"]
			assert.Equal(t, tt.input.HealthWorkerVaccinationTarget, hw.Target)
			assert.Equal(t, tt.input.HealthWorkerFirstVaccinationReceived, hw.Dose1.Daily)
			assert.Equal(t, tt.input.CumulativeHealthWorkerFirstVaccinationReceived, hw.Dose1.Cumulative)
			assert.Equal(t, tt.input.HealthWorkerSecondVaccinationReceived, hw.Dose2.Daily)
			assert.Equal(t, tt.input.CumulativeHealthWorkerSecondVaccinationReceived, hw.Dose2.Cumulative)

			// elderly group
			el := result.Groups["elderly"]
			assert.Equal(t, tt.input.ElderlyVaccinationTarget, el.Target)
			assert.Equal(t, tt.input.ElderlyFirstVaccinationReceived, el.Dose1.Daily)
			assert.Equal(t, tt.input.CumulativeElderlyFirstVaccinationReceived, el.Dose1.Cumulative)
			assert.Equal(t, tt.input.ElderlySecondVaccinationReceived, el.Dose2.Daily)
			assert.Equal(t, tt.input.CumulativeElderlySecondVaccinationReceived, el.Dose2.Cumulative)

			// public_officer group
			po := result.Groups["public_officer"]
			assert.Equal(t, tt.input.PublicOfficerVaccinationTarget, po.Target)
			assert.Equal(t, tt.input.PublicOfficerFirstVaccinationReceived, po.Dose1.Daily)
			assert.Equal(t, tt.input.CumulativePublicOfficerFirstVaccinationReceived, po.Dose1.Cumulative)
			assert.Equal(t, tt.input.PublicOfficerSecondVaccinationReceived, po.Dose2.Daily)
			assert.Equal(t, tt.input.CumulativePublicOfficerSecondVaccinationReceived, po.Dose2.Cumulative)

			// public group
			pub := result.Groups["public"]
			assert.Equal(t, tt.input.PublicVaccinationTarget, pub.Target)
			assert.Equal(t, tt.input.PublicFirstVaccinationReceived, pub.Dose1.Daily)
			assert.Equal(t, tt.input.CumulativePublicFirstVaccinationReceived, pub.Dose1.Cumulative)
			assert.Equal(t, tt.input.PublicSecondVaccinationReceived, pub.Dose2.Daily)
			assert.Equal(t, tt.input.CumulativePublicSecondVaccinationReceived, pub.Dose2.Cumulative)

			// teenager group
			teen := result.Groups["teenager"]
			assert.Equal(t, tt.input.TeenagerVaccinationTarget, teen.Target)
			assert.Equal(t, tt.input.TeenagerFirstVaccinationReceived, teen.Dose1.Daily)
			assert.Equal(t, tt.input.CumulativeTeenagerFirstVaccinationReceived, teen.Dose1.Cumulative)
			assert.Equal(t, tt.input.TeenagerSecondVaccinationReceived, teen.Dose2.Daily)
			assert.Equal(t, tt.input.CumulativeTeenagerSecondVaccinationReceived, teen.Dose2.Cumulative)
		})
	}
}

func TestTransformProvinceVaccine(t *testing.T) {
	tests := []struct {
		name  string
		input models.ProvinceVaccine
	}{
		{
			name:  "zero values",
			input: models.ProvinceVaccine{},
		},
		{
			name: "sample real-world data with province_id",
			input: models.ProvinceVaccine{
				NationalVaccine: sampleNationalVaccine(),
				ProvinceID:      72,
			},
		},
		{
			name: "different province ID",
			input: models.ProvinceVaccine{
				NationalVaccine: sampleNationalVaccine(),
				ProvinceID:      31,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TransformProvinceVaccine(tt.input)

			// Province ID should be included
			assert.Equal(t, tt.input.ProvinceID, result.ProvinceID, "province_id should be included")

			// All national vaccine fields should be correctly mapped
			assert.Equal(t, tt.input.ID, result.ID)
			assert.Equal(t, tt.input.Day, result.Day)
			assert.Equal(t, tt.input.TotalVaccinationTarget, result.Target)

			// Dose fields
			assert.Equal(t, tt.input.FirstVaccinationReceived, result.Dose1.Daily)
			assert.Equal(t, tt.input.CumulativeFirstVaccinationReceived, result.Dose1.Cumulative)
			assert.Equal(t, tt.input.SecondVaccinationReceived, result.Dose2.Daily)
			assert.Equal(t, tt.input.CumulativeSecondVaccinationReceived, result.Dose2.Cumulative)

			// Groups
			assert.Len(t, result.Groups, 5)
			requiredGroups := []string{"health_worker", "elderly", "public_officer", "public", "teenager"}
			for _, key := range requiredGroups {
				_, ok := result.Groups[key]
				assert.True(t, ok, "groups should contain key: %s", key)
			}
		})
	}
}
