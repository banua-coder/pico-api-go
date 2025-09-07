package handler

import (
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/gorilla/mux"
)

func SetupRoutes(covidService service.CovidService) *mux.Router {
	router := mux.NewRouter()

	covidHandler := NewCovidHandler(covidService)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/health", covidHandler.HealthCheck).Methods("GET")
	api.HandleFunc("/national", covidHandler.GetNationalCases).Methods("GET")
	api.HandleFunc("/national/latest", covidHandler.GetLatestNationalCase).Methods("GET")
	api.HandleFunc("/provinces", covidHandler.GetProvinces).Methods("GET")
	api.HandleFunc("/provinces/cases", covidHandler.GetProvinceCases).Methods("GET")
	api.HandleFunc("/provinces/{provinceId}/cases", covidHandler.GetProvinceCases).Methods("GET")

	return router
}