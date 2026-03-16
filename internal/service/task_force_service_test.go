package service

import (
	"errors"
	"testing"

	"github.com/banua-coder/pico-api-go/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockTaskForceRepository mocks repository.TaskForceRepositoryInterface
type MockTaskForceRepository struct {
	mock.Mock
}

func (m *MockTaskForceRepository) GetPaginatedByProvinceID(provinceID, limit, offset int) ([]models.TaskForceByRegency, int, error) {
	args := m.Called(provinceID, limit, offset)
	return args.Get(0).([]models.TaskForceByRegency), args.Int(1), args.Error(2)
}

func (m *MockTaskForceRepository) GetAllByProvinceID(provinceID int) ([]models.TaskForceByRegency, error) {
	args := m.Called(provinceID)
	return args.Get(0).([]models.TaskForceByRegency), args.Error(1)
}

func setupTaskForceService() (*MockTaskForceRepository, *TaskForceService) {
	mockRepo := new(MockTaskForceRepository)
	svc := NewTaskForceService(mockRepo)
	return mockRepo, svc
}

func TestTaskForceService_GetTaskForces(t *testing.T) {
	mockRepo, svc := setupTaskForceService()

	expected := []models.TaskForceByRegency{
		{
			RegencyID:   7201,
			RegencyName: "Kabupaten Banggai",
			TaskForces: []models.TaskForce{
				{ID: 1, RegencyID: 7201, Name: "Posko COVID-19 Banggai"},
			},
		},
		{
			RegencyID:   7202,
			RegencyName: "Kabupaten Banggai Kepulauan",
			TaskForces:  []models.TaskForce{},
		},
	}
	mockRepo.On("GetAllByProvinceID", 72).Return(expected, nil)

	result, err := svc.GetTaskForces()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskForceService_GetTaskForces_Empty(t *testing.T) {
	mockRepo, svc := setupTaskForceService()

	mockRepo.On("GetAllByProvinceID", 72).Return([]models.TaskForceByRegency{}, nil)

	result, err := svc.GetTaskForces()

	assert.NoError(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskForceService_GetTaskForces_Error(t *testing.T) {
	mockRepo, svc := setupTaskForceService()

	mockRepo.On("GetAllByProvinceID", 72).Return([]models.TaskForceByRegency{}, errors.New("db error"))

	result, err := svc.GetTaskForces()

	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestTaskForceService_GetTaskForcesPaginated(t *testing.T) {
	mockRepo, svc := setupTaskForceService()

	expected := []models.TaskForceByRegency{
		{RegencyID: 7201, RegencyName: "Kabupaten Banggai", TaskForces: []models.TaskForce{
			{ID: 1, RegencyID: 7201, Name: "Posko COVID-19 Banggai"},
		}},
	}
	mockRepo.On("GetPaginatedByProvinceID", 72, 10, 0).Return(expected, 1, nil)

	result, total, err := svc.GetTaskForcesPaginated(10, 0)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	assert.Equal(t, 1, total)
	mockRepo.AssertExpectations(t)
}

func TestTaskForceService_GetTaskForcesPaginated_Error(t *testing.T) {
	mockRepo, svc := setupTaskForceService()

	mockRepo.On("GetPaginatedByProvinceID", 72, 10, 0).Return([]models.TaskForceByRegency{}, 0, errors.New("db error"))

	result, total, err := svc.GetTaskForcesPaginated(10, 0)

	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Equal(t, 0, total)
	mockRepo.AssertExpectations(t)
}
