package tests

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/google/uuid"
    api "github.com/niv-e/phonebook-api/internal/delivery/http/endpoint"
    "github.com/stretchr/testify/assert"
)

func TestDeleteContactHttpHandler_Success(t *testing.T) {
    mockRepo := &MockContactRepository{}
    id := uuid.New()

    req, err := http.NewRequest("DELETE", "/contacts?id="+id.String(), nil)
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(api.DeleteContactHttpHandler(mockRepo))
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusNoContent, rr.Code)
    assert.Len(t, mockRepo.Contacts, 0)
}

func TestDeleteContactHttpHandler_InvalidID_ShouldReturnInvalidPayload(t *testing.T) {
    req, err := http.NewRequest("DELETE", "/contacts?id=invalid-uuid", nil)
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(api.DeleteContactHttpHandler(nil))
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusBadRequest, rr.Code)
    assert.Equal(t, "Invalid request payload\n", rr.Body.String())
}

func TestDeleteContactHttpHandler_InternalServerError_ShouldReturnInternalServerError(t *testing.T) {
    mockRepo := &MockContactRepository{
        Err: assert.AnError,
    }
    id := uuid.New()

    req, err := http.NewRequest("DELETE", "/contacts?id="+id.String(), nil)
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(api.DeleteContactHttpHandler(mockRepo))
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusInternalServerError, rr.Code)
}