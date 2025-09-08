package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/database"
	"github.com/gorilla/mux"
)

func SetupRoutes(covidService service.CovidService, db *database.DB) *mux.Router {
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

	// API specification endpoint
	api.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "./docs/swagger.json")
	}).Methods("GET", "OPTIONS")

	// Redirect /docs to static swagger UI
	router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	}).Methods("GET")

	return router
}
