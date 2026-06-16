// Package main provides the entry point for the Sulawesi Tengah COVID-19 Data API
//
//	@title			Sulawesi Tengah COVID-19 Data API
//	@version		2.9.0
//	@description	A comprehensive REST API for COVID-19 data in Sulawesi Tengah (Central Sulawesi), with additional national and provincial data for context. Features enhanced ODP/PDP grouping, hybrid pagination, and rate limiting protection. Rate limiting: 100 requests per minute per IP address by default, with appropriate HTTP headers for client guidance.
//	@termsOfService	http://swagger.io/terms/
//
//	@contact.name	API Support
//	@contact.url	https://github.com/banua-coder/pico-api-go
//	@contact.email	support@banuacoder.com
//
//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT
//
//	@host		pico-api-go.banuacoder.com
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
	"time"

	"github.com/banua-coder/pico-api-go/docs"
	"github.com/banua-coder/pico-api-go/internal/config"
	"github.com/banua-coder/pico-api-go/internal/handler"
	"github.com/banua-coder/pico-api-go/internal/middleware"
	"github.com/banua-coder/pico-api-go/internal/repository"
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/cache"
	"github.com/banua-coder/pico-api-go/pkg/database"
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

	// Initialize cache — use Redis-backed dual-layer if REDIS_ADDR is set, otherwise in-memory only
	var c *cache.Cache
	var cacheInvalidator service.CacheInvalidator

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr != "" {
		rac, err := cache.NewRedisAwareCache(time.Hour, cache.RedisOptions{
			Addr:     redisAddr,
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       0,
		})
		if err != nil {
			log.Printf("Redis unavailable (%v), falling back to in-memory cache only", err)
			c = cache.New(time.Hour)
			cacheInvalidator = c
		} else {
			log.Printf("Redis connected: %s (dual-layer cache active)", redisAddr)
			c = rac.Unwrap()
			cacheInvalidator = rac
		}
	} else {
		c = cache.New(time.Hour)
		cacheInvalidator = c
	}
	c.StartCleanup(5 * time.Minute)

	covidService := service.NewCachedCovidService(
		service.NewCovidService(nationalCaseRepo, provinceRepo, provinceCaseRepo),
		c,
	)

	// New repositories and services for migrated Lumen endpoints
	regencyRepo := repository.NewRegencyRepository(db)
	regencyCaseRepo := repository.NewRegencyCaseRepository(db)
	hospitalRepo := repository.NewHospitalRepository(db)
	taskForceRepo := repository.NewTaskForceRepository(db)

	regencyService := service.NewCachedRegencyService(
		service.NewRegencyService(regencyRepo, regencyCaseRepo),
		c,
	)
	hospitalService := service.NewHospitalService(hospitalRepo)
	taskForceService := service.NewTaskForceService(taskForceRepo)

	vaccinationRepo := repository.NewVaccinationRepository(db)
	vaccinationService := service.NewVaccinationService(vaccinationRepo)

	provinceStatsRepo := repository.NewProvinceStatsRepository(db)
	provinceStatsService := service.NewProvinceStatsService(provinceStatsRepo)

	// Override Swagger host/basePath from environment variables if set
	if host := os.Getenv("SWAGGER_HOST"); host != "" {
		docs.SwaggerInfo.Host = host
	}
	if basePath := os.Getenv("SWAGGER_BASE_PATH"); basePath != "" {
		docs.SwaggerInfo.BasePath = basePath
	}
	if schemes := os.Getenv("SWAGGER_SCHEMES"); schemes != "" {
		docs.SwaggerInfo.Schemes = []string{schemes}
	}

	enableSwagger := true
	svc := handler.Services{
		CovidService:     covidService,
		RegencyService:   regencyService,
		CacheInvalidator: cacheInvalidator,
		HospitalService:  hospitalService,
		TaskForceService:    taskForceService,
		VaccinationService:   vaccinationService,
		ProvinceStatsService: provinceStatsService,
	}
	router := handler.SetupRoutes(svc, db, enableSwagger)

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
