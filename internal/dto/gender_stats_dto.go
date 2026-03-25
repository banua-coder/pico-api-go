package dto

import "github.com/banua-coder/pico-api-go/internal/models"

// AgeGroups represents COVID-19 case counts broken down by age group
type AgeGroups struct {
	Gr0_14  int `json:"0_14"`
	Gr15_19 int `json:"15_19"`
	Gr20_24 int `json:"20_24"`
	Gr25_49 int `json:"25_49"`
	Gr50_54 int `json:"50_54"`
	Gr55Plus int `json:"55_plus"`
}

// GenderData represents total and age-group breakdown for a single gender
type GenderData struct {
	Total     int       `json:"total"`
	AgeGroups AgeGroups `json:"age_groups"`
}

// CategoryData represents male and female data for a category (positive/pdp)
type CategoryData struct {
	Male   GenderData `json:"male"`
	Female GenderData `json:"female"`
}

// GenderStatsResponse is the restructured API response for gender/age stats
// @Description Nested gender and age group statistics
type GenderStatsResponse struct {
	ID         int64        `json:"id"`
	Day        int64        `json:"day"`
	ProvinceID int          `json:"province_id"`
	Positive   CategoryData `json:"positive"`
	PDP        CategoryData `json:"pdp"`
}

// ToGenderStatsResponse transforms a ProvinceGenderCase model to GenderStatsResponse DTO
func ToGenderStatsResponse(g models.ProvinceGenderCase) GenderStatsResponse {
	return GenderStatsResponse{
		ID:         g.ID,
		Day:        g.Day,
		ProvinceID: g.ProvinceID,
		Positive: CategoryData{
			Male: GenderData{
				Total: g.PositiveMale,
				AgeGroups: AgeGroups{
					Gr0_14:   g.PositiveMale0_14,
					Gr15_19:  g.PositiveMale15_19,
					Gr20_24:  g.PositiveMale20_24,
					Gr25_49:  g.PositiveMale25_49,
					Gr50_54:  g.PositiveMale50_54,
					Gr55Plus: g.PositiveMale55,
				},
			},
			Female: GenderData{
				Total: g.PositiveFemale,
				AgeGroups: AgeGroups{
					Gr0_14:   g.PositiveFemale0_14,
					Gr15_19:  g.PositiveFemale15_19,
					Gr20_24:  g.PositiveFemale20_24,
					Gr25_49:  g.PositiveFemale25_49,
					Gr50_54:  g.PositiveFemale50_54,
					Gr55Plus: g.PositiveFemale55,
				},
			},
		},
		PDP: CategoryData{
			Male: GenderData{
				Total: g.PDPMale,
				AgeGroups: AgeGroups{
					Gr0_14:   g.PDPMale0_14,
					Gr15_19:  g.PDPMale15_19,
					Gr20_24:  g.PDPMale20_24,
					Gr25_49:  g.PDPMale25_49,
					Gr50_54:  g.PDPMale50_54,
					Gr55Plus: g.PDPMale55,
				},
			},
			Female: GenderData{
				Total: g.PDPFemale,
				AgeGroups: AgeGroups{
					Gr0_14:   g.PDPFemale0_14,
					Gr15_19:  g.PDPFemale15_19,
					Gr20_24:  g.PDPFemale20_24,
					Gr25_49:  g.PDPFemale25_49,
					Gr50_54:  g.PDPFemale50_54,
					Gr55Plus: g.PDPFemale55,
				},
			},
		},
	}
}

// ToGenderStatsResponseList transforms a slice of ProvinceGenderCase to GenderStatsResponse slice
func ToGenderStatsResponseList(cases []models.ProvinceGenderCase) []GenderStatsResponse {
	result := make([]GenderStatsResponse, len(cases))
	for i, c := range cases {
		result[i] = ToGenderStatsResponse(c)
	}
	return result
}
