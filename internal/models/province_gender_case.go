package models

// ProvinceGenderCase represents COVID-19 cases by gender and age group for a province
type ProvinceGenderCase struct {
	ID         int64 `json:"id"`
	Day        int64 `json:"day"`
	ProvinceID int   `json:"province_id"`

	PositiveMale   int `json:"positive_male"`
	PositiveFemale int `json:"positive_female"`
	PDPMale        int `json:"pdp_male"`
	PDPFemale      int `json:"pdp_female"`

	// Positive by age group - Male
	PositiveMale0_14  int `json:"positive_male_0_14"`
	PositiveMale15_19 int `json:"positive_male_15_19"`
	PositiveMale20_24 int `json:"positive_male_20_24"`
	PositiveMale25_49 int `json:"positive_male_25_49"`
	PositiveMale50_54 int `json:"positive_male_50_54"`
	PositiveMale55    int `json:"positive_male_55"`

	// Positive by age group - Female
	PositiveFemale0_14  int `json:"positive_female_0_14"`
	PositiveFemale15_19 int `json:"positive_female_15_19"`
	PositiveFemale20_24 int `json:"positive_female_20_24"`
	PositiveFemale25_49 int `json:"positive_female_25_49"`
	PositiveFemale50_54 int `json:"positive_female_50_54"`
	PositiveFemale55    int `json:"positive_female_55"`

	// PDP by age group - Male
	PDPMale0_14  int `json:"pdp_male_0_14"`
	PDPMale15_19 int `json:"pdp_male_15_19"`
	PDPMale20_24 int `json:"pdp_male_20_24"`
	PDPMale25_49 int `json:"pdp_male_25_49"`
	PDPMale50_54 int `json:"pdp_male_50_54"`
	PDPMale55    int `json:"pdp_male_55"`

	// PDP by age group - Female
	PDPFemale0_14  int `json:"pdp_female_0_14"`
	PDPFemale15_19 int `json:"pdp_female_15_19"`
	PDPFemale20_24 int `json:"pdp_female_20_24"`
	PDPFemale25_49 int `json:"pdp_female_25_49"`
	PDPFemale50_54 int `json:"pdp_female_50_54"`
	PDPFemale55    int `json:"pdp_female_55"`
}
