package models

// PaginationMeta contains metadata for paginated responses
type PaginationMeta struct {
	Limit      int  `json:"limit"`
	Offset     int  `json:"offset"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	Page       int  `json:"page"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

// PaginatedResponse wraps data with pagination metadata
type PaginatedResponse struct {
	Data       interface{}    `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// CalculatePaginationMeta calculates pagination metadata
func CalculatePaginationMeta(limit, offset, total int) PaginationMeta {
	totalPages := (total + limit - 1) / limit // Ceiling division
	page := (offset / limit) + 1

	return PaginationMeta{
		Limit:      limit,
		Offset:     offset,
		Total:      total,
		TotalPages: totalPages,
		Page:       page,
		HasNext:    offset+limit < total,
		HasPrev:    offset > 0,
	}
}
