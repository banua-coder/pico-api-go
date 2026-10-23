package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/service"
)

// ContentHandler handles read-only site content endpoints
type ContentHandler struct {
	service service.ContentServiceInterface
}

// NewContentHandler creates a new ContentHandler
func NewContentHandler(service service.ContentServiceInterface) *ContentHandler {
	return &ContentHandler{service: service}
}

// maxContentPage caps ?page. Beyond it the request is rejected with 400, which
// prevents (page-1)*per_page from overflowing and keeps arbitrary page numbers
// from creating unbounded cache entries (the cache also skips empty pages).
const maxContentPage = 1000000

// parseStrictBool reads a boolean query param with strconv.ParseBool.
// Empty means "not provided" (ok=false); a non-empty unparsable value is an error.
func parseStrictBool(r *http.Request, name string) (val, ok bool, err error) {
	raw := r.URL.Query().Get(name)
	if raw == "" {
		return false, false, nil
	}
	val, err = strconv.ParseBool(raw)
	if err != nil {
		return false, false, fmt.Errorf("%s must be a boolean (true/false/1/0)", name)
	}
	return val, true, nil
}

// writeList runs fetch with the request's pagination (limit 0 = load_all) and writes the result.
// Repository errors are logged and reported as a generic 500.
func writeList(w http.ResponseWriter, r *http.Request, fetch func(limit, offset int) (interface{}, int, error)) {
	loadAll, _, err := parseStrictBool(r, "load_all")
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	p := parsePaginationParams(r)
	if !loadAll && p.Page > maxContentPage {
		writeErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("page must be at most %d", maxContentPage))
		return
	}
	limit := p.PerPage
	if loadAll {
		limit = 0
	}
	offset := 0
	if !loadAll {
		offset = (p.Page - 1) * p.PerPage // safe: Page <= maxContentPage, PerPage <= 100
	}
	items, total, err := fetch(limit, offset)
	if err != nil {
		log.Printf("content endpoint %s failed: %v", r.URL.Path, err)
		writeErrorResponse(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if loadAll {
		writeSuccessResponse(w, items)
		return
	}
	writePaginatedResponse(w, items, buildPaginationMeta(p, total))
}

// GetBanners godoc
// @Summary Get banners (paginated)
// @Description Returns banners ordered by id. Only active banners unless include_inactive=true. link is url_type + url as stored; image is returned as stored.
// @Description Response shape: with load_all=true data is a bare array; otherwise data is {data: [...], pagination: {...}}. Boolean params accept true/false/1/0 (case-insensitive); other values return 400. page is capped at 1000000.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Param include_inactive query bool false "Include banners with status 0"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /banners [get]
func (h *ContentHandler) GetBanners(w http.ResponseWriter, r *http.Request) {
	inactive, _, err := parseStrictBool(r, "include_inactive")
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetBanners(inactive, limit, offset)
	})
}

// GetInfographics godoc
// @Summary Get infographics (paginated)
// @Description Returns infographics newest first, each with images ordered by queue.
// @Description Response shape: with load_all=true data is a bare array; otherwise data is {data: [...], pagination: {...}}. Boolean params accept true/false/1/0 (case-insensitive); other values return 400. page is capped at 1000000.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /infographics [get]
func (h *ContentHandler) GetInfographics(w http.ResponseWriter, r *http.Request) {
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetInfographics(limit, offset)
	})
}

// GetDonations godoc
// @Summary Get donations (paginated)
// @Description Returns donation campaigns with bank accounts (including bank) and contacts.
// @Description Response shape: with load_all=true data is a bare array; otherwise data is {data: [...], pagination: {...}}. Boolean params accept true/false/1/0 (case-insensitive); other values return 400. page is capped at 1000000.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /donations [get]
func (h *ContentHandler) GetDonations(w http.ResponseWriter, r *http.Request) {
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetDonations(limit, offset)
	})
}

// GetTelemedicines godoc
// @Summary Get telemedicine providers (paginated)
// @Description Returns telemedicine providers with schedules and contacts.
// @Description Response shape: with load_all=true data is a bare array; otherwise data is {data: [...], pagination: {...}}. Boolean params accept true/false/1/0 (case-insensitive); other values return 400. page is capped at 1000000.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /telemedicines [get]
func (h *ContentHandler) GetTelemedicines(w http.ResponseWriter, r *http.Request) {
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetTelemedicines(limit, offset)
	})
}

// GetPartners godoc
// @Summary Get partners (paginated)
// @Description Returns partners and data sources. image is a file name as stored.
// @Description Response shape: with load_all=true data is a bare array; otherwise data is {data: [...], pagination: {...}}. Boolean params accept true/false/1/0 (case-insensitive); other values return 400. page is capped at 1000000.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Param in_footer query bool false "Filter on the in_footer flag"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /partners [get]
func (h *ContentHandler) GetPartners(w http.ResponseWriter, r *http.Request) {
	v, ok, err := parseStrictBool(r, "in_footer")
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	var inFooter *bool
	if ok {
		inFooter = &v
	}
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetPartners(inFooter, limit, offset)
	})
}

// GetContacts godoc
// @Summary Get contacts (paginated)
// @Description Returns contacts of hospitals, task forces, donations and telemedicines, each with its contact type.
// @Description Response shape: with load_all=true data is a bare array; otherwise data is {data: [...], pagination: {...}}. Boolean params accept true/false/1/0 (case-insensitive); other values return 400. page is capped at 1000000.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Param owner_type query string false "hospital, task_force, donation or telemedicine"
// @Param owner_id query int false "Positive id of the owner (requires owner_type)"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 500 {object} Response
// @Router /contacts [get]
func (h *ContentHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	ownerType := r.URL.Query().Get("owner_type")
	if _, ok := models.ContactOwnerClasses[ownerType]; ownerType != "" && !ok {
		writeErrorResponse(w, http.StatusBadRequest, "owner_type must be one of: hospital, task_force, donation, telemedicine")
		return
	}
	var ownerID int64
	if raw := r.URL.Query().Get("owner_id"); raw != "" {
		id, err := strconv.ParseInt(raw, 10, 64)
		if err != nil || id <= 0 {
			writeErrorResponse(w, http.StatusBadRequest, "owner_id must be a positive integer")
			return
		}
		if ownerType == "" {
			writeErrorResponse(w, http.StatusBadRequest, "owner_type is required when owner_id is given")
			return
		}
		ownerID = id
	}
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetContacts(ownerType, ownerID, limit, offset)
	})
}
