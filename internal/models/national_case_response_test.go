package models

import (
	"testing"
	"time"
)

func TestNationalCase_TransformToResponse(t *testing.T) {
	// Test case with reproduction rate data
	rtValue := 1.2
	rtUpper := 1.5
	rtLower := 0.9

	nc := NationalCase{
		ID:                  1,
		Day:                 10,
		Date:                time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
		Positive:            50,
		Recovered:           10,
		Deceased:            5,
		CumulativePositive:  500,
		CumulativeRecovered: 100,
		CumulativeDeceased:  20,
		Rt:                  &rtValue,
		RtUpper:             &rtUpper,
		RtLower:             &rtLower,
	}

	response := nc.TransformToResponse()

	// Check basic fields
	if response.Day != nc.Day {
		t.Errorf("Expected Day %d, got %d", nc.Day, response.Day)
	}

	if !response.Date.Equal(nc.Date) {
		t.Errorf("Expected Date %v, got %v", nc.Date, response.Date)
	}

	// Check daily cases
	if response.Daily.Positive != nc.Positive {
		t.Errorf("Expected Daily.Positive %d, got %d", nc.Positive, response.Daily.Positive)
	}

	if response.Daily.Recovered != nc.Recovered {
		t.Errorf("Expected Daily.Recovered %d, got %d", nc.Recovered, response.Daily.Recovered)
	}

	if response.Daily.Deceased != nc.Deceased {
		t.Errorf("Expected Daily.Deceased %d, got %d", nc.Deceased, response.Daily.Deceased)
	}

	expectedDailyActive := nc.Positive - nc.Recovered - nc.Deceased
	if response.Daily.Active != expectedDailyActive {
		t.Errorf("Expected Daily.Active %d, got %d", expectedDailyActive, response.Daily.Active)
	}

	// Check cumulative cases
	if response.Cumulative.Positive != nc.CumulativePositive {
		t.Errorf("Expected Cumulative.Positive %d, got %d", nc.CumulativePositive, response.Cumulative.Positive)
	}

	expectedCumulativeActive := nc.CumulativePositive - nc.CumulativeRecovered - nc.CumulativeDeceased
	if response.Cumulative.Active != expectedCumulativeActive {
		t.Errorf("Expected Cumulative.Active %d, got %d", expectedCumulativeActive, response.Cumulative.Active)
	}

	// Check reproduction rate
	if response.Statistics.ReproductionRate == nil {
		t.Error("Expected ReproductionRate to be present")
	} else {
		if response.Statistics.ReproductionRate.Value != rtValue {
			t.Errorf("Expected Rt.Value %f, got %f", rtValue, response.Statistics.ReproductionRate.Value)
		}
		if response.Statistics.ReproductionRate.UpperBound != rtUpper {
			t.Errorf("Expected Rt.UpperBound %f, got %f", rtUpper, response.Statistics.ReproductionRate.UpperBound)
		}
		if response.Statistics.ReproductionRate.LowerBound != rtLower {
			t.Errorf("Expected Rt.LowerBound %f, got %f", rtLower, response.Statistics.ReproductionRate.LowerBound)
		}
	}

	// Check percentages
	expectedActivePercentage := (float64(expectedCumulativeActive) / float64(nc.CumulativePositive)) * 100
	if response.Statistics.Percentages.Active != expectedActivePercentage {
		t.Errorf("Expected Active percentage %f, got %f", expectedActivePercentage, response.Statistics.Percentages.Active)
	}
}

func TestNationalCase_TransformToResponse_NoReproductionRate(t *testing.T) {
	nc := NationalCase{
		ID:                  1,
		Day:                 1,
		Date:                time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
		Positive:            2,
		Recovered:           0,
		Deceased:            0,
		CumulativePositive:  2,
		CumulativeRecovered: 0,
		CumulativeDeceased:  0,
		// No Rt values
	}

	response := nc.TransformToResponse()

	if response.Statistics.ReproductionRate != nil {
		t.Error("Expected ReproductionRate to be nil when not provided")
	}
}

func TestCalculatePercentages_ZeroPositive(t *testing.T) {
	percentages := calculatePercentages(0, 0, 0, 0)

	if percentages.Active != 0 {
		t.Errorf("Expected Active percentage 0, got %f", percentages.Active)
	}
	if percentages.Recovered != 0 {
		t.Errorf("Expected Recovered percentage 0, got %f", percentages.Recovered)
	}
	if percentages.Deceased != 0 {
		t.Errorf("Expected Deceased percentage 0, got %f", percentages.Deceased)
	}
}

func TestTransformSliceToResponse(t *testing.T) {
	cases := []NationalCase{
		{
			ID:                  1,
			Day:                 1,
			Date:                time.Date(2020, 3, 1, 0, 0, 0, 0, time.UTC),
			Positive:            2,
			Recovered:           0,
			Deceased:            0,
			CumulativePositive:  2,
			CumulativeRecovered: 0,
			CumulativeDeceased:  0,
		},
		{
			ID:                  2,
			Day:                 2,
			Date:                time.Date(2020, 3, 2, 0, 0, 0, 0, time.UTC),
			Positive:            3,
			Recovered:           1,
			Deceased:            0,
			CumulativePositive:  5,
			CumulativeRecovered: 1,
			CumulativeDeceased:  0,
		},
	}

	responses := TransformSliceToResponse(cases)

	if len(responses) != len(cases) {
		t.Errorf("Expected %d responses, got %d", len(cases), len(responses))
	}

	for i, response := range responses {
		if response.Day != cases[i].Day {
			t.Errorf("Response %d: Expected Day %d, got %d", i, cases[i].Day, response.Day)
		}
	}
}