package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/database"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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

	// Swagger documentation
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods("GET")
	
	// Redirect root to swagger docs for convenience  
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	}).Methods("GET")

	return router
}