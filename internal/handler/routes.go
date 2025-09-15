package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/database"
	"github.com/gorilla/mux"
	// httpSwagger "github.com/swaggo/http-swagger" // Disabled for minimal production build
)

func SetupRoutes(covidService service.CovidService, db *database.DB, enableSwagger bool) *mux.Router {
	router := mux.NewRouter()

	covidHandler := NewCovidHandler(covidService, db)

	api := router.PathPrefix("/api/v1").Subrouter()

	// API index endpoint
	api.HandleFunc("", covidHandler.GetAPIIndex).Methods("GET", "OPTIONS")
	api.HandleFunc("/", covidHandler.GetAPIIndex).Methods("GET", "OPTIONS")

	// Main endpoints
	api.HandleFunc("/health", covidHandler.HealthCheck).Methods("GET", "OPTIONS")
	api.HandleFunc("/national", covidHandler.GetNationalCases).Methods("GET", "OPTIONS")
	api.HandleFunc("/national/latest", covidHandler.GetLatestNationalCase).Methods("GET", "OPTIONS")
	api.HandleFunc("/provinces", covidHandler.GetProvinces).Methods("GET", "OPTIONS")
	api.HandleFunc("/provinces/cases", covidHandler.GetProvinceCases).Methods("GET", "OPTIONS")
	api.HandleFunc("/provinces/{provinceId}/cases", covidHandler.GetProvinceCases).Methods("GET", "OPTIONS")

	// Conditionally add Swagger documentation based on environment
	if enableSwagger {
		// Development: Add Swagger documentation
		// Note: httpSwagger import is disabled for minimal production builds
		router.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Swagger UI not available in minimal build - see static documentation site", http.StatusNotFound)
		}).Methods("GET")

		// Redirect root to API index (Swagger disabled for minimal build)
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/api/v1", http.StatusFound)
		}).Methods("GET")
	} else {
		// Production: Redirect root to API index
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/api/v1", http.StatusFound)
		}).Methods("GET")
	}

	return router
}
