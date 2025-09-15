// Package main provides the entry point for the Sulawesi Tengah COVID-19 Data API
//
//	@title			Sulawesi Tengah COVID-19 Data API
//	@version		2.4.0
//	@description	A comprehensive REST API for COVID-19 data in Sulawesi Tengah (Central Sulawesi), with additional national and provincial data for context. Features enhanced ODP/PDP grouping, hybrid pagination, and rate limiting protection. Rate limiting:		2.4.0 requests per minute per IP address by default, with appropriate HTTP headers for client guidance.
//	@termsOfService	http://swagger.io/terms/
//
//	@contact.name	API Support
//	@contact.url	https://github.com/banua-coder/pico-api-go
//	@contact.email	support@banuacoder.com
//
//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT
//
//	@host		pico-api.banuacoder.com
//	@BasePath	/api/v1
//
//	@schemes	https http
//
//	@tag.name			health
//	@tag.description	Health check operations
//
//	@tag.name			national
//	@tag.description	National COVID-19 case operations (for context)
//
//	@tag.name			provinces
//	@tag.description	Province information and COVID-19 case operations (focus on Sulawesi Tengah)
//
//	@tag.name			province-cases
//	@tag.description	Province-level COVID-19 case data with pagination support
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/banua-coder/pico-api-go/internal/config"
	"github.com/banua-coder/pico-api-go/internal/handler"
	"github.com/banua-coder/pico-api-go/internal/middleware"
	"github.com/banua-coder/pico-api-go/internal/repository"
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/database"
	// Import docs for Swagger (disabled for production deployment)
	// _ "github.com/banua-coder/pico-api-go/docs"
)

func main() {
	cfg := config.Load()

	db, err := database.NewMySQLConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	log.Println("Database connected successfully")

	nationalCaseRepo := repository.NewNationalCaseRepository(db)
	provinceRepo := repository.NewProvinceRepository(db)
	provinceCaseRepo := repository.NewProvinceCaseRepository(db)

	covidService := service.NewCovidService(nationalCaseRepo, provinceRepo, provinceCaseRepo)

	// Check if we should enable Swagger (disabled in production)
	enableSwagger := os.Getenv("ENV") != "production"
	router := handler.SetupRoutes(covidService, db, enableSwagger)

	router.Use(middleware.Recovery)
	router.Use(middleware.Logging)
	router.Use(middleware.RateLimit(cfg.RateLimit))
	router.Use(middleware.CORS)

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Server starting on %s", address)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
