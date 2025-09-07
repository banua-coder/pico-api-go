package utils

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIntQueryParam(t *testing.T) {
	tests := []struct {
		name         string
		queryValue   string
		defaultValue int
		expected     int
	}{
		{
			name:         "Valid integer",
			queryValue:   "100",
			defaultValue: 50,
			expected:     100,
		},
		{
			name:         "Empty value uses default",
			queryValue:   "",
			defaultValue: 50,
			expected:     50,
		},
		{
			name:         "Invalid integer uses default",
			queryValue:   "not-a-number",
			defaultValue: 50,
			expected:     50,
		},
		{
			name:         "Zero value",
			queryValue:   "0",
			defaultValue: 50,
			expected:     0,
		},
		{
			name:         "Negative value",
			queryValue:   "-10",
			defaultValue: 50,
			expected:     -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				URL: &url.URL{
					RawQuery: url.Values{"test_param": []string{tt.queryValue}}.Encode(),
				},
			}

			result := ParseIntQueryParam(req, "test_param", tt.defaultValue)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseBoolQueryParam(t *testing.T) {
	tests := []struct {
		name       string
		queryValue string
		expected   bool
	}{
		{
			name:       "true value",
			queryValue: "true",
			expected:   true,
		},
		{
			name:       "false value",
			queryValue: "false",
			expected:   false,
		},
		{
			name:       "empty value",
			queryValue: "",
			expected:   false,
		},
		{
			name:       "non-boolean value",
			queryValue: "yes",
			expected:   false,
		},
		{
			name:       "1 is not true",
			queryValue: "1",
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				URL: &url.URL{
					RawQuery: url.Values{"test_param": []string{tt.queryValue}}.Encode(),
				},
			}

			result := ParseBoolQueryParam(req, "test_param")
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestParseStringArrayQueryParam(t *testing.T) {
	tests := []struct {
		name       string
		queryValue string
		expected   []string
	}{
		{
			name:       "Single value",
			queryValue: "value1",
			expected:   []string{"value1"},
		},
		{
			name:       "Multiple values",
			queryValue: "value1,value2,value3",
			expected:   []string{"value1", "value2", "value3"},
		},
		{
			name:       "Values with spaces",
			queryValue: "value1, value2 , value3",
			expected:   []string{"value1", "value2", "value3"},
		},
		{
			name:       "Empty string returns nil",
			queryValue: "",
			expected:   nil,
		},
		{
			name:       "Only commas and spaces",
			queryValue: " , , ",
			expected:   nil,
		},
		{
			name:       "Mixed empty and valid values",
			queryValue: "value1,,value2, ,value3",
			expected:   []string{"value1", "value2", "value3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &http.Request{
				URL: &url.URL{
					RawQuery: url.Values{"test_param": []string{tt.queryValue}}.Encode(),
				},
			}

			result := ParseStringArrayQueryParam(req, "test_param")
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidatePaginationParams(t *testing.T) {
	tests := []struct {
		name           string
		inputLimit     int
		inputOffset    int
		expectedLimit  int
		expectedOffset int
	}{
		{
			name:           "Valid parameters",
			inputLimit:     100,
			inputOffset:    50,
			expectedLimit:  100,
			expectedOffset: 50,
		},
		{
			name:           "Zero limit uses default",
			inputLimit:     0,
			inputOffset:    10,
			expectedLimit:  50,
			expectedOffset: 10,
		},
		{
			name:           "Negative limit uses default",
			inputLimit:     -10,
			inputOffset:    10,
			expectedLimit:  50,
			expectedOffset: 10,
		},
		{
			name:           "Limit exceeds max",
			inputLimit:     2000,
			inputOffset:    10,
			expectedLimit:  1000,
			expectedOffset: 10,
		},
		{
			name:           "Negative offset uses zero",
			inputLimit:     100,
			inputOffset:    -10,
			expectedLimit:  100,
			expectedOffset: 0,
		},
		{
			name:           "Both invalid values",
			inputLimit:     -5,
			inputOffset:    -10,
			expectedLimit:  50,
			expectedOffset: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			limit, offset := ValidatePaginationParams(tt.inputLimit, tt.inputOffset)
			assert.Equal(t, tt.expectedLimit, limit)
			assert.Equal(t, tt.expectedOffset, offset)
		})
	}
}