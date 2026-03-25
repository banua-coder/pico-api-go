package service

import (
	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/banua-coder/pico-api-go/internal/repository"
)

// TaskForceService handles business logic for task forces
type TaskForceService struct {
	taskForceRepo repository.TaskForceRepositoryInterface
}

// NewTaskForceService creates a new TaskForceService
func NewTaskForceService(taskForceRepo repository.TaskForceRepositoryInterface) *TaskForceService {
	return &TaskForceService{taskForceRepo: taskForceRepo}
}

// GetTaskForces returns all task forces grouped by regency in SulTeng (province_id=72)
func (s *TaskForceService) GetTaskForces() ([]models.TaskForceByRegency, error) {
	return s.taskForceRepo.GetAllByProvinceID(72)
}

// GetTaskForcesPaginated returns a page of task forces grouped by regency with total count
func (s *TaskForceService) GetTaskForcesPaginated(limit, offset int) ([]models.TaskForceByRegency, int, error) {
	return s.taskForceRepo.GetPaginatedByProvinceID(72, limit, offset)
}
