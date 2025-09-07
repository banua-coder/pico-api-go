package handler

import (
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/database"
	"github.com/gorilla/mux"
)

func SetupRoutes(covidService service.CovidService, db *database.DB) *mux.Router {
	router := mux.NewRouter()

	covidHandler := NewCovidHandler(covidService, db)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/health", covidHandler.HealthCheck).Methods("GET", "OPTIONS")
	api.HandleFunc("/national", covidHandler.GetNationalCases).Methods("GET", "OPTIONS")
	api.HandleFunc("/national/latest", covidHandler.GetLatestNationalCase).Methods("GET", "OPTIONS")
	api.HandleFunc("/provinces", covidHandler.GetProvinces).Methods("GET", "OPTIONS")
	api.HandleFunc("/provinces/cases", covidHandler.GetProvinceCases).Methods("GET", "OPTIONS")
	api.HandleFunc("/provinces/{provinceId}/cases", covidHandler.GetProvinceCases).Methods("GET", "OPTIONS")

	return router
}