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

// SortParams represents sorting parameters
type SortParams struct {
	Field string `json:"field"`
	Order string `json:"order"` // "asc" or "desc"
}

// ParseSortParam parses sort parameter from query string
// Format: ?sort=field:order or ?sort=field (defaults to asc)
// Example: ?sort=date:desc or ?sort=date
func ParseSortParam(r *http.Request, defaultField string) SortParams {
	sortParam := r.URL.Query().Get("sort")

	// Default sorting by date ascending
	if sortParam == "" {
		return SortParams{
			Field: defaultField,
			Order: "asc",
		}
	}

	parts := strings.Split(sortParam, ":")
	field := strings.TrimSpace(parts[0])
	order := "asc" // default order

	if len(parts) > 1 {
		orderParam := strings.ToLower(strings.TrimSpace(parts[1]))
		if orderParam == "desc" || orderParam == "asc" {
			order = orderParam
		}
	}

	// Validate field name (prevent SQL injection)
	if !IsValidSortField(field) {
		field = defaultField
	}

	return SortParams{
		Field: field,
		Order: order,
	}
}

// IsValidSortField validates if the field name is allowed for sorting
func IsValidSortField(field string) bool {
	allowedFields := map[string]bool{
		"date":          true,
		"day":           true,
		"positive":      true,
		"recovered":     true,
		"deceased":      true,
		"active":        true,
		"province_id":   true,
		"province_name": true,
		"created_at":    true,
		"updated_at":    true,
	}

	return allowedFields[field]
}

// GetSQLOrderClause generates SQL ORDER BY clause from sort parameters
func (s SortParams) GetSQLOrderClause() string {
	// Map API field names to database column names
	fieldMapping := map[string]string{
		"date":          "date",
		"day":           "day",
		"positive":      "positive",
		"recovered":     "recovered",
		"deceased":      "deceased",
		"active":        "active",
		"province_id":   "province_id",
		"province_name": "province_name",
		"created_at":    "created_at",
		"updated_at":    "updated_at",
	}

	dbField, exists := fieldMapping[s.Field]
	if !exists {
		dbField = "date" // fallback to date
	}

	order := strings.ToUpper(s.Order)
	if order != "DESC" {
		order = "ASC" // default to ASC
	}

	return dbField + " " + order
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

// ParsePaginationParams parses pagination parameters from request
// Supports both offset-based and page-based pagination
func ParsePaginationParams(r *http.Request) (limit, offset int) {
	// Parse limit (records per page)
	limit = ParseIntQueryParam(r, "limit", 50)

	// Check if page parameter is provided
	page := ParseIntQueryParam(r, "page", 0)
	if page > 0 {
		// Page-based pagination (page starts from 1)
		offset = (page - 1) * limit
	} else {
		// Offset-based pagination (fallback)
		offset = ParseIntQueryParam(r, "offset", 0)
	}

	// Validate and adjust parameters
	return ValidatePaginationParams(limit, offset)
}
