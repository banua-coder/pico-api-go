package models

import "time"

// NationalCaseResponse represents the structured response for national COVID-19 case data
type NationalCaseResponse struct {
	Day        int64                        `json:"day"`
	Date       time.Time                    `json:"date"`
	Daily      DailyCases                   `json:"daily"`
	Cumulative CumulativeCases              `json:"cumulative"`
	Statistics NationalCaseStatistics       `json:"statistics"`
}

// DailyCases represents new cases for a single day
type DailyCases struct {
	Positive  int64 `json:"positive"`
	Recovered int64 `json:"recovered"`
	Deceased  int64 `json:"deceased"`
	Active    int64 `json:"active"`
}

// CumulativeCases represents total cases accumulated over time
type CumulativeCases struct {
	Positive  int64 `json:"positive"`
	Recovered int64 `json:"recovered"`
	Deceased  int64 `json:"deceased"`
	Active    int64 `json:"active"`
}

// NationalCaseStatistics contains calculated statistics and metrics
type NationalCaseStatistics struct {
	Percentages      CasePercentages      `json:"percentages"`
	ReproductionRate *ReproductionRate    `json:"reproduction_rate,omitempty"`
}

// CasePercentages represents percentage distribution of cases
type CasePercentages struct {
	Active    float64 `json:"active"`
	Recovered float64 `json:"recovered"`
	Deceased  float64 `json:"deceased"`
}

// ReproductionRate represents the R-value with confidence bounds
type ReproductionRate struct {
	Value      *float64 `json:"value"`
	UpperBound *float64 `json:"upper_bound"`
	LowerBound *float64 `json:"lower_bound"`
}

// TransformToResponse converts a NationalCase model to the response format
func (nc *NationalCase) TransformToResponse() NationalCaseResponse {
	// Calculate active cases
	dailyActive := nc.Positive - nc.Recovered - nc.Deceased
	cumulativeActive := nc.CumulativePositive - nc.CumulativeRecovered - nc.CumulativeDeceased

	// Build response
	response := NationalCaseResponse{
		Day:  nc.Day,
		Date: nc.Date,
		Daily: DailyCases{
			Positive:  nc.Positive,
			Recovered: nc.Recovered,
			Deceased:  nc.Deceased,
			Active:    dailyActive,
		},
		Cumulative: CumulativeCases{
			Positive:  nc.CumulativePositive,
			Recovered: nc.CumulativeRecovered,
			Deceased:  nc.CumulativeDeceased,
			Active:    cumulativeActive,
		},
		Statistics: NationalCaseStatistics{
			Percentages: calculatePercentages(nc.CumulativePositive, nc.CumulativeRecovered, nc.CumulativeDeceased, cumulativeActive),
		},
	}

	// Always include reproduction rate structure, even when values are null
	response.Statistics.ReproductionRate = &ReproductionRate{
		Value:      nc.Rt,
		UpperBound: nc.RtUpper,
		LowerBound: nc.RtLower,
	}

	return response
}

// TransformSliceToResponse converts a slice of NationalCase models to response format
func TransformSliceToResponse(cases []NationalCase) []NationalCaseResponse {
	responses := make([]NationalCaseResponse, len(cases))
	for i, c := range cases {
		responses[i] = c.TransformToResponse()
	}
	return responses
}

// calculatePercentages calculates the percentage distribution of cases
func calculatePercentages(positive, recovered, deceased, active int64) CasePercentages {
	if positive == 0 {
		return CasePercentages{
			Active:    0,
			Recovered: 0,
			Deceased:  0,
		}
	}

	total := float64(positive)
	return CasePercentages{
		Active:    (float64(active) / total) * 100,
		Recovered: (float64(recovered) / total) * 100,
		Deceased:  (float64(deceased) / total) * 100,
	}
}