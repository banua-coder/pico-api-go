package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/database"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Services holds all service dependencies for route setup
type Services struct {
	CovidService    service.CovidService
	RegencyService  *service.RegencyService
	HospitalService *service.HospitalService
	TaskForceService *service.TaskForceService
}

func SetupRoutes(svc Services, db *database.DB, enableSwagger bool) *mux.Router {
	router := mux.NewRouter()

	covidHandler := NewCovidHandler(svc.CovidService, db)

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

	// Regency endpoints
	if svc.RegencyService != nil {
		regencyHandler := NewRegencyHandler(svc.RegencyService)
		api.HandleFunc("/regencies", regencyHandler.GetRegencies).Methods("GET", "OPTIONS")
		api.HandleFunc("/regencies/{code}", regencyHandler.GetRegencyByID).Methods("GET", "OPTIONS")
		api.HandleFunc("/regencies/{code}/cases", regencyHandler.GetRegencyCases).Methods("GET", "OPTIONS")
	}

	// Hospital endpoints
	if svc.HospitalService != nil {
		hospitalHandler := NewHospitalHandler(svc.HospitalService)
		api.HandleFunc("/hospitals", hospitalHandler.GetHospitals).Methods("GET", "OPTIONS")
		api.HandleFunc("/hospitals/{code}", hospitalHandler.GetHospitalByCode).Methods("GET", "OPTIONS")
	}

	// Task force endpoints
	if svc.TaskForceService != nil {
		taskForceHandler := NewTaskForceHandler(svc.TaskForceService)
		api.HandleFunc("/task-forces", taskForceHandler.GetTaskForces).Methods("GET", "OPTIONS")
	}

	// Conditionally add Swagger documentation based on environment
	if enableSwagger {
		router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
		}).Methods("GET")
	} else {
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/api/v1", http.StatusFound)
		}).Methods("GET")
	}

	return router
}
