package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProvinceCase_TransformToResponse(t *testing.T) {
	testDate := time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC)
	rt := 1.5
	rtUpper := 1.8
	rtLower := 1.2

	tests := []struct {
		name           string
		provinceCase   ProvinceCase
		date           time.Time
		expectedResult ProvinceCaseResponse
	}{
		{
			name: "complete province case data",
			provinceCase: ProvinceCase{
				ID:                                       1,
				Day:                                      100,
				ProvinceID:                               "ID-JK",
				Positive:                                 150,
				Recovered:                                120,
				Deceased:                                 10,
				PersonUnderObservation:                   25,
				FinishedPersonUnderObservation:           20,
				PersonUnderSupervision:                   30,
				FinishedPersonUnderSupervision:           25,
				CumulativePositive:                       5000,
				CumulativeRecovered:                      4500,
				CumulativeDeceased:                       300,
				CumulativePersonUnderObservation:         800,
				CumulativeFinishedPersonUnderObservation: 750,
				CumulativePersonUnderSupervision:         600,
				CumulativeFinishedPersonUnderSupervision: 580,
				Rt:                                       &rt,
				RtUpper:                                  &rtUpper,
				RtLower:                                  &rtLower,
				Province: &Province{
					ID:   "ID-JK",
					Name: "DKI Jakarta",
				},
			},
			date: testDate,
			expectedResult: ProvinceCaseResponse{
				Day:  100,
				Date: testDate,
				Daily: ProvinceDailyCases{
					Positive:  150,
					Recovered: 120,
					Deceased:  10,
					Active:    20, // 150 - 120 - 10
					ODP: DailyObservationData{
						Active:   5, // 25 - 20
						Finished: 20,
					},
					PDP: DailySupervisionData{
						Active:   5, // 30 - 25
						Finished: 25,
					},
				},
				Cumulative: ProvinceCumulativeCases{
					Positive:  5000,
					Recovered: 4500,
					Deceased:  300,
					Active:    200, // 5000 - 4500 - 300
					ODP: ObservationData{
						Active:   50, // 800 - 750
						Finished: 750,
						Total:    800,
					},
					PDP: SupervisionData{
						Active:   20, // 600 - 580
						Finished: 580,
						Total:    600,
					},
				},
				Statistics: ProvinceCaseStatistics{
					Percentages: CasePercentages{
						Active:    4.0,  // (200 / 5000) * 100
						Recovered: 90.0, // (4500 / 5000) * 100
						Deceased:  6.0,  // (300 / 5000) * 100
					},
					ReproductionRate: &ReproductionRate{
						Value:      &[]float64{1.5}[0],
						UpperBound: &[]float64{1.8}[0],
						LowerBound: &[]float64{1.2}[0],
					},
				},
				Province: &Province{
					ID:   "ID-JK",
					Name: "DKI Jakarta",
				},
			},
		},
		{
			name: "province case without reproduction rate",
			provinceCase: ProvinceCase{
				ID:                                       2,
				Day:                                      50,
				ProvinceID:                               "ID-JB",
				Positive:                                 100,
				Recovered:                                80,
				Deceased:                                 5,
				PersonUnderObservation:                   15,
				FinishedPersonUnderObservation:           10,
				PersonUnderSupervision:                   20,
				FinishedPersonUnderSupervision:           15,
				CumulativePositive:                       2000,
				CumulativeRecovered:                      1800,
				CumulativeDeceased:                       100,
				CumulativePersonUnderObservation:         400,
				CumulativeFinishedPersonUnderObservation: 350,
				CumulativePersonUnderSupervision:         300,
				CumulativeFinishedPersonUnderSupervision: 290,
				Rt:                                       nil,
				RtUpper:                                  nil,
				RtLower:                                  nil,
				Province: &Province{
					ID:   "ID-JB",
					Name: "Jawa Barat",
				},
			},
			date: testDate,
			expectedResult: ProvinceCaseResponse{
				Day:  50,
				Date: testDate,
				Daily: ProvinceDailyCases{
					Positive:  100,
					Recovered: 80,
					Deceased:  5,
					Active:    15, // 100 - 80 - 5
					ODP: DailyObservationData{
						Active:   5, // 15 - 10
						Finished: 10,
					},
					PDP: DailySupervisionData{
						Active:   5, // 20 - 15
						Finished: 15,
					},
				},
				Cumulative: ProvinceCumulativeCases{
					Positive:  2000,
					Recovered: 1800,
					Deceased:  100,
					Active:    100, // 2000 - 1800 - 100
					ODP: ObservationData{
						Active:   50, // 400 - 350
						Finished: 350,
						Total:    400,
					},
					PDP: SupervisionData{
						Active:   10, // 300 - 290
						Finished: 290,
						Total:    300,
					},
				},
				Statistics: ProvinceCaseStatistics{
					Percentages: CasePercentages{
						Active:    5.0,  // (100 / 2000) * 100
						Recovered: 90.0, // (1800 / 2000) * 100
						Deceased:  5.0,  // (100 / 2000) * 100
					},
					ReproductionRate: &ReproductionRate{
						Value:      nil,
						UpperBound: nil,
						LowerBound: nil,
					},
				},
				Province: &Province{
					ID:   "ID-JB",
					Name: "Jawa Barat",
				},
			},
		},
		{
			name: "province case with zero cumulative positive",
			provinceCase: ProvinceCase{
				ID:                                       3,
				Day:                                      1,
				ProvinceID:                               "ID-AC",
				Positive:                                 0,
				Recovered:                                0,
				Deceased:                                 0,
				PersonUnderObservation:                   0,
				FinishedPersonUnderObservation:           0,
				PersonUnderSupervision:                   0,
				FinishedPersonUnderSupervision:           0,
				CumulativePositive:                       0,
				CumulativeRecovered:                      0,
				CumulativeDeceased:                       0,
				CumulativePersonUnderObservation:         0,
				CumulativeFinishedPersonUnderObservation: 0,
				CumulativePersonUnderSupervision:         0,
				CumulativeFinishedPersonUnderSupervision: 0,
				Rt:                                       nil,
				RtUpper:                                  nil,
				RtLower:                                  nil,
				Province: &Province{
					ID:   "ID-AC",
					Name: "Aceh",
				},
			},
			date: testDate,
			expectedResult: ProvinceCaseResponse{
				Day:  1,
				Date: testDate,
				Daily: ProvinceDailyCases{
					Positive:  0,
					Recovered: 0,
					Deceased:  0,
					Active:    0,
					ODP: DailyObservationData{
						Active:   0,
						Finished: 0,
					},
					PDP: DailySupervisionData{
						Active:   0,
						Finished: 0,
					},
				},
				Cumulative: ProvinceCumulativeCases{
					Positive:  0,
					Recovered: 0,
					Deceased:  0,
					Active:    0,
					ODP: ObservationData{
						Active:   0,
						Finished: 0,
						Total:    0,
					},
					PDP: SupervisionData{
						Active:   0,
						Finished: 0,
						Total:    0,
					},
				},
				Statistics: ProvinceCaseStatistics{
					Percentages: CasePercentages{
						Active:    0.0,
						Recovered: 0.0,
						Deceased:  0.0,
					},
					ReproductionRate: &ReproductionRate{
						Value:      nil,
						UpperBound: nil,
						LowerBound: nil,
					},
				},
				Province: &Province{
					ID:   "ID-AC",
					Name: "Aceh",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.provinceCase.TransformToResponse(tt.date)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestProvinceCaseWithDate_TransformToResponse(t *testing.T) {
	testDate := time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC)
	rt := 1.2
	rtUpper := 1.5
	rtLower := 0.9

	provinceCaseWithDate := ProvinceCaseWithDate{
		ProvinceCase: ProvinceCase{
			ID:                                       1,
			Day:                                      200,
			ProvinceID:                               "ID-JT",
			Positive:                                 50,
			Recovered:                                40,
			Deceased:                                 2,
			PersonUnderObservation:                   10,
			FinishedPersonUnderObservation:           8,
			PersonUnderSupervision:                   12,
			FinishedPersonUnderSupervision:           10,
			CumulativePositive:                       3000,
			CumulativeRecovered:                      2700,
			CumulativeDeceased:                       200,
			CumulativePersonUnderObservation:         500,
			CumulativeFinishedPersonUnderObservation: 450,
			CumulativePersonUnderSupervision:         350,
			CumulativeFinishedPersonUnderSupervision: 320,
			Rt:                                       &rt,
			RtUpper:                                  &rtUpper,
			RtLower:                                  &rtLower,
			Province: &Province{
				ID:   "ID-JT",
				Name: "Jawa Tengah",
			},
		},
		Date: testDate,
	}

	result := provinceCaseWithDate.TransformToResponse()

	expected := ProvinceCaseResponse{
		Day:  200,
		Date: testDate,
		Daily: ProvinceDailyCases{
			Positive:  50,
			Recovered: 40,
			Deceased:  2,
			Active:    8, // 50 - 40 - 2
			ODP: DailyObservationData{
				Active:   2, // 10 - 8
				Finished: 8,
			},
			PDP: DailySupervisionData{
				Active:   2, // 12 - 10
				Finished: 10,
			},
		},
		Cumulative: ProvinceCumulativeCases{
			Positive:  3000,
			Recovered: 2700,
			Deceased:  200,
			Active:    100, // 3000 - 2700 - 200
			ODP: ObservationData{
				Active:   50, // 500 - 450
				Finished: 450,
				Total:    500,
			},
			PDP: SupervisionData{
				Active:   30, // 350 - 320
				Finished: 320,
				Total:    350,
			},
		},
		Statistics: ProvinceCaseStatistics{
			Percentages: CasePercentages{
				Active:    3.3333333333333335, // (100 / 3000) * 100
				Recovered: 90.0,               // (2700 / 3000) * 100
				Deceased:  6.666666666666667,  // (200 / 3000) * 100
			},
			ReproductionRate: &ReproductionRate{
				Value:      &[]float64{1.2}[0],
				UpperBound: &[]float64{1.5}[0],
				LowerBound: &[]float64{0.9}[0],
			},
		},
		Province: &Province{
			ID:   "ID-JT",
			Name: "Jawa Tengah",
		},
	}

	assert.Equal(t, expected, result)
}

func TestTransformProvinceCaseSliceToResponse(t *testing.T) {
	testDate1 := time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC)
	testDate2 := time.Date(2023, 10, 16, 0, 0, 0, 0, time.UTC)
	rt := 1.3
	rtUpper := 1.6
	rtLower := 1.0

	cases := []ProvinceCaseWithDate{
		{
			ProvinceCase: ProvinceCase{
				ID:                                       1,
				Day:                                      1,
				ProvinceID:                               "ID-JK",
				Positive:                                 100,
				Recovered:                                80,
				Deceased:                                 5,
				PersonUnderObservation:                   20,
				FinishedPersonUnderObservation:           15,
				PersonUnderSupervision:                   25,
				FinishedPersonUnderSupervision:           20,
				CumulativePositive:                       1000,
				CumulativeRecovered:                      800,
				CumulativeDeceased:                       50,
				CumulativePersonUnderObservation:         200,
				CumulativeFinishedPersonUnderObservation: 180,
				CumulativePersonUnderSupervision:         250,
				CumulativeFinishedPersonUnderSupervision: 230,
				Rt:                                       &rt,
				RtUpper:                                  &rtUpper,
				RtLower:                                  &rtLower,
				Province: &Province{
					ID:   "ID-JK",
					Name: "DKI Jakarta",
				},
			},
			Date: testDate1,
		},
		{
			ProvinceCase: ProvinceCase{
				ID:                                       2,
				Day:                                      2,
				ProvinceID:                               "ID-JK",
				Positive:                                 50,
				Recovered:                                45,
				Deceased:                                 2,
				PersonUnderObservation:                   10,
				FinishedPersonUnderObservation:           8,
				PersonUnderSupervision:                   12,
				FinishedPersonUnderSupervision:           10,
				CumulativePositive:                       1050,
				CumulativeRecovered:                      845,
				CumulativeDeceased:                       52,
				CumulativePersonUnderObservation:         210,
				CumulativeFinishedPersonUnderObservation: 188,
				CumulativePersonUnderSupervision:         262,
				CumulativeFinishedPersonUnderSupervision: 240,
				Rt:                                       &rt,
				RtUpper:                                  &rtUpper,
				RtLower:                                  &rtLower,
				Province: &Province{
					ID:   "ID-JK",
					Name: "DKI Jakarta",
				},
			},
			Date: testDate2,
		},
	}

	result := TransformProvinceCaseSliceToResponse(cases)

	assert.Len(t, result, 2)

	// Test first case
	assert.Equal(t, int64(1), result[0].Day)
	assert.Equal(t, testDate1, result[0].Date)
	assert.Equal(t, int64(100), result[0].Daily.Positive)
	assert.Equal(t, int64(15), result[0].Daily.Active) // 100 - 80 - 5
	assert.Equal(t, int64(1000), result[0].Cumulative.Positive)
	assert.Equal(t, int64(150), result[0].Cumulative.Active) // 1000 - 800 - 50

	// Test second case
	assert.Equal(t, int64(2), result[1].Day)
	assert.Equal(t, testDate2, result[1].Date)
	assert.Equal(t, int64(50), result[1].Daily.Positive)
	assert.Equal(t, int64(3), result[1].Daily.Active) // 50 - 45 - 2
	assert.Equal(t, int64(1050), result[1].Cumulative.Positive)
	assert.Equal(t, int64(153), result[1].Cumulative.Active) // 1050 - 845 - 52
}

func TestTransformProvinceCaseSliceToResponse_EmptySlice(t *testing.T) {
	var cases []ProvinceCaseWithDate
	result := TransformProvinceCaseSliceToResponse(cases)
	assert.Empty(t, result)
}

func TestProvinceCase_TransformToResponseWithoutProvince(t *testing.T) {
	testDate := time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC)
	rt := 1.5
	rtUpper := 1.8
	rtLower := 1.2

	provinceCase := ProvinceCase{
		ID:                                       1,
		Day:                                      100,
		ProvinceID:                               "ID-JK",
		Positive:                                 150,
		Recovered:                                120,
		Deceased:                                 10,
		PersonUnderObservation:                   25,
		FinishedPersonUnderObservation:           20,
		PersonUnderSupervision:                   30,
		FinishedPersonUnderSupervision:           25,
		CumulativePositive:                       5000,
		CumulativeRecovered:                      4500,
		CumulativeDeceased:                       300,
		CumulativePersonUnderObservation:         800,
		CumulativeFinishedPersonUnderObservation: 750,
		CumulativePersonUnderSupervision:         600,
		CumulativeFinishedPersonUnderSupervision: 580,
		Rt:                                       &rt,
		RtUpper:                                  &rtUpper,
		RtLower:                                  &rtLower,
		Province: &Province{
			ID:   "ID-JK",
			Name: "DKI Jakarta",
		},
	}

	result := provinceCase.TransformToResponseWithoutProvince(testDate)

	expectedResult := ProvinceCaseResponse{
		Day:  100,
		Date: testDate,
		Daily: ProvinceDailyCases{
			Positive:  150,
			Recovered: 120,
			Deceased:  10,
			Active:    20, // 150 - 120 - 10
			ODP: DailyObservationData{
				Active:   5, // 25 - 20
				Finished: 20,
			},
			PDP: DailySupervisionData{
				Active:   5, // 30 - 25
				Finished: 25,
			},
		},
		Cumulative: ProvinceCumulativeCases{
			Positive:  5000,
			Recovered: 4500,
			Deceased:  300,
			Active:    200, // 5000 - 4500 - 300
			ODP: ObservationData{
				Active:   50, // 800 - 750
				Finished: 750,
				Total:    800,
			},
			PDP: SupervisionData{
				Active:   20, // 600 - 580
				Finished: 580,
				Total:    600,
			},
		},
		Statistics: ProvinceCaseStatistics{
			Percentages: CasePercentages{
				Active:    4.0,  // (200 / 5000) * 100
				Recovered: 90.0, // (4500 / 5000) * 100
				Deceased:  6.0,  // (300 / 5000) * 100
			},
			ReproductionRate: &ReproductionRate{
				Value:      &[]float64{1.5}[0],
				UpperBound: &[]float64{1.8}[0],
				LowerBound: &[]float64{1.2}[0],
			},
		},
		// Province should be nil in this case
		Province: nil,
	}

	assert.Equal(t, expectedResult, result)
	assert.Nil(t, result.Province, "Province should be nil when using TransformToResponseWithoutProvince")
}

func TestProvinceCaseWithDate_TransformToResponseWithoutProvince(t *testing.T) {
	testDate := time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC)
	rt := 1.2

	provinceCaseWithDate := ProvinceCaseWithDate{
		ProvinceCase: ProvinceCase{
			ID:                                       1,
			Day:                                      200,
			ProvinceID:                               "ID-JT",
			Positive:                                 50,
			Recovered:                                40,
			Deceased:                                 2,
			PersonUnderObservation:                   10,
			FinishedPersonUnderObservation:           8,
			PersonUnderSupervision:                   12,
			FinishedPersonUnderSupervision:           10,
			CumulativePositive:                       3000,
			CumulativeRecovered:                      2700,
			CumulativeDeceased:                       200,
			CumulativePersonUnderObservation:         500,
			CumulativeFinishedPersonUnderObservation: 450,
			CumulativePersonUnderSupervision:         350,
			CumulativeFinishedPersonUnderSupervision: 320,
			Rt:                                       &rt,
			RtUpper:                                  nil,
			RtLower:                                  nil,
			Province: &Province{
				ID:   "ID-JT",
				Name: "Jawa Tengah",
			},
		},
		Date: testDate,
	}

	result := provinceCaseWithDate.TransformToResponseWithoutProvince()

	assert.Equal(t, int64(200), result.Day)
	assert.Equal(t, testDate, result.Date)
	assert.Equal(t, int64(50), result.Daily.Positive)
	assert.Equal(t, int64(8), result.Daily.Active) // 50 - 40 - 2
	assert.Equal(t, int64(3000), result.Cumulative.Positive)
	assert.Equal(t, int64(100), result.Cumulative.Active) // 3000 - 2700 - 200
	assert.Nil(t, result.Province, "Province should be nil when using TransformToResponseWithoutProvince")
}

