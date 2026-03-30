package handler

import (
	"net/http"
	"os"

	"github.com/banua-coder/pico-api-go/internal/service"
)

// AdminHandler handles admin endpoints.
type AdminHandler struct {
	invalidator service.CacheInvalidator
}

// NewAdminHandler creates a new AdminHandler.
func NewAdminHandler(invalidator service.CacheInvalidator) *AdminHandler {
	return &AdminHandler{invalidator: invalidator}
}

// ClearCache godoc
//
//	@Summary		Clear all in-memory cache
//	@Description	Clears all cached data. Requires X-Admin-Key header matching ADMIN_KEY env var.
//	@Tags			admin
//	@Produce		json
//	@Param			X-Admin-Key	header		string	true	"Admin key"
//	@Success		200			{object}	map[string]string
//	@Failure		401			{object}	map[string]string
//	@Router			/admin/cache/clear [post]
func (h *AdminHandler) ClearCache(w http.ResponseWriter, r *http.Request) {
	adminKey := os.Getenv("ADMIN_KEY")
	if adminKey == "" || r.Header.Get("X-Admin-Key") != adminKey {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"unauthorized"}`)) //nolint:errcheck
		return
	}
	h.invalidator.Clear()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"cache cleared"}`)) //nolint:errcheck
}
