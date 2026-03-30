package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCacheInvalidator struct {
	mock.Mock
}

func (m *MockCacheInvalidator) Clear() {
	m.Called()
}

func TestNewAdminHandler(t *testing.T) {
	invalidator := new(MockCacheInvalidator)
	h := NewAdminHandler(invalidator)
	assert.NotNil(t, h)
}

func TestAdminHandler_ClearCache_Success(t *testing.T) {
	t.Setenv("ADMIN_KEY", "test-secret-key")

	invalidator := new(MockCacheInvalidator)
	invalidator.On("Clear").Once()

	h := NewAdminHandler(invalidator)

	req := httptest.NewRequest(http.MethodPost, "/admin/cache/clear", nil)
	req.Header.Set("X-Admin-Key", "test-secret-key")
	w := httptest.NewRecorder()

	h.ClearCache(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "cache cleared")
	invalidator.AssertExpectations(t)
}

func TestAdminHandler_ClearCache_WrongKey(t *testing.T) {
	t.Setenv("ADMIN_KEY", "test-secret-key")

	invalidator := new(MockCacheInvalidator)
	h := NewAdminHandler(invalidator)

	req := httptest.NewRequest(http.MethodPost, "/admin/cache/clear", nil)
	req.Header.Set("X-Admin-Key", "wrong-key")
	w := httptest.NewRecorder()

	h.ClearCache(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "unauthorized")
	invalidator.AssertNotCalled(t, "Clear")
}

func TestAdminHandler_ClearCache_NoKey(t *testing.T) {
	t.Setenv("ADMIN_KEY", "test-secret-key")

	invalidator := new(MockCacheInvalidator)
	h := NewAdminHandler(invalidator)

	req := httptest.NewRequest(http.MethodPost, "/admin/cache/clear", nil)
	w := httptest.NewRecorder()

	h.ClearCache(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	invalidator.AssertNotCalled(t, "Clear")
}

func TestAdminHandler_ClearCache_EmptyAdminKeyEnv(t *testing.T) {
	t.Setenv("ADMIN_KEY", "")

	invalidator := new(MockCacheInvalidator)
	h := NewAdminHandler(invalidator)

	req := httptest.NewRequest(http.MethodPost, "/admin/cache/clear", strings.NewReader(""))
	req.Header.Set("X-Admin-Key", "any-key")
	w := httptest.NewRecorder()

	h.ClearCache(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	invalidator.AssertNotCalled(t, "Clear")
}
