package dto

import (
	"math"
	"testing"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
)

func sampleProvinceGenderCase() models.ProvinceGenderCase {
	return models.ProvinceGenderCase{
		ID:         42,
		Day:        200,
		ProvinceID: 72,

		PositiveMale:   1500,
		PositiveFemale: 1200,
		PDPMale:        300,
		PDPFemale:      250,

		PositiveMale0_14:  100,
		PositiveMale15_19: 150,
		PositiveMale20_24: 200,
		PositiveMale25_49: 800,
		PositiveMale50_54: 150,
		PositiveMale55:    100,

		PositiveFemale0_14:  80,
		PositiveFemale15_19: 120,
		PositiveFemale20_24: 180,
		PositiveFemale25_49: 650,
		PositiveFemale50_54: 100,
		PositiveFemale55:    70,

		PDPMale0_14:  20,
		PDPMale15_19: 30,
		PDPMale20_24: 50,
		PDPMale25_49: 150,
		PDPMale50_54: 30,
		PDPMale55:    20,

		PDPFemale0_14:  15,
		PDPFemale15_19: 25,
		PDPFemale20_24: 45,
		PDPFemale25_49: 120,
		PDPFemale50_54: 25,
		PDPFemale55:    20,
	}
}

func TestToGenderStatsResponse(t *testing.T) {
	maxVal := math.MaxInt

	tests := []struct {
		name  string
		input models.ProvinceGenderCase
	}{
		{
			name:  "zero values",
			input: models.ProvinceGenderCase{},
		},
		{
			name:  "sample real-world data",
			input: sampleProvinceGenderCase(),
		},
		{
			name: "max int values",
			input: models.ProvinceGenderCase{
				ID:                 math.MaxInt64,
				Day:                math.MaxInt64,
				ProvinceID:         maxVal,
				PositiveMale:       maxVal,
				PositiveFemale:     maxVal,
				PDPMale:            maxVal,
				PDPFemale:          maxVal,
				PositiveMale0_14:   maxVal,
				PositiveMale15_19:  maxVal,
				PositiveMale20_24:  maxVal,
				PositiveMale25_49:  maxVal,
				PositiveMale50_54:  maxVal,
				PositiveMale55:     maxVal,
				PositiveFemale0_14:  maxVal,
				PositiveFemale15_19: maxVal,
				PositiveFemale20_24: maxVal,
				PositiveFemale25_49: maxVal,
				PositiveFemale50_54: maxVal,
				PositiveFemale55:    maxVal,
				PDPMale0_14:        maxVal,
				PDPMale15_19:       maxVal,
				PDPMale20_24:       maxVal,
				PDPMale25_49:       maxVal,
				PDPMale50_54:       maxVal,
				PDPMale55:          maxVal,
				PDPFemale0_14:      maxVal,
				PDPFemale15_19:     maxVal,
				PDPFemale20_24:     maxVal,
				PDPFemale25_49:     maxVal,
				PDPFemale50_54:     maxVal,
				PDPFemale55:        maxVal,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToGenderStatsResponse(tt.input)

			// Basic fields
			assert.Equal(t, tt.input.ID, result.ID)
			assert.Equal(t, tt.input.Day, result.Day)
			assert.Equal(t, tt.input.ProvinceID, result.ProvinceID)

			// Positive.Male.Total
			assert.Equal(t, tt.input.PositiveMale, result.Positive.Male.Total, "positive.male.total should match PositiveMale")
			// Positive.Female.Total
			assert.Equal(t, tt.input.PositiveFemale, result.Positive.Female.Total, "positive.female.total should match PositiveFemale")

			// All 6 age groups in positive.male
			assert.Equal(t, tt.input.PositiveMale0_14, result.Positive.Male.AgeGroups.Gr0_14, "positive.male.age_groups.0_14")
			assert.Equal(t, tt.input.PositiveMale15_19, result.Positive.Male.AgeGroups.Gr15_19, "positive.male.age_groups.15_19")
			assert.Equal(t, tt.input.PositiveMale20_24, result.Positive.Male.AgeGroups.Gr20_24, "positive.male.age_groups.20_24")
			assert.Equal(t, tt.input.PositiveMale25_49, result.Positive.Male.AgeGroups.Gr25_49, "positive.male.age_groups.25_49")
			assert.Equal(t, tt.input.PositiveMale50_54, result.Positive.Male.AgeGroups.Gr50_54, "positive.male.age_groups.50_54")
			assert.Equal(t, tt.input.PositiveMale55, result.Positive.Male.AgeGroups.Gr55Plus, "positive.male.age_groups.55_plus")

			// All 6 age groups in positive.female
			assert.Equal(t, tt.input.PositiveFemale0_14, result.Positive.Female.AgeGroups.Gr0_14, "positive.female.age_groups.0_14")
			assert.Equal(t, tt.input.PositiveFemale15_19, result.Positive.Female.AgeGroups.Gr15_19, "positive.female.age_groups.15_19")
			assert.Equal(t, tt.input.PositiveFemale20_24, result.Positive.Female.AgeGroups.Gr20_24, "positive.female.age_groups.20_24")
			assert.Equal(t, tt.input.PositiveFemale25_49, result.Positive.Female.AgeGroups.Gr25_49, "positive.female.age_groups.25_49")
			assert.Equal(t, tt.input.PositiveFemale50_54, result.Positive.Female.AgeGroups.Gr50_54, "positive.female.age_groups.50_54")
			assert.Equal(t, tt.input.PositiveFemale55, result.Positive.Female.AgeGroups.Gr55Plus, "positive.female.age_groups.55_plus")

			// PDP category mapping
			assert.Equal(t, tt.input.PDPMale, result.PDP.Male.Total, "pdp.male.total should match PDPMale")
			assert.Equal(t, tt.input.PDPFemale, result.PDP.Female.Total, "pdp.female.total should match PDPFemale")

			// All 6 age groups in pdp.male
			assert.Equal(t, tt.input.PDPMale0_14, result.PDP.Male.AgeGroups.Gr0_14, "pdp.male.age_groups.0_14")
			assert.Equal(t, tt.input.PDPMale15_19, result.PDP.Male.AgeGroups.Gr15_19, "pdp.male.age_groups.15_19")
			assert.Equal(t, tt.input.PDPMale20_24, result.PDP.Male.AgeGroups.Gr20_24, "pdp.male.age_groups.20_24")
			assert.Equal(t, tt.input.PDPMale25_49, result.PDP.Male.AgeGroups.Gr25_49, "pdp.male.age_groups.25_49")
			assert.Equal(t, tt.input.PDPMale50_54, result.PDP.Male.AgeGroups.Gr50_54, "pdp.male.age_groups.50_54")
			assert.Equal(t, tt.input.PDPMale55, result.PDP.Male.AgeGroups.Gr55Plus, "pdp.male.age_groups.55_plus")

			// All 6 age groups in pdp.female
			assert.Equal(t, tt.input.PDPFemale0_14, result.PDP.Female.AgeGroups.Gr0_14, "pdp.female.age_groups.0_14")
			assert.Equal(t, tt.input.PDPFemale15_19, result.PDP.Female.AgeGroups.Gr15_19, "pdp.female.age_groups.15_19")
			assert.Equal(t, tt.input.PDPFemale20_24, result.PDP.Female.AgeGroups.Gr20_24, "pdp.female.age_groups.20_24")
			assert.Equal(t, tt.input.PDPFemale25_49, result.PDP.Female.AgeGroups.Gr25_49, "pdp.female.age_groups.25_49")
			assert.Equal(t, tt.input.PDPFemale50_54, result.PDP.Female.AgeGroups.Gr50_54, "pdp.female.age_groups.50_54")
			assert.Equal(t, tt.input.PDPFemale55, result.PDP.Female.AgeGroups.Gr55Plus, "pdp.female.age_groups.55_plus")
		})
	}
}