func TestProvinceCaseResponse_JSONStructure(t *testing.T) {
	// This test verifies that the JSON structure matches the expected format
	testDate := time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC)
	rt := 1.5

	provinceCase := ProvinceCase{
		ID:                                       1,
		Day:                                      100,
		ProvinceID:                               "ID-JK",
		Positive:                                 150,
		Recovered:                                120,
		Deceased:                                 10,
		PersonUnderObservation:                   25,
		FinishedPersonUnderObservation:           20,
		PersonUnderSupervision:                   30,
		FinishedPersonUnderSupervision:           25,
		CumulativePositive:                       5000,
		CumulativeRecovered:                      4500,
		CumulativeDeceased:                       300,
		CumulativePersonUnderObservation:         800,
		CumulativeFinishedPersonUnderObservation: 750,
		CumulativePersonUnderSupervision:         600,
		CumulativeFinishedPersonUnderSupervision: 580,
		Rt:                                       &rt,
		RtUpper:                                  &rt,
		RtLower:                                  &rt,
		Province: &Province{
			ID:   "ID-JK",
			Name: "DKI Jakarta",
		},
	}

	result := provinceCase.TransformToResponse(testDate)

	// Verify the nested structure exists
	assert.NotNil(t, result.Daily)
	assert.NotNil(t, result.Cumulative)
	assert.NotNil(t, result.Statistics)
	assert.NotNil(t, result.Statistics.Percentages)
	assert.NotNil(t, result.Statistics.ReproductionRate)
	assert.NotNil(t, result.Province)

	// Verify key field names are in English
	assert.Equal(t, int64(100), result.Day) // "day"
	assert.Equal(t, testDate, result.Date)  // "date"
	// "daily" nested structure
	assert.Equal(t, int64(150), result.Daily.Positive)  // "positive"
	assert.Equal(t, int64(120), result.Daily.Recovered) // "recovered"
	assert.Equal(t, int64(10), result.Daily.Deceased)   // "deceased"
	assert.Equal(t, int64(20), result.Daily.Active)     // "active"
	// "cumulative" nested structure
	assert.Equal(t, int64(5000), result.Cumulative.Positive)  // "positive"
	assert.Equal(t, int64(4500), result.Cumulative.Recovered) // "recovered"
	assert.Equal(t, int64(300), result.Cumulative.Deceased)   // "deceased"
	assert.Equal(t, int64(200), result.Cumulative.Active)     // "active"
	// "statistics" nested structure with "percentages" and "reproduction_rate"
	assert.True(t, result.Statistics.Percentages.Active > 0)
	assert.True(t, result.Statistics.Percentages.Recovered > 0)
	assert.True(t, result.Statistics.Percentages.Deceased > 0)
}
