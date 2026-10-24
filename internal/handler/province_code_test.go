package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// MySQL coerces '072', '72abc', '+72' and '72.0' to 72, so non-canonical codes must be
// rejected before they reach the service (and its cache keys).
var nonCanonicalCodes = []string{"072", "72abc", "+72", "72.0", "-1", "0", "abc"}

func TestCovidHandler_GetProvinceByID_RejectsNonCanonicalCode(t *testing.T) {
	for _, code := range nonCanonicalCodes {
		svc := new(MockCovidService)
		router := mux.NewRouter()
		router.HandleFunc("/api/v1/provinces/{code}", NewCovidHandler(svc, nil).GetProvinceByID)

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, "/api/v1/provinces/"+code, nil))

		assert.Equal(t, http.StatusNotFound, rr.Code, code)
		svc.AssertNotCalled(t, "GetProvinceByID")
	}
}

func TestCovidHandler_GetProvinceCases_RejectsNonCanonicalCode(t *testing.T) {
	for _, code := range nonCanonicalCodes {
		svc := new(MockCovidService)
		rr := serveProvinceCases(NewCovidHandler(svc, nil), "/api/v1/provinces/"+code+"/cases")

		assert.Equal(t, http.StatusNotFound, rr.Code, code)
		svc.AssertNotCalled(t, "GetProvinceByID")
	}
}
