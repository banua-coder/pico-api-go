package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePaginationMeta(t *testing.T) {
	tests := []struct {
		name          string
		limit         int
		offset        int
		total         int
		expectedMeta  PaginationMeta
	}{
		{
			name:   "First page with results",
			limit:  50,
			offset: 0,
			total:  200,
			expectedMeta: PaginationMeta{
				Limit:      50,
				Offset:     0,
				Total:      200,
				TotalPages: 4,
				Page:       1,
				HasNext:    true,
				HasPrev:    false,
			},
		},
		{
			name:   "Middle page",
			limit:  50,
			offset: 50,
			total:  200,
			expectedMeta: PaginationMeta{
				Limit:      50,
				Offset:     50,
				Total:      200,
				TotalPages: 4,
				Page:       2,
				HasNext:    true,
				HasPrev:    true,
			},
		},
		{
			name:   "Last page",
			limit:  50,
			offset: 150,
			total:  200,
			expectedMeta: PaginationMeta{
				Limit:      50,
				Offset:     150,
				Total:      200,
				TotalPages: 4,
				Page:       4,
				HasNext:    false,
				HasPrev:    true,
			},
		},
		{
			name:   "Single page with all data",
			limit:  100,
			offset: 0,
			total:  50,
			expectedMeta: PaginationMeta{
				Limit:      100,
				Offset:     0,
				Total:      50,
				TotalPages: 1,
				Page:       1,
				HasNext:    false,
				HasPrev:    false,
			},
		},
		{
			name:   "Exact fit last page",
			limit:  25,
			offset: 75,
			total:  100,
			expectedMeta: PaginationMeta{
				Limit:      25,
				Offset:     75,
				Total:      100,
				TotalPages: 4,
				Page:       4,
				HasNext:    false,
				HasPrev:    true,
			},
		},
		{
			name:   "Empty result set",
			limit:  50,
			offset: 0,
			total:  0,
			expectedMeta: PaginationMeta{
				Limit:      50,
				Offset:     0,
				Total:      0,
				TotalPages: 0,
				Page:       1,
				HasNext:    false,
				HasPrev:    false,
			},
		},
		{
			name:   "Large offset beyond total",
			limit:  50,
			offset: 500,
			total:  100,
			expectedMeta: PaginationMeta{
				Limit:      50,
				Offset:     500,
				Total:      100,
				TotalPages: 2,
				Page:       11,
				HasNext:    false,
				HasPrev:    true,
			},
		},
		{
			name:   "Partial last page",
			limit:  30,
			offset: 90,
			total:  100,
			expectedMeta: PaginationMeta{
				Limit:      30,
				Offset:     90,
				Total:      100,
				TotalPages: 4,
				Page:       4,
				HasNext:    false,
				HasPrev:    true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			meta := CalculatePaginationMeta(tt.limit, tt.offset, tt.total)
			assert.Equal(t, tt.expectedMeta, meta)
		})
	}
}

func TestPaginationMetaCalculations(t *testing.T) {
	t.Run("Total pages calculation for different scenarios", func(t *testing.T) {
		// Test ceiling division for total pages
		assert.Equal(t, 4, CalculatePaginationMeta(33, 0, 100).TotalPages) // 100/33 = 3.03 -> 4 pages
		assert.Equal(t, 2, CalculatePaginationMeta(50, 0, 100).TotalPages) // 100/50 = 2 -> 2 pages  
		assert.Equal(t, 3, CalculatePaginationMeta(33, 0, 99).TotalPages)  // 99/33 = 3 -> 3 pages
	})

	t.Run("Page number calculation", func(t *testing.T) {
		assert.Equal(t, 1, CalculatePaginationMeta(50, 0, 200).Page)   // offset 0 = page 1
		assert.Equal(t, 2, CalculatePaginationMeta(50, 50, 200).Page)  // offset 50 = page 2
		assert.Equal(t, 3, CalculatePaginationMeta(50, 100, 200).Page) // offset 100 = page 3
	})

	t.Run("Has next and previous flags", func(t *testing.T) {
		// First page
		meta := CalculatePaginationMeta(50, 0, 200)
		assert.False(t, meta.HasPrev)
		assert.True(t, meta.HasNext)

		// Middle page
		meta = CalculatePaginationMeta(50, 50, 200)
		assert.True(t, meta.HasPrev)
		assert.True(t, meta.HasNext)

		// Last page
		meta = CalculatePaginationMeta(50, 150, 200)
		assert.True(t, meta.HasPrev)
		assert.False(t, meta.HasNext)

		// Single page
		meta = CalculatePaginationMeta(100, 0, 50)
		assert.False(t, meta.HasPrev)
		assert.False(t, meta.HasNext)
	})
}

func TestPaginatedResponse(t *testing.T) {
	t.Run("PaginatedResponse structure", func(t *testing.T) {
		testData := []string{"item1", "item2", "item3"}
		pagination := PaginationMeta{
			Limit:      10,
			Offset:     0,
			Total:      3,
			TotalPages: 1,
			Page:       1,
			HasNext:    false,
			HasPrev:    false,
		}

		response := PaginatedResponse{
			Data:       testData,
			Pagination: pagination,
		}

		assert.Equal(t, testData, response.Data)
		assert.Equal(t, pagination, response.Pagination)
	})
}