func TestToGenderStatsResponseList(t *testing.T) {
	t.Run("empty slice returns empty slice", func(t *testing.T) {
		result := ToGenderStatsResponseList([]models.ProvinceGenderCase{})
		assert.NotNil(t, result)
		assert.Len(t, result, 0)
	})

	t.Run("batch transform preserves order and count", func(t *testing.T) {
		cases := []models.ProvinceGenderCase{
			{ID: 1, Day: 1, ProvinceID: 72, PositiveMale: 100, PositiveFemale: 80},
			{ID: 2, Day: 2, ProvinceID: 72, PositiveMale: 200, PositiveFemale: 160},
			{ID: 3, Day: 3, ProvinceID: 72, PositiveMale: 300, PositiveFemale: 240},
		}

		result := ToGenderStatsResponseList(cases)

		assert.Len(t, result, 3)
		for i, c := range cases {
			assert.Equal(t, c.ID, result[i].ID, "ID should match at index %d", i)
			assert.Equal(t, c.Day, result[i].Day, "Day should match at index %d", i)
			assert.Equal(t, c.ProvinceID, result[i].ProvinceID, "ProvinceID should match at index %d", i)
			assert.Equal(t, c.PositiveMale, result[i].Positive.Male.Total, "positive.male.total at index %d", i)
			assert.Equal(t, c.PositiveFemale, result[i].Positive.Female.Total, "positive.female.total at index %d", i)
		}
	})

	t.Run("single item list", func(t *testing.T) {
		cases := []models.ProvinceGenderCase{sampleProvinceGenderCase()}
		result := ToGenderStatsResponseList(cases)
		assert.Len(t, result, 1)
		assert.Equal(t, cases[0].PositiveMale, result[0].Positive.Male.Total)
		assert.Equal(t, cases[0].PositiveFemale, result[0].Positive.Female.Total)
	})
}
