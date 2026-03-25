package handler

import (
	"net/http"

	"github.com/banua-coder/pico-api-go/internal/service"
)

// TaskForceHandler handles HTTP requests for task force endpoints
type TaskForceHandler struct {
	service service.TaskForceServiceInterface
}

// NewTaskForceHandler creates a new TaskForceHandler
func NewTaskForceHandler(service service.TaskForceServiceInterface) *TaskForceHandler {
	return &TaskForceHandler{service: service}
}

// GetTaskForces godoc
// @Summary Get task force posts in Sulawesi Tengah (paginated)
// @Description Returns paginated gugus tugas/posko grouped by regency. Use ?load_all=true for all.
// @Tags task-forces
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param per_page query int false "Items per page (default: 10, max: 100)"
// @Param load_all query bool false "Set true to return all task forces without pagination"
// @Success 200 {object} Response
// @Router /task-forces [get]
func (h *TaskForceHandler) GetTaskForces(w http.ResponseWriter, r *http.Request) {
	p := parsePaginationParams(r)

	if p.LoadAll {
		taskForces, err := h.service.GetTaskForces()
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		writeSuccessResponse(w, taskForces)
		return
	}

	taskForces, total, err := h.service.GetTaskForcesPaginated(p.PerPage, p.Offset)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writePaginatedResponse(w, taskForces, buildPaginationMeta(p, total))
}
