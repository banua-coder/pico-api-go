package models

import "time"

// ProvinceCaseResponse represents the structured response for province COVID-19 case data
type ProvinceCaseResponse struct {
	Day        int64                   `json:"day"`
	Date       time.Time               `json:"date"`
	Daily      ProvinceDailyCases      `json:"daily"`
	Cumulative ProvinceCumulativeCases `json:"cumulative"`
	Statistics ProvinceCaseStatistics  `json:"statistics"`
	Province   *Province               `json:"province,omitempty"`
}

// ProvinceDailyCases represents new cases for a single day in a province
type ProvinceDailyCases struct {
	Positive  int64                `json:"positive"`
	Recovered int64                `json:"recovered"`
	Deceased  int64                `json:"deceased"`
	Active    int64                `json:"active"`
	ODP       DailyObservationData `json:"odp"`
	PDP       DailySupervisionData `json:"pdp"`
}

// ProvinceCumulativeCases represents total cases accumulated over time in a province
type ProvinceCumulativeCases struct {
	Positive  int64           `json:"positive"`
	Recovered int64           `json:"recovered"`
	Deceased  int64           `json:"deceased"`
	Active    int64           `json:"active"`
	ODP       ObservationData `json:"odp"`
	PDP       SupervisionData `json:"pdp"`
}

// DailyObservationData represents daily Person Under Observation (ODP) data
type DailyObservationData struct {
	Active   int64 `json:"active"`
	Finished int64 `json:"finished"`
}

// DailySupervisionData represents daily Patient Under Supervision (PDP) data
type DailySupervisionData struct {
	Active   int64 `json:"active"`
	Finished int64 `json:"finished"`
}

// ObservationData represents cumulative Person Under Observation (ODP) data
type ObservationData struct {
	Active   int64 `json:"active"`
	Finished int64 `json:"finished"`
	Total    int64 `json:"total"`
}

// SupervisionData represents cumulative Patient Under Supervision (PDP) data
type SupervisionData struct {
	Active   int64 `json:"active"`
	Finished int64 `json:"finished"`
	Total    int64 `json:"total"`
}

// ProvinceCaseStatistics contains calculated statistics and metrics for province data
type ProvinceCaseStatistics struct {
	Percentages      CasePercentages   `json:"percentages"`
	ReproductionRate *ReproductionRate `json:"reproduction_rate"`
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
			Positive:  pc.Positive,
			Recovered: pc.Recovered,
			Deceased:  pc.Deceased,
			Active:    dailyActive,
			ODP: DailyObservationData{
				Active:   pc.PersonUnderObservation - pc.FinishedPersonUnderObservation,
				Finished: pc.FinishedPersonUnderObservation,
			},
			PDP: DailySupervisionData{
				Active:   pc.PersonUnderSupervision - pc.FinishedPersonUnderSupervision,
				Finished: pc.FinishedPersonUnderSupervision,
			},
		},
		Cumulative: ProvinceCumulativeCases{
			Positive:  pc.CumulativePositive,
			Recovered: pc.CumulativeRecovered,
			Deceased:  pc.CumulativeDeceased,
			Active:    cumulativeActive,
			ODP: ObservationData{
				Active:   activePersonUnderObservation,
				Finished: pc.CumulativeFinishedPersonUnderObservation,
				Total:    pc.CumulativePersonUnderObservation,
			},
			PDP: SupervisionData{
				Active:   activePersonUnderSupervision,
				Finished: pc.CumulativeFinishedPersonUnderSupervision,
				Total:    pc.CumulativePersonUnderSupervision,
			},
		},
		Statistics: ProvinceCaseStatistics{
			Percentages: calculatePercentages(pc.CumulativePositive, pc.CumulativeRecovered, pc.CumulativeDeceased, cumulativeActive),
		},
		Province: pc.Province,
	}

	// Always include reproduction rate structure, even when values are null
	response.Statistics.ReproductionRate = &ReproductionRate{
		Value:      pc.Rt,
		UpperBound: pc.RtUpper,
		LowerBound: pc.RtLower,
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
