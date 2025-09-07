package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/config"
	"github.com/banua-coder/pico-api-go/internal/handler"
	"github.com/banua-coder/pico-api-go/internal/middleware"
	"github.com/banua-coder/pico-api-go/internal/repository"
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/database"
)

func main() {
	cfg := config.Load()

	db, err := database.NewMySQLConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	nationalCaseRepo := repository.NewNationalCaseRepository(db)
	provinceRepo := repository.NewProvinceRepository(db)
	provinceCaseRepo := repository.NewProvinceCaseRepository(db)

	covidService := service.NewCovidService(nationalCaseRepo, provinceRepo, provinceCaseRepo)

	router := handler.SetupRoutes(covidService, db)

	router.Use(middleware.Recovery)
	router.Use(middleware.Logging)
	router.Use(middleware.CORS)

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", address)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}