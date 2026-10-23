package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/service"
	"github.com/banua-coder/pico-api-go/pkg/utils"
)

// ContentHandler handles read-only site content endpoints
type ContentHandler struct {
	service service.ContentServiceInterface
}

// NewContentHandler creates a new ContentHandler
func NewContentHandler(service service.ContentServiceInterface) *ContentHandler {
	return &ContentHandler{service: service}
}

// writeList runs fetch with the request's pagination (limit 0 = load_all) and writes the result.
func writeList(w http.ResponseWriter, r *http.Request, fetch func(limit, offset int) (interface{}, int, error)) {
	p := parsePaginationParams(r)
	limit := p.PerPage
	if p.LoadAll {
		limit = 0
	}
	items, total, err := fetch(limit, p.Offset)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	if p.LoadAll {
		writeSuccessResponse(w, items)
		return
	}
	writePaginatedResponse(w, items, buildPaginationMeta(p, total))
}

// GetBanners godoc
// @Summary Get banners (paginated)
// @Description Returns banners ordered by id. Only active banners unless include_inactive=true. link is url_type + url as stored; image is returned as stored.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Param include_inactive query bool false "Include banners with status 0"
// @Success 200 {object} Response
// @Router /banners [get]
func (h *ContentHandler) GetBanners(w http.ResponseWriter, r *http.Request) {
	inactive := utils.ParseBoolQueryParam(r, "include_inactive")
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetBanners(inactive, limit, offset)
	})
}

// GetInfographics godoc
// @Summary Get infographics (paginated)
// @Description Returns infographics newest first, each with images ordered by queue.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Success 200 {object} Response
// @Router /infographics [get]
func (h *ContentHandler) GetInfographics(w http.ResponseWriter, r *http.Request) {
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetInfographics(limit, offset)
	})
}

// GetDonations godoc
// @Summary Get donations (paginated)
// @Description Returns donation campaigns with bank accounts (including bank) and contacts.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Success 200 {object} Response
// @Router /donations [get]
func (h *ContentHandler) GetDonations(w http.ResponseWriter, r *http.Request) {
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetDonations(limit, offset)
	})
}

// GetTelemedicines godoc
// @Summary Get telemedicine providers (paginated)
// @Description Returns telemedicine providers with schedules and contacts.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Success 200 {object} Response
// @Router /telemedicines [get]
func (h *ContentHandler) GetTelemedicines(w http.ResponseWriter, r *http.Request) {
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetTelemedicines(limit, offset)
	})
}

// GetPartners godoc
// @Summary Get partners (paginated)
// @Description Returns partners and data sources. image is a file name as stored.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Param in_footer query bool false "Filter on the in_footer flag"
// @Success 200 {object} Response
// @Router /partners [get]
func (h *ContentHandler) GetPartners(w http.ResponseWriter, r *http.Request) {
	var inFooter *bool
	if r.URL.Query().Get("in_footer") != "" {
		v := utils.ParseBoolQueryParam(r, "in_footer")
		inFooter = &v
	}
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetPartners(inFooter, limit, offset)
	})
}

// GetContacts godoc
// @Summary Get contacts (paginated)
// @Description Returns contacts of hospitals, task forces, donations and telemedicines, each with its contact type.
// @Tags content
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all without pagination"
// @Param owner_type query string false "hospital, task_force, donation or telemedicine"
// @Param owner_id query int false "Id of the owner (requires nothing else, usually combined with owner_type)"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /contacts [get]
func (h *ContentHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	ownerType := r.URL.Query().Get("owner_type")
	if _, ok := models.ContactOwnerClasses[ownerType]; ownerType != "" && !ok {
		writeErrorResponse(w, http.StatusBadRequest, "owner_type must be one of: hospital, task_force, donation, telemedicine")
		return
	}
	ownerID := int64(utils.ParseIntQueryParam(r, "owner_id", 0))
	writeList(w, r, func(limit, offset int) (interface{}, int, error) {
		return h.service.GetContacts(ownerType, ownerID, limit, offset)
	})
}
