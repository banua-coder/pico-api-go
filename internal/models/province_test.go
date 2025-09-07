package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvince_Structure(t *testing.T) {
	province := Province{
		ID:   "11",
		Name: "Aceh",
	}

	assert.Equal(t, "11", province.ID)
	assert.Equal(t, "Aceh", province.Name)
}

func TestProvince_IndonesianCodes(t *testing.T) {
	testCases := []struct {
		name         string
		provinceCode string
		provinceName string
	}{
		{"Aceh", "11", "Aceh"},
		{"Sulawesi Tengah", "72", "Sulawesi Tengah"},
		{"DKI Jakarta", "31", "DKI Jakarta"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			province := Province{
				ID:   tc.provinceCode,
				Name: tc.provinceName,
			}

			assert.Equal(t, tc.provinceCode, province.ID)
			assert.Equal(t, tc.provinceName, province.Name)
		})
	}
}