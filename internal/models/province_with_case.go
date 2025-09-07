package models

// ProvinceWithLatestCase represents a province with its latest COVID-19 case data
type ProvinceWithLatestCase struct {
	Province
	LatestCase *ProvinceCaseResponse `json:"latest_case,omitempty"`
}