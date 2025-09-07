package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNationalCase_Structure(t *testing.T) {
	now := time.Now()
	rt := 1.2
	rtUpper := 1.5
	rtLower := 0.9

	nationalCase := NationalCase{
		ID:                  1,
		Day:                 1,
		Date:                now,
		Positive:            100,
		Recovered:           80,
		Deceased:            5,
		CumulativePositive:  1000,
		CumulativeRecovered: 800,
		CumulativeDeceased:  50,
		Rt:                  &rt,
		RtUpper:             &rtUpper,
		RtLower:             &rtLower,
	}

	assert.Equal(t, int64(1), nationalCase.ID)
	assert.Equal(t, int64(1), nationalCase.Day)
	assert.Equal(t, now, nationalCase.Date)
	assert.Equal(t, int64(100), nationalCase.Positive)
	assert.Equal(t, int64(80), nationalCase.Recovered)
	assert.Equal(t, int64(5), nationalCase.Deceased)
	assert.Equal(t, int64(1000), nationalCase.CumulativePositive)
	assert.Equal(t, int64(800), nationalCase.CumulativeRecovered)
	assert.Equal(t, int64(50), nationalCase.CumulativeDeceased)
	assert.NotNil(t, nationalCase.Rt)
	assert.Equal(t, 1.2, *nationalCase.Rt)
	assert.NotNil(t, nationalCase.RtUpper)
	assert.Equal(t, 1.5, *nationalCase.RtUpper)
	assert.NotNil(t, nationalCase.RtLower)
	assert.Equal(t, 0.9, *nationalCase.RtLower)
}

func TestNationalCase_NullableFields(t *testing.T) {
	nationalCase := NationalCase{
		ID:                  1,
		Day:                 1,
		Date:                time.Now(),
		Positive:            100,
		Recovered:           80,
		Deceased:            5,
		CumulativePositive:  1000,
		CumulativeRecovered: 800,
		CumulativeDeceased:  50,
		Rt:                  nil,
		RtUpper:             nil,
		RtLower:             nil,
	}

	assert.Nil(t, nationalCase.Rt)
	assert.Nil(t, nationalCase.RtUpper)
	assert.Nil(t, nationalCase.RtLower)
}

func TestNullFloat64_Scan(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{}
		expected    NullFloat64
		expectError bool
	}{
		{
			name:     "nil value",
			input:    nil,
			expected: NullFloat64{Float64: 0, Valid: false},
		},
		{
			name:     "float64 value",
			input:    1.5,
			expected: NullFloat64{Float64: 1.5, Valid: true},
		},
		{
			name:     "empty byte slice",
			input:    []byte(""),
			expected: NullFloat64{Float64: 0, Valid: false},
		},
		{
			name:     "byte slice with valid float",
			input:    []byte("2.5"),
			expected: NullFloat64{Float64: 2.5, Valid: true},
		},
		{
			name:        "invalid type",
			input:       "invalid",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nf NullFloat64
			err := nf.Scan(tt.input)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.Float64, nf.Float64)
				assert.Equal(t, tt.expected.Valid, nf.Valid)
			}
		})
	}
}

func TestNullFloat64_Value(t *testing.T) {
	tests := []struct {
		name     string
		input    NullFloat64
		expected interface{}
	}{
		{
			name:     "valid float",
			input:    NullFloat64{Float64: 1.5, Valid: true},
			expected: 1.5,
		},
		{
			name:     "invalid float",
			input:    NullFloat64{Float64: 0, Valid: false},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := tt.input.Value()
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}