package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProvinceCase_Structure(t *testing.T) {
	rt := 1.1
	rtUpper := 1.3
	rtLower := 0.8

	provinceCase := ProvinceCase{
		ID:                                               1,
		Day:                                              1,
		ProvinceID:                                       "11",
		Positive:                                         50,
		Recovered:                                        40,
		Deceased:                                         2,
		PersonUnderObservation:                           10,
		FinishedPersonUnderObservation:                   8,
		PersonUnderSupervision:                           5,
		FinishedPersonUnderSupervision:                   3,
		CumulativePositive:                               500,
		CumulativeRecovered:                              400,
		CumulativeDeceased:                               20,
		CumulativePersonUnderObservation:                 100,
		CumulativeFinishedPersonUnderObservation:         80,
		CumulativePersonUnderSupervision:                 50,
		CumulativeFinishedPersonUnderSupervision:         30,
		Rt:                                               &rt,
		RtUpper:                                          &rtUpper,
		RtLower:                                          &rtLower,
		Province:                                         &Province{ID: "11", Name: "Aceh"},
	}

	assert.Equal(t, int64(1), provinceCase.ID)
	assert.Equal(t, int64(1), provinceCase.Day)
	assert.Equal(t, "11", provinceCase.ProvinceID)
	assert.Equal(t, int64(50), provinceCase.Positive)
	assert.Equal(t, int64(40), provinceCase.Recovered)
	assert.Equal(t, int64(2), provinceCase.Deceased)
	assert.Equal(t, int64(10), provinceCase.PersonUnderObservation)
	assert.Equal(t, int64(8), provinceCase.FinishedPersonUnderObservation)
	assert.Equal(t, int64(5), provinceCase.PersonUnderSupervision)
	assert.Equal(t, int64(3), provinceCase.FinishedPersonUnderSupervision)
	assert.Equal(t, int64(500), provinceCase.CumulativePositive)
	assert.Equal(t, int64(400), provinceCase.CumulativeRecovered)
	assert.Equal(t, int64(20), provinceCase.CumulativeDeceased)
	assert.NotNil(t, provinceCase.Rt)
	assert.Equal(t, 1.1, *provinceCase.Rt)
	assert.NotNil(t, provinceCase.Province)
	assert.Equal(t, "11", provinceCase.Province.ID)
	assert.Equal(t, "Aceh", provinceCase.Province.Name)
}

func TestProvinceCase_WithoutProvince(t *testing.T) {
	provinceCase := ProvinceCase{
		ID:         1,
		Day:        1,
		ProvinceID: "11",
		Positive:   50,
		Province:   nil,
	}

	assert.Equal(t, "11", provinceCase.ProvinceID)
	assert.Nil(t, provinceCase.Province)
}

func TestProvinceCaseWithDate_Structure(t *testing.T) {
	now := time.Now()
	rt := 1.1

	provinceCaseWithDate := ProvinceCaseWithDate{
		ProvinceCase: ProvinceCase{
			ID:         1,
			Day:        1,
			ProvinceID: "11",
			Positive:   50,
			Rt:         &rt,
		},
		Date: now,
	}

	assert.Equal(t, int64(1), provinceCaseWithDate.ID)
	assert.Equal(t, now, provinceCaseWithDate.Date)
	assert.Equal(t, "11", provinceCaseWithDate.ProvinceID)
	assert.Equal(t, int64(50), provinceCaseWithDate.Positive)
	assert.NotNil(t, provinceCaseWithDate.Rt)
	assert.Equal(t, 1.1, *provinceCaseWithDate.Rt)
}
