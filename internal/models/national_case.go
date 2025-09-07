package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type NationalCase struct {
	ID                     int64      `json:"id" db:"id"`
	Day                    int64      `json:"day" db:"day"`
	Date                   time.Time  `json:"date" db:"date"`
	Positive               int64      `json:"positive" db:"positive"`
	Recovered              int64      `json:"recovered" db:"recovered"`
	Deceased               int64      `json:"deceased" db:"deceased"`
	CumulativePositive     int64      `json:"cumulative_positive" db:"cumulative_positive"`
	CumulativeRecovered    int64      `json:"cumulative_recovered" db:"cumulative_recovered"`
	CumulativeDeceased     int64      `json:"cumulative_deceased" db:"cumulative_deceased"`
	Rt                     *float64   `json:"rt" db:"rt"`
	RtUpper                *float64   `json:"rt_upper" db:"rt_upper"`
	RtLower                *float64   `json:"rt_lower" db:"rt_lower"`
}

type NullFloat64 struct {
	Float64 float64
	Valid   bool
}

func (nf *NullFloat64) Scan(value interface{}) error {
	if value == nil {
		nf.Float64, nf.Valid = 0, false
		return nil
	}
	switch v := value.(type) {
	case float64:
		nf.Float64, nf.Valid = v, true
	case []byte:
		if len(v) == 0 {
			nf.Float64, nf.Valid = 0, false
		} else {
			var f float64
			if _, err := fmt.Sscanf(string(v), "%f", &f); err != nil {
				return err
			}
			nf.Float64, nf.Valid = f, true
		}
	default:
		return fmt.Errorf("cannot scan %T into NullFloat64", value)
	}
	return nil
}

func (nf NullFloat64) Value() (driver.Value, error) {
	if !nf.Valid {
		return nil, nil
	}
	return nf.Float64, nil
}