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
// @Summary Get all task force posts in Sulawesi Tengah
// @Description Returns all gugus tugas/posko grouped by regency with contacts
// @Tags task-forces
// @Produce json
// @Success 200 {object} Response
// @Router /api/v1/task-forces [get]
func (h *TaskForceHandler) GetTaskForces(w http.ResponseWriter, r *http.Request) {
	taskForces, err := h.service.GetTaskForces()
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeSuccessResponse(w, taskForces)
}
