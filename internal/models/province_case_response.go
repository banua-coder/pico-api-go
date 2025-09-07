package models

import "time"

// ProvinceCaseResponse represents the structured response for province COVID-19 case data
type ProvinceCaseResponse struct {
	Day        int64                     `json:"day"`
	Date       time.Time                 `json:"date"`
	Daily      ProvinceDailyCases        `json:"daily"`
	Cumulative ProvinceCumulativeCases   `json:"cumulative"`
	Statistics ProvinceCaseStatistics    `json:"statistics"`
	Province   *Province                 `json:"province,omitempty"`
}

// ProvinceDailyCases represents new cases for a single day in a province
type ProvinceDailyCases struct {
	Positive                         int64 `json:"positive"`
	Recovered                        int64 `json:"recovered"`
	Deceased                         int64 `json:"deceased"`
	Active                           int64 `json:"active"`
	PersonUnderObservation           int64 `json:"person_under_observation"`
	FinishedPersonUnderObservation   int64 `json:"finished_person_under_observation"`
	PersonUnderSupervision           int64 `json:"person_under_supervision"`
	FinishedPersonUnderSupervision   int64 `json:"finished_person_under_supervision"`
}

// ProvinceCumulativeCases represents total cases accumulated over time in a province
type ProvinceCumulativeCases struct {
	Positive                         int64 `json:"positive"`
	Recovered                        int64 `json:"recovered"`
	Deceased                         int64 `json:"deceased"`
	Active                           int64 `json:"active"`
	PersonUnderObservation           int64 `json:"person_under_observation"`
	ActivePersonUnderObservation     int64 `json:"active_person_under_observation"`
	FinishedPersonUnderObservation   int64 `json:"finished_person_under_observation"`
	PersonUnderSupervision           int64 `json:"person_under_supervision"`
	ActivePersonUnderSupervision     int64 `json:"active_person_under_supervision"`
	FinishedPersonUnderSupervision   int64 `json:"finished_person_under_supervision"`
}

// ProvinceCaseStatistics contains calculated statistics and metrics for province data
type ProvinceCaseStatistics struct {
	Percentages      CasePercentages   `json:"percentages"`
	ReproductionRate *ReproductionRate `json:"reproduction_rate,omitempty"`
}

// TransformToResponse converts a ProvinceCase model to the response format
func (pc *ProvinceCase) TransformToResponse(date time.Time) ProvinceCaseResponse {
	// Calculate active cases
	dailyActive := pc.Positive - pc.Recovered - pc.Deceased
	cumulativeActive := pc.CumulativePositive - pc.CumulativeRecovered - pc.CumulativeDeceased

	// Calculate active under observation and supervision
	activePersonUnderObservation := pc.CumulativePersonUnderObservation - pc.CumulativeFinishedPersonUnderObservation
	activePersonUnderSupervision := pc.CumulativePersonUnderSupervision - pc.CumulativeFinishedPersonUnderSupervision

	// Build response
	response := ProvinceCaseResponse{
		Day:  pc.Day,
		Date: date,
		Daily: ProvinceDailyCases{
			Positive:                         pc.Positive,
			Recovered:                        pc.Recovered,
			Deceased:                         pc.Deceased,
			Active:                           dailyActive,
			PersonUnderObservation:           pc.PersonUnderObservation,
			FinishedPersonUnderObservation:   pc.FinishedPersonUnderObservation,
			PersonUnderSupervision:           pc.PersonUnderSupervision,
			FinishedPersonUnderSupervision:   pc.FinishedPersonUnderSupervision,
		},
		Cumulative: ProvinceCumulativeCases{
			Positive:                         pc.CumulativePositive,
			Recovered:                        pc.CumulativeRecovered,
			Deceased:                         pc.CumulativeDeceased,
			Active:                           cumulativeActive,
			PersonUnderObservation:           pc.CumulativePersonUnderObservation,
			ActivePersonUnderObservation:     activePersonUnderObservation,
			FinishedPersonUnderObservation:   pc.CumulativeFinishedPersonUnderObservation,
			PersonUnderSupervision:           pc.CumulativePersonUnderSupervision,
			ActivePersonUnderSupervision:     activePersonUnderSupervision,
			FinishedPersonUnderSupervision:   pc.CumulativeFinishedPersonUnderSupervision,
		},
		Statistics: ProvinceCaseStatistics{
			Percentages: calculatePercentages(pc.CumulativePositive, pc.CumulativeRecovered, pc.CumulativeDeceased, cumulativeActive),
		},
		Province: pc.Province,
	}

	// Add reproduction rate if available
	if pc.Rt != nil && pc.RtUpper != nil && pc.RtLower != nil {
		response.Statistics.ReproductionRate = &ReproductionRate{
			Value:      *pc.Rt,
			UpperBound: *pc.RtUpper,
			LowerBound: *pc.RtLower,
		}
	}

	return response
}

// TransformProvinceCaseWithDateToResponse converts a ProvinceCaseWithDate model to the response format
func (pcd *ProvinceCaseWithDate) TransformToResponse() ProvinceCaseResponse {
	return pcd.ProvinceCase.TransformToResponse(pcd.Date)
}

// TransformProvinceCaseSliceToResponse converts a slice of ProvinceCaseWithDate models to response format
func TransformProvinceCaseSliceToResponse(cases []ProvinceCaseWithDate) []ProvinceCaseResponse {
	responses := make([]ProvinceCaseResponse, len(cases))
	for i, c := range cases {
		responses[i] = c.TransformToResponse()
	}
	return responses
}