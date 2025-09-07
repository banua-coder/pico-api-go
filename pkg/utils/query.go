package utils

import (
	"net/http"
	"strconv"
	"strings"
)

// ParseIntQueryParam parses an integer query parameter with a default value
func ParseIntQueryParam(r *http.Request, key string, defaultValue int) int {
	valueStr := r.URL.Query().Get(key)
	if valueStr == "" {
		return defaultValue
	}
	
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	
	return value
}

// ParseBoolQueryParam parses a boolean query parameter
func ParseBoolQueryParam(r *http.Request, key string) bool {
	return r.URL.Query().Get(key) == "true"
}

// ParseStringArrayQueryParam parses a comma-separated string parameter into array
func ParseStringArrayQueryParam(r *http.Request, key string) []string {
	valueStr := r.URL.Query().Get(key)
	if valueStr == "" {
		return nil
	}
	
	values := strings.Split(valueStr, ",")
	var result []string
	for _, v := range values {
		v = strings.TrimSpace(v)
		if v != "" {
			result = append(result, v)
		}
	}
	
	return result
}

// ValidatePaginationParams validates and adjusts pagination parameters
func ValidatePaginationParams(limit, offset int) (int, int) {
	// Validate limit
	if limit <= 0 {
		limit = 50 // Default limit
	} else if limit > 1000 {
		limit = 1000 // Max limit
	}
	
	// Validate offset
	if offset < 0 {
		offset = 0
	}
	
	return limit, offset
}