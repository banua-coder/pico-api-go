package models

import "time"

// TestType represents a COVID-19 testing method
type TestType struct {
	ID            int64  `json:"id"`
	Key           string `json:"key"`
	Name          string `json:"name"`
	Sample        string `json:"sample"`
	Duration      string `json:"duration"`
	IsRecommended bool   `json:"is_recommended"`
}

// ProvinceTest represents province-level COVID-19 test results
type ProvinceTest struct {
	ID         int64     `json:"id"`
	TestTypeID int64     `json:"test_type_id"`
	Day        int64     `json:"day"`
	ProvinceID int       `json:"province_id"`
	DateFrom   time.Time `json:"date_from"`
	Process    int       `json:"process"`
	Invalid    int       `json:"invalid"`
	Positive   int       `json:"positive"`
	Negative   int       `json:"negative"`

	// Joined field
	TestType *TestType `json:"test_type,omitempty"`
}
